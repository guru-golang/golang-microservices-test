package api

import (
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/user/src/api/user"
)

type (
	Interface interface {
		Init(g *gin_lib.Gin, r *repository.Repository)
	}
	API struct {
		Auth user.Interface
	}
)

func NewAPI() Interface {
	var i = API{Auth: user.NewRoute()}
	return &i
}

func (a *API) Init(g *gin_lib.Gin, r *repository.Repository) {
	a.Auth.Init(g, r)
}
