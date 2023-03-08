package notification

import (
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	userNotification "car-rent-platform/backend/common/src/repository/user-notification"
	"car-rent-platform/backend/user/src/api/common"
	"github.com/gin-gonic/gin"
)

type (
	RpcInterface interface {
		common.FullRpcInterface
	}

	Rpc struct {
		service UserNotificationInterface
	}
)

func (rpc *Rpc) Init(n *net_lib.Net, r *repository.Repository) {
	rpc.service = NewUserNotificationService(r)
	n.Pattern(net_lib.UserNotificationFindAll, rpc.FindAll)
	n.Pattern(net_lib.UserNotificationFindOne, rpc.FindOne)
	n.Pattern(net_lib.UserNotificationCreate, rpc.Create)
	n.Pattern(net_lib.UserNotificationUpdate, rpc.Update)
	n.Pattern(net_lib.UserNotificationRemove, rpc.Remove)
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
	var input userNotification.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.FindOne(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Create(ctx *net_lib.Context) {
	var input userNotification.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.Create(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Update(ctx *net_lib.Context) {
	var input userNotification.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.Update(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Remove(ctx *net_lib.Context) {
	var input userNotification.Input

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
