package profile

import (
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	user_profile "car-rent-platform/backend/common/src/repository/user-profile"
	"car-rent-platform/backend/user/src/api/common"
	"github.com/gin-gonic/gin"
)

type (
	RpcInterface interface {
		common.FullRpcInterface
	}

	Rpc struct {
		service UserProfileInterface
	}
)

func (rpc *Rpc) Init(n *net_lib.Net, r *repository.Repository) {
	rpc.service = NewUserProfileService(r)
	n.Pattern(net_lib.UserProfileFindAll, rpc.FindAll)
	n.Pattern(net_lib.UserProfileFindOne, rpc.FindOne)
	n.Pattern(net_lib.UserProfileCreate, rpc.Create)
	n.Pattern(net_lib.UserProfileUpdate, rpc.Update)
	n.Pattern(net_lib.UserProfileRemove, rpc.Remove)
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
	var input user_profile.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.FindOne(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Create(ctx *net_lib.Context) {
	var input user_profile.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.Create(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Update(ctx *net_lib.Context) {
	var input user_profile.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.Update(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Remove(ctx *net_lib.Context) {
	var input user_profile.Input

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
