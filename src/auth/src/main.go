package main

import (
	"car-rent-platform/backend/auth/src/api"
	"car-rent-platform/backend/common/src/lib/config_lib"
	"car-rent-platform/backend/common/src/lib/gin_lib"
	"car-rent-platform/backend/common/src/lib/log_lib"
	"github.com/rs/zerolog/log"
)

func init() {
	config_lib.Init()
	log_lib.Init()
}

func main() {
	srv := gin_lib.NewServer().Init()
	api.NewAPI().Init(srv)

	if err := srv.Run(); err != nil {
		log.Fatal().Msg(err.Error())
	}

}
