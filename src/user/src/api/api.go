package api

import (
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/user/src/api/user"
	"car-rent-platform/backend/user/src/api/user/profile"
)

type (
	Interface interface {
		InitRoute(g *gin_lib.Gin, r *repository.Repository)
		InitRpc(n *net_lib.Net, r *repository.Repository)
	}
	API struct {
		Route struct {
			User        user.RouteInterface
			UserProfile profile.RouteInterface
		}
		Rpc struct {
			User user.RpcInterface
		}
	}
)

func NewAPI() Interface {
	i := API{
		Route: struct {
			User        user.RouteInterface
			UserProfile profile.RouteInterface
		}{
			User:        user.NewRoute(),
			UserProfile: profile.NewRoute(),
		},
		Rpc: struct {
			User user.RpcInterface
		}{
			User: user.NewRpc(),
		},
	}
	return &i
}

func (a *API) InitRoute(g *gin_lib.Gin, r *repository.Repository) {
	br := g.Route(g.Conf.Version)
	rg := a.Route.User.Init(g, r, br.Group)
	{
		a.Route.UserProfile.Init(g, r, rg)
	}
}

func (a *API) InitRpc(n *net_lib.Net, r *repository.Repository) {
	a.Rpc.User.Init(n, r)
}
