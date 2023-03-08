package api

import (
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/wallet/src/api/portfolio"
)

type (
	Interface interface {
		InitRoute(g *gin_lib.Gin, r *repository.Repository)
		InitRpc(n *net_lib.Net, r *repository.Repository)
	}
	API struct {
		Route struct {
			Portfolio portfolio.RouteInterface
		}
		Rpc struct {
			Portfolio portfolio.RpcInterface
		}
	}
)

func NewAPI() Interface {
	i := API{
		Route: struct {
			Portfolio portfolio.RouteInterface
		}{Portfolio: portfolio.NewRoute()},
		Rpc: struct {
			Portfolio portfolio.RpcInterface
		}{Portfolio: portfolio.NewRpc()},
	}
	return &i
}

func (a *API) InitRoute(g *gin_lib.Gin, r *repository.Repository) {
	br := g.Route(g.Conf.Version)
	_ = a.Route.Portfolio.Init(g, r, br.Group)
}

func (a *API) InitRpc(n *net_lib.Net, r *repository.Repository) {
	a.Rpc.Portfolio.Init(n, r)
}
