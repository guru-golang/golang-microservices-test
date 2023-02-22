package repository

import (
	"car-rent-platform/backend/common/src/lib/gorm_lib"
	"car-rent-platform/backend/common/src/repository/user"
)

type (
	Repository struct {
		models map[string]any
	}
)

func NewRepository() *Repository {
	var i Repository
	return &i
}

func (r *Repository) Init(db *gorm_lib.Gorm) *Repository {
	r.models = make(map[string]any)
	r.models["user"] = user.New(user.Input{}, user.Output{}, db)
	return r
}

func (r *Repository) Model(name string) any {
	return r.models[name]
}
