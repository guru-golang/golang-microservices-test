package user_notification

import (
	"car-rent-platform/backend/common/src/lib/gorm_lib"
	"car-rent-platform/backend/common/src/repository/common"
	"car-rent-platform/backend/common/src/repository/user-notification/schema"
)

type (
	Input struct {
		schema.UserNotification
	}
	Output struct {
		schema.UserNotification
	}
)

func (Input) TableName() string {
	return "configurations"
}

func New(input Input, output Output, db *gorm_lib.Gorm) *common.Repo[Input, Output] {
	var repo common.Repo[Input, Output]
	tableName := "usersNotifications"
	repo.Init(input, output, tableName, db)
	//repo.Db().AutoMigrate(&input)
	return &repo
}
