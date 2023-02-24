package schema

import (
	"car-rent-platform/backend/common/src/repository/common"
)

type (
	Configuration struct {
		common.Model[Configuration]

		Module string `gorm:"column:module; type:text;index" json:"module" validate:"omitempty"`
		Name   string `gorm:"column:name; type:text;index" json:"name" validate:"omitempty"`
		Value  string `gorm:"column:value; type:jsonb" json:"value" validate:"omitempty"`
	}
)

func (Configuration) TableName() string {
	return "configurations"
}
