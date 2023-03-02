package user

import (
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/user/src/api/common"
	"fmt"
	"github.com/gin-gonic/gin"
)

type (
	RpcInterface interface {
		common.FullRpcInterface
	}

	Rpc struct {
		service UserInterface
	}
)

func (rpc Rpc) Init(n *net_lib.Net, r *repository.Repository) {
	rpc.service = NewUserService(r)
	n.Pattern(net_lib.UserFindAll, rpc.FindAll)
	n.Pattern(net_lib.UserFindOne, rpc.FindOne)
	n.Pattern(net_lib.UserCreate, rpc.Create)
	n.Pattern(net_lib.UserUpdate, rpc.Update)
	n.Pattern(net_lib.UserRemove, rpc.Remove)
}

func (rpc Rpc) FindAll(ctx net_lib.Context) {
	var input wql_lib.FilterInput
	err := ctx.Msg.ShouldBind(&input)
	fmt.Println(err, input)

	if out, err := rpc.service.FindAll(&input); err != nil {
		_ = ctx.SendErr(err)
		//ctx.Write(msg.ToByte(gin.H{"status": false, "reason": err.Error()}))
	} else {
		_ = ctx.SendResp(gin.H{"result": out})
		//ctx.Write(msg.ToByte(gin.H{"status": true, "result": out}))
	}
}

func (rpc Rpc) FindOne(ctx net_lib.Context) {
	//TODO implement me
	panic("implement me")
}

func (rpc Rpc) Create(ctx net_lib.Context) {
	//TODO implement me
	panic("implement me")
}

func (rpc Rpc) Update(ctx net_lib.Context) {
	//TODO implement me
	panic("implement me")
}

func (rpc Rpc) Remove(ctx net_lib.Context) {
	//TODO implement me
	panic("implement me")
}

func NewRpc() RpcInterface {
	var i = Rpc{}
	return &i
}
