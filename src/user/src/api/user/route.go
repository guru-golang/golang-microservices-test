package user

import (
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/common/src/repository/user"
	"car-rent-platform/backend/user/src/api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	Interface interface {
		common.FullInterface
		Init(g *gin_lib.Gin, r *repository.Repository)
	}

	Route struct {
		Service AuthInterface
	}
)

func NewRoute() Interface {
	var i = Route{}
	return &i
}

func (r *Route) Init(g *gin_lib.Gin, repo *repository.Repository) {
	r.Service = NewAuthService(repo)
	route := g.Route("user")
	route.Group.GET("", r.FindAll)
	route.Group.POST("", r.Create)
	route.Group.GET(":uuid", r.FindOne)
	route.Group.PATCH(":uuid", r.Update)
	route.Group.DELETE(":uuid", r.Remove)
}

func (r *Route) FindAll(ctx *gin.Context) {
	var input user.Input
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := r.Service.FindAll(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}

func (r *Route) FindOne(ctx *gin.Context) {
	var input user.Input
	uuid := ctx.Params.ByName("uuid")
	input.UUID = &uuid
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := r.Service.FindOne(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}

func (r *Route) Create(ctx *gin.Context) {
	var input user.Input
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := r.Service.Create(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}

func (r *Route) Update(ctx *gin.Context) {
	var input user.Input
	uuid := ctx.Params.ByName("uuid")
	input.UUID = &uuid
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := r.Service.Update(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}

func (r *Route) Remove(ctx *gin.Context) {
	var input user.Input
	uuid := ctx.Params.ByName("uuid")
	input.UUID = &uuid
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else if out, err := r.Service.Remove(&input); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "reason": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": true, "result": out})
	}
}
