package api

import (
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/worker/src/api/notification"
)

type (
	Interface interface {
		InitRoute(g *gin_lib.Gin, r *repository.Repository)
		InitRpc(n *net_lib.Net, r *repository.Repository)
	}
	API struct {
		Route struct {
			Notification notification.RouteInterface
		}
		Rpc struct {
			Notification notification.RpcInterface
		}
	}
)

func NewAPI() Interface {
	i := API{
		Route: struct {
			Notification notification.RouteInterface
		}{
			Notification: notification.NewRoute(),
		},
		Rpc: struct {
			Notification notification.RpcInterface
		}{
			Notification: notification.NewRpc(),
		},
	}
	return &i
}

func (a *API) InitRoute(g *gin_lib.Gin, r *repository.Repository) {
	br := g.Route(g.Conf.Version)
	_ = a.Route.Notification.Init(g, r, br.Group)
}

func (a *API) InitRpc(n *net_lib.Net, r *repository.Repository) {
	a.Rpc.Notification.Init(n, r)
}
