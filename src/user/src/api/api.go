package api

import (
	"car-rent-platform/backend/common/src/lib/config_lib"
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/user/src/api/user"
	"car-rent-platform/backend/user/src/api/user/profile"
)

type (
	Interface interface {
		Init(g *gin_lib.Gin, r *repository.Repository)
	}
	API struct {
		User        user.Interface
		UserProfile profile.Interface
	}
)

func NewAPI() Interface {
	var i = API{
		User:        user.NewRoute(),
		UserProfile: profile.NewRoute(),
	}
	return &i
}

func (a *API) Init(g *gin_lib.Gin, r *repository.Repository) {
	br := g.Route(config_lib.Config.Get("server_gin_version").(string))
	rg := a.User.Init(g, r, br.Group)
	{
		a.UserProfile.Init(g, r, rg)
	}
}
