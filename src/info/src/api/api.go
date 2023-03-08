package api

import (
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/info/src/api/configuration"
)

type (
	Interface interface {
		InitRoute(g *gin_lib.Gin, r *repository.Repository)
		InitRpc(n *net_lib.Net, r *repository.Repository)
	}
	API struct {
		Route struct {
			Configuration configuration.RouteInterface
		}
		Rpc struct {
			Configuration configuration.RpcInterface
		}
	}
)

func NewAPI() Interface {
	i := API{
		Route: struct {
			Configuration configuration.RouteInterface
		}{
			Configuration: configuration.NewRoute(),
		},
		Rpc: struct {
			Configuration configuration.RpcInterface
		}{
			Configuration: configuration.NewRpc(),
		},
	}
	return &i
}

func (a *API) InitRoute(g *gin_lib.Gin, r *repository.Repository) {
	br := g.Route(g.Conf.Version)
	_ = a.Route.Configuration.Init(g, r, br.Group)
}

func (a *API) InitRpc(n *net_lib.Net, r *repository.Repository) {
	a.Rpc.Configuration.Init(n, r)
}
