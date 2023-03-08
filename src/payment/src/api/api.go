package api

import (
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/payment/src/api/binance"
)

type (
	Interface interface {
		InitRoute(g *gin_lib.Gin, r *repository.Repository)
		InitRpc(n *net_lib.Net, r *repository.Repository)
	}
	API struct {
		Route struct {
			Binance binance.RouteInterface
		}
		Rpc struct {
			Binance binance.RpcInterface
		}
	}
)

func NewAPI() Interface {
	i := API{
		Route: struct {
			Binance binance.RouteInterface
		}{
			Binance: binance.NewRoute(),
		},
		Rpc: struct {
			Binance binance.RpcInterface
		}{
			Binance: binance.NewRpc(),
		},
	}
	return &i
}

func (a *API) InitRoute(g *gin_lib.Gin, r *repository.Repository) {
	br := g.Route(g.Conf.Version)
	_ = a.Route.Binance.Init(g, r, br.Group)
}

func (a *API) InitRpc(n *net_lib.Net, r *repository.Repository) {
	a.Rpc.Binance.Init(n, r)
}
