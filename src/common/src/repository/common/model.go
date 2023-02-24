package common

import (
	"car-rent-platform/backend/common/src/lib/gorm_lib"
	"car-rent-platform/backend/common/src/lib/strings_lib"
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"fmt"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type (
	Repo[Input any, Output any] struct {
		db       *gorm_lib.Gorm
		validate *validator.Validate

		table  string
		input  Input
		output Output
	}
	Model[Type any] struct {
		UUID *string `gorm:"column:uuid;default:uuid_generate_v4();type:uuid;primarykey" json:"uuid" validate:"omitempty"`

		Sequence  *uint      `gorm:"column:sequence;type:uint;index" json:"sequence" validate:"omitempty"`
		CreatedAt *time.Time `gorm:"column:createdAt;type:timestamp;index" json:"createdAt" validate:"omitempty"`
		UpdatedAt *time.Time `gorm:"column:updatedAt;type:timestamp;index" json:"updatedAt" validate:"omitempty"`
		DeletedAt *time.Time `gorm:"column:deletedAt;type:timestamp;index" json:"deletedAt" validate:"omitempty"`
	}
)

func (r *Repo[Input, Output]) Init(input Input, output Output, table string, db *gorm_lib.Gorm) {
	r.SetDb(db)
	r.SetTable(table)
	r.SetInput(input)
	r.SetOutput(output)
	r.validate = validator.New()

	r.validate.RegisterStructValidation(r.inputStruct, r.Input())
	r.validate.RegisterStructValidation(r.outputStruct, r.Output())
}

func (r *Repo[Input, Output]) Scan(filter *wql_lib.FilterInput) *gorm.DB {
	db := r.db.Base.Table(r.Table()).Model(r.Input())

	db = r.BuildFilterMeta(filter, db)
	db = r.BuildQueryMeta(filter, db)
	db = r.BuildOrderMeta(filter, db)
	db = r.BuildAttributeMeta(filter, db)
	db = r.BuildIncludeMeta(filter, db)

	return db
}

func (r *Repo[Input, Output]) BuildFilterMeta(filter *wql_lib.FilterInput, scan *gorm.DB) *gorm.DB {
	if filter.FilterMeta != nil {
		return scan.Where(filter.FilterMeta)
	}
	return scan
}

func (r *Repo[Input, Output]) BuildQueryMeta(filter *wql_lib.FilterInput, scan *gorm.DB) *gorm.DB {
	if filter.QueryMeta != nil {
		page, _ := strconv.Atoi(filter.QueryMeta["page"])
		count, _ := strconv.Atoi(filter.QueryMeta["count"])
		offset := (page - 1) * count

		scan = scan.Offset(offset).Limit(count)
	}
	return scan
}

func (r *Repo[Input, Output]) BuildOrderMeta(filter *wql_lib.FilterInput, scan *gorm.DB) *gorm.DB {
	if filter.OrderMeta != nil {
		order := []string{"asc", "desc"}
		for key, orderMeta := range filter.OrderMeta {
			if strings_lib.Contains(order, orderMeta) {
				scan = scan.Order(fmt.Sprintf("%v %v", key, orderMeta))
			}
		}
	}
	return scan
}

func (r *Repo[Input, Output]) BuildIncludeMeta(filter *wql_lib.FilterInput, scan *gorm.DB) *gorm.DB {
	if filter.IncludeMeta != nil {
		for key, includeMeta := range filter.IncludeMeta {
			if includeMeta == "join" {
				scan = scan.Joins(key)
			}
			if includeMeta == "innerJoin" {
				scan = scan.InnerJoins(key)
			}
		}
	}
	return scan
}

func (r *Repo[Input, Output]) BuildAttributeMeta(filter *wql_lib.FilterInput, scan *gorm.DB) *gorm.DB {
	if filter.AttributeMeta != nil {
		for key, attributeMeta := range filter.AttributeMeta {
			if attributeMeta == "pick" {
				scan = scan.Select(key)
			}
			if attributeMeta == "omit" {
				scan = scan.Omit(key)
			}
		}
	}
	return scan
}

func (r *Repo[Input, Output]) SetTable(name string) {
	r.table = name
}

func (r *Repo[Input, Output]) Table() string {
	return r.table
}

func (r *Repo[Input, Output]) SetDb(db *gorm_lib.Gorm) {
	r.db = db
}

func (r *Repo[Input, Output]) Db() *gorm.DB {
	return r.db.Base.Table(r.Table()).Model(r.Input())
}

func (r *Repo[Input, Output]) Output() *Output {
	return &r.output
}

func (r *Repo[Input, Output]) SetOutput(output Output) {
	r.output = output
}

func (r *Repo[Input, Output]) Input() *Input {
	return &r.input
}

func (r *Repo[Input, Output]) SetInput(input Input) {
	r.input = input
}

func (r *Repo[Input, Output]) OutputValidate(out *Output) (result *Output, err error) {
	result = out
	if err = r.validate.Struct(result); err != nil {
		return
	}
	return
}

func (r *Repo[Input, Output]) OutputListValidate(out *[]*Output) (result []*Output, err error) {
	swap, swapSlice := *out, make([]*Output, 0)
	for rowI := 0; len(swap) > rowI; rowI++ {
		if res, err := r.OutputValidate(swap[rowI]); err != nil {
			return nil, err
		} else {
			swapSlice = append(swapSlice, res)
		}
	}
	return swapSlice, nil
}

func (r *Repo[Input, Output]) InputValidate(input *Input) (result *Input, err error) {
	result = input
	if err = r.validate.Struct(result); err != nil {
		return
	}
	return
}

func (r *Repo[Input, Output]) outputStruct(sl validator.StructLevel) {
	// fmt.Println("outputStruct")
	return
}

func (r *Repo[Input, Output]) inputStruct(sl validator.StructLevel) {
	// fmt.Println("inputStruct")
	return
}
