package main

import (
	"car-rent-platform/backend/common/src/lib/config_lib"
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/lib/gorm_lib"
	"car-rent-platform/backend/common/src/lib/log_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/user/src/api"
	"github.com/rs/zerolog/log"
)

func init() {
	config_lib.Init()
	log_lib.Init()
}

func main() {
	db := gorm_lib.NewConn().Init()
	repo := repository.NewRepository().Init(db)
	srv := gin_lib.NewServer().Init()
	api.NewAPI().Init(srv, repo)

	if err := srv.Run(); err != nil {
		log.Fatal().Msg(err.Error())
	}

}
