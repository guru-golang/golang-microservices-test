package auth

import (
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"github.com/gin-gonic/gin"

	"car-rent-platform/backend/auth/src/api/common"
)

type (
	Interface interface {
		common.FullInterface
		Init(g *gin_lib.Gin)
	}

	Route struct {
		Service *Service
	}
)

func NewRoute() Interface {
	var i = Route{}
	return &i
}

func (r *Route) Init(g *gin_lib.Gin) {
	route := g.Route("auth")
	route.Group.GET("", r.FindAll)
	route.Group.POST("", r.Create)
	route.Group.GET(":uuid", r.FindOne)
	route.Group.PATCH(":uuid", r.Update)
	route.Group.DELETE(":uuid", r.Remove)
}

func (r *Route) FindAll(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r *Route) FindOne(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r *Route) Create(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r *Route) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r *Route) Remove(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
