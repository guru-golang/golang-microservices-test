package api

import (
	"car-rent-platform/backend/auth/src/api/auth"
	"car-rent-platform/backend/common/src/lib/gin_lib"
)

type (
	Interface interface {
		Init(g *gin_lib.Gin)
	}
	API struct {
		Auth auth.Interface
	}
)

func NewAPI() Interface {
	var i = API{Auth: auth.NewRoute()}
	return &i
}

func (a *API) Init(g *gin_lib.Gin) {
	a.Auth.Init(g)
}
