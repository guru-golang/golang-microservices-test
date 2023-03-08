package main

import (
	"car-rent-platform/backend/common/src/lib/config_lib"
	"car-rent-platform/backend/common/src/lib/log_lib"
	"car-rent-platform/backend/proxy/src/tasks"
	"os"
)

func init() {
	_ = os.Setenv("app_name", "user")
	config_lib.Init()
	log_lib.Init()
}

func main() {
	tasks.Init()
}
