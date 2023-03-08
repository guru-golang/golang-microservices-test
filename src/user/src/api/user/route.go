package user

import (
	"car-rent-platform/backend/common/src/lib/builtin_lib"
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/common/src/repository/user"
	"car-rent-platform/backend/user/src/api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	RouteInterface interface {
		common.FullRouteInterface
	}

	Route struct {
		service UserInterface
	}
)

func NewRoute() RouteInterface {
	var i = Route{}
	return &i
}

func (route *Route) Init(g *gin_lib.Gin, r *repository.Repository, gr *gin.RouterGroup) *gin.RouterGroup {
	route.service = NewUserService(r)

	var rg *gin.RouterGroup
	if gr == nil {
		rg = g.Route(builtin_lib.GetLocalPkgName()).Group
	} else {
		rg = gr.Group(builtin_lib.GetLocalPkgName())
	}
	rg.GET("", route.FindAll)
	rg.POST("", route.Create)
	rg.GET(":uuid", route.FindOne)
	rg.PATCH(":uuid", route.Update)
	rg.DELETE(":uuid", route.Remove)

	return rg
}

func (route *Route) FindAll(ctx *gin.Context) {
	var input wql_lib.FilterInput
	if filterInput, err := input.GinScan(ctx); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := route.service.FindAll(filterInput); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out, "meta": filterInput.QueryMeta})
	}
}

func (route *Route) FindOne(ctx *gin.Context) {
	var input user.Input
	uuid := ctx.Params.ByName("uuid")
	input.UUID = &uuid
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := route.service.FindOne(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}

func (route *Route) Create(ctx *gin.Context) {
	var input user.Input
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := route.service.Create(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}

func (route *Route) Update(ctx *gin.Context) {
	var input user.Input
	uuid := ctx.Params.ByName("uuid")
	input.UUID = &uuid
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := route.service.Update(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}

func (route *Route) Remove(ctx *gin.Context) {
	var input user.Input
	uuid := ctx.Params.ByName("uuid")
	input.UUID = &uuid
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := route.service.Remove(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}
