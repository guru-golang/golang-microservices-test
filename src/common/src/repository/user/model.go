package user

import (
	"car-rent-platform/backend/common/src/lib/gorm_lib"
	"car-rent-platform/backend/common/src/repository/common"
	schema2 "car-rent-platform/backend/common/src/repository/user-profile/schema"
	"car-rent-platform/backend/common/src/repository/user/schema"
)

type (
	Input struct {
		schema.User

		Profile *schema2.UserProfile `gorm:"foreignKey:userUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
	Output struct {
		schema.User

		Profile *schema2.UserProfile `gorm:"foreignKey:userUUID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}
)

func New(input Input, output Output, db *gorm_lib.Gorm) *common.Repo[Input, Output] {
	var repo common.Repo[Input, Output]

	tableName := "users"
	repo.Init(input, output, tableName, db)
	//repo.Db().AutoMigrate(&input)
	return &repo
}
