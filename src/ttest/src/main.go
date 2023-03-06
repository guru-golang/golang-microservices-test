package main

import (
	"car-rent-platform/backend/common/src/lib/config_lib"
	"car-rent-platform/backend/common/src/lib/log_lib"
	"car-rent-platform/backend/common/src/lib/net_lib"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func init() {
	_ = os.Setenv("app_name", "ttest")
	config_lib.Init()
	log_lib.Init()
}

func main() {
	service := net_lib.New().Service("user")
	for i := 0; i < 1; i++ {
		send, err := service.Emit(net_lib.UserFindAll, gin.H{"filterMeta": gin.H{"email": "asdasd@asd.asd"}})
		if err != nil {
			fmt.Println(i, err)
			return
		}
		fmt.Println(i, send)
	}
	if err := service.Release(); err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 1; i++ {
		send, err := service.Emit(net_lib.UserFindOne, gin.H{"uuid": "0513a061-8b1b-48ec-aa42-172ea9abfde4"})
		if err != nil {
			fmt.Println(i, err)
			return
		}
		fmt.Println(i, send)
	}
	if err := service.Release(); err != nil {
		fmt.Println(err)
	}
}
