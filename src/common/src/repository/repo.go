package repository

import (
	"car-rent-platform/backend/common/src/lib/gorm_lib"
	"car-rent-platform/backend/common/src/repository/configuration"
	"car-rent-platform/backend/common/src/repository/user"
	user_notification "car-rent-platform/backend/common/src/repository/user-notification"
	user_portfolio "car-rent-platform/backend/common/src/repository/user-portfolio"
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
	r.models["userNotification"] = user_notification.New(user_notification.Input{}, user_notification.Output{}, db)
	r.models["userPortfolio"] = user_portfolio.New(user_portfolio.Input{}, user_portfolio.Output{}, db)
	return r
}

func (r *Repository) Model(name string) any {
	return r.models[name]
}
