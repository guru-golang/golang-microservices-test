package repository

import (
	"car-rent-platform/backend/common/src/lib/gorm_lib"
	"car-rent-platform/backend/common/src/repository/configuration"
	"car-rent-platform/backend/common/src/repository/user"
	userprofile "car-rent-platform/backend/common/src/repository/user-profile"
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
	r.models["configuration"] = configuration.New(configuration.Input{}, configuration.Output{}, db)
	r.models["userProfile"] = userprofile.New(userprofile.Input{}, userprofile.Output{}, db)
	r.models["user"] = user.New(user.Input{}, user.Output{}, db)
	return r
}

func (r *Repository) Model(name string) any {
	return r.models[name]
}
