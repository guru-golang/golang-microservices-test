package common

import (
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"github.com/gin-gonic/gin"
)

type (
	CommonInterface interface {
		Init(g *gin_lib.Gin)
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
