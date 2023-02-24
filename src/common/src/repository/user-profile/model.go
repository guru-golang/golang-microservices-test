package user_profile

import (
	"car-rent-platform/backend/common/src/lib/gorm_lib"
	"car-rent-platform/backend/common/src/repository/common"
	"car-rent-platform/backend/common/src/repository/user-profile/schema"
)

type (
	Input struct {
		schema.UserProfile
	}
	Output struct {
		schema.UserProfile
	}
)

func New(input Input, output Output, db *gorm_lib.Gorm) *common.Repo[Input, Output] {
	var repo common.Repo[Input, Output]

	tableName := "usersProfiles"
	repo.Init(input, output, tableName, db)
	//repo.Db().AutoMigrate(&input)
	return &repo
}
