package user

import (
	"car-rent-platform/backend/common/src/lib/gorm_lib"
	"car-rent-platform/backend/common/src/repository/model"
)

type (
	Input struct {
		model.Model[Input]
	}
	Output struct {
		model.Model[Output]
	}
)

func New(input Input, output Output, db *gorm_lib.Gorm) *model.Repo[Input, Output] {
	var repo model.Repo[Input, Output]
	repo.Init(input, output, "users", db)
	//repo.Db().AutoMigrate(&input)
	return &repo
}
