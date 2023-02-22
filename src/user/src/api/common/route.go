package common

import (
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/repository"
	"github.com/gin-gonic/gin"
)

type (
	CommonInterface interface {
		Init(g *gin_lib.Gin, r *repository.Repository)
	}
	FullInterface interface {
		CommonInterface
		FindAll(ctx *gin.Context)
		FindOne(ctx *gin.Context)
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
		Remove(ctx *gin.Context)
	}
)
