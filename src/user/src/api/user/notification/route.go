package notification

import (
	"car-rent-platform/backend/common/src/lib/builtin_lib"
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	user_notification "car-rent-platform/backend/common/src/repository/user-notification"
	"car-rent-platform/backend/user/src/api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	RouteInterface interface {
		common.FullRouteInterface
	}

	Route struct {
		service UserNotificationInterface
	}
)

func NewRoute() RouteInterface {
	var i = Route{}
	return &i
}

func (r *Route) Init(g *gin_lib.Gin, repo *repository.Repository, gr *gin.RouterGroup) *gin.RouterGroup {
	r.service = NewUserNotificationService(repo)

	var route *gin.RouterGroup
	if gr == nil {
		route = g.Route(builtin_lib.GetLocalPkgName()).Group
	} else {
		route = gr.Group(builtin_lib.GetLocalPkgName())
	}
	route.GET("", r.FindAll)
	route.POST("", r.Create)
	route.GET(":uuid", r.FindOne)
	route.PATCH(":uuid", r.Update)
	route.DELETE(":uuid", r.Remove)

	return route
}

func (r *Route) FindAll(ctx *gin.Context) {
	var input wql_lib.FilterInput
	if filterInput, err := input.GinScan(ctx); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := r.service.FindAll(filterInput); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out, "meta": filterInput.QueryMeta})
	}
}

func (r *Route) FindOne(ctx *gin.Context) {
	var input user_notification.Input
	uuid := ctx.Params.ByName("uuid")
	input.UUID = &uuid
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := r.service.FindOne(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}

func (r *Route) Create(ctx *gin.Context) {
	var input user_notification.Input
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := r.service.Create(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}

func (r *Route) Update(ctx *gin.Context) {
	var input user_notification.Input
	uuid := ctx.Params.ByName("uuid")
	input.UUID = &uuid
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := r.service.Update(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}

func (r *Route) Remove(ctx *gin.Context) {
	var input user_notification.Input
	uuid := ctx.Params.ByName("uuid")
	input.UUID = &uuid
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := r.service.Remove(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}
