package main

import (
	"car-rent-platform/backend/common/src/lib/config_lib"
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/lib/gorm_lib"
	"car-rent-platform/backend/common/src/lib/log_lib"
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/payment/src/api"
	"github.com/rs/zerolog/log"
	"os"
	"sync"
)

func init() {
	_ = os.Setenv("app_name", "payment")
	config_lib.Init()
	log_lib.Init()
}

func main() {
	db := gorm_lib.NewConn().Init()
	repo := repository.NewRepository().Init(db)
	srv := gin_lib.NewServer().Init()
	nSrv := net_lib.NewServer().Init()
	api.NewAPI().InitRpc(nSrv, repo)
	api.NewAPI().InitRoute(srv, repo)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := srv.Run(); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()
	go func() {
		defer wg.Done()
		if err := nSrv.Run(); err != nil {
			log.Fatal().Msg(err.Error())
		}
	}()
	wg.Wait()
}
