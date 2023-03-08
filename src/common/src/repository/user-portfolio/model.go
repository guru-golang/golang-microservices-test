package user_portfolio

import (
	"car-rent-platform/backend/common/src/lib/gorm_lib"
	"car-rent-platform/backend/common/src/repository/common"
	"car-rent-platform/backend/common/src/repository/user-portfolio/schema"
)

type (
	Input struct {
		schema.UserPortfolio
	}
	Output struct {
		schema.UserPortfolio
	}
)

func New(input Input, output Output, db *gorm_lib.Gorm) *common.Repo[Input, Output] {
	var repo common.Repo[Input, Output]

	tableName := "usersPortfolios"
	repo.Init(input, output, tableName, db)
	//repo.Db().AutoMigrate(&input)
	return &repo
}
