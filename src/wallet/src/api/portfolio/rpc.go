package portfolio

import (
	common "car-rent-platform/backend/common/src/context"
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/common/src/repository/user-portfolio"
	"github.com/gin-gonic/gin"
)

type (
	RpcInterface interface {
		common.FullRpcInterface
	}

	Rpc struct {
		service PortfolioInterface
	}
)

func (rpc *Rpc) Init(n *net_lib.Net, r *repository.Repository) {
	rpc.service = NewPortfolioService(r)
	n.Pattern(net_lib.PortfolioFindAll, rpc.FindAll)
	n.Pattern(net_lib.PortfolioFindOne, rpc.FindOne)
	n.Pattern(net_lib.PortfolioCreate, rpc.Create)
	n.Pattern(net_lib.PortfolioUpdate, rpc.Update)
	n.Pattern(net_lib.PortfolioRemove, rpc.Remove)
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
	var input user_portfolio.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.FindOne(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Create(ctx *net_lib.Context) {
	var input user_portfolio.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.Create(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Update(ctx *net_lib.Context) {
	var input user_portfolio.Input

	if err := ctx.Msg.ShouldBind(&input); err != nil {
		_ = ctx.Error(err)
	} else if out, err := rpc.service.Update(&input); err != nil {
		_ = ctx.Error(err)
	} else {
		_ = ctx.Response(gin.H{"result": out})
	}
}

func (rpc *Rpc) Remove(ctx *net_lib.Context) {
	var input user_portfolio.Input

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
