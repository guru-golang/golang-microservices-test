package configuration

import (
	common "car-rent-platform/backend/common/src/context"
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/common/src/repository/configuration"
	"github.com/gin-gonic/gin"
)

type (
	RpcInterface interface {
		common.FullRpcInterface
	}

	Rpc struct {
		service ConfigurationInterface
	}
)

func (rpc *Rpc) Init(n *net_lib.Net, r *repository.Repository) {
	rpc.service = NewConfigurationService(r)
	n.Pattern(net_lib.UserFindAll, rpc.FindAll)
	n.Pattern(net_lib.UserFindOne, rpc.FindOne)
	n.Pattern(net_lib.UserCreate, rpc.Create)
	n.Pattern(net_lib.UserUpdate, rpc.Update)
	n.Pattern(net_lib.UserRemove, rpc.Remove)
}

func (rpc *Rpc) FindAll(ctx *net_lib.Context) {
	var input wql_lib.FilterInput

	if filterInput, err := input.NetScan(ctx); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.FindAll(filterInput); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out, "meta": filterInput.QueryMeta})
	}
}

func (rpc *Rpc) FindOne(ctx *net_lib.Context) {
	var input configuration.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.FindOne(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Create(ctx *net_lib.Context) {
	var input configuration.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.Create(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Update(ctx *net_lib.Context) {
	var input configuration.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.Update(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Remove(ctx *net_lib.Context) {
	var input configuration.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.Remove(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func NewRpc() RpcInterface {
	var i = Rpc{}
	return &i
}
