package common

import (
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/repository"
)

type (
	RpcInterface interface {
		Init(n *net_lib.Net, r *repository.Repository)
	}

	FullRpcInterface interface {
		RpcInterface
		FindAll(ctx net_lib.Context)
		FindOne(ctx net_lib.Context)
		Create(ctx net_lib.Context)
		Update(ctx net_lib.Context)
		Remove(ctx net_lib.Context)
	}
)
