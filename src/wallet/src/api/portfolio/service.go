package portfolio

import (
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/common/src/repository/common"
	"car-rent-platform/backend/common/src/repository/user-portfolio"
)

type (
	PortfolioInterface interface {
		FindAll(input *wql_lib.FilterInput) (output []*user_portfolio.Output, err error)
		FindOne(input *user_portfolio.Input) (output *user_portfolio.Output, err error)
		Create(input *user_portfolio.Input) (output *user_portfolio.Output, err error)
		Update(input *user_portfolio.Input) (output *user_portfolio.Output, err error)
		Remove(input *user_portfolio.Input) (output *user_portfolio.Output, err error)
	}
	PortfolioService struct {
		repo      *repository.Repository
		userModel *common.Repo[user_portfolio.Input, user_portfolio.Output]
	}
)

func NewPortfolioService(repo *repository.Repository) PortfolioInterface {
	var i = PortfolioService{repo: repo}
	i.userModel = i.repo.Model("userPortfolio").(*common.Repo[user_portfolio.Input, user_portfolio.Output])
	return &i
}

func (s *PortfolioService) FindAll(input *wql_lib.FilterInput) (output []*user_portfolio.Output, err error) {
	if res := s.userModel.Scan(input).Find(&output); res.Error != nil {
		return nil, res.Error
	} else {
		return s.userModel.OutputListValidate(&output)
	}
}

func (s *PortfolioService) FindOne(input *user_portfolio.Input) (output *user_portfolio.Output, err error) {
	if res := s.userModel.Db().Where(input).First(&output); res.Error != nil {
		return nil, res.Error
	} else {
		return s.userModel.OutputValidate(output)
	}
}

func (s *PortfolioService) Create(input *user_portfolio.Input) (output *user_portfolio.Output, err error) {
	if res := s.userModel.Db().Create(input); res.Error != nil {
		return nil, res.Error
	} else {
		dest := res.Statement.Dest.(*user_portfolio.Input)
		return s.FindOne(dest)
	}
}

func (s *PortfolioService) Update(input *user_portfolio.Input) (output *user_portfolio.Output, err error) {
	var where user_portfolio.Input
	where.UUID = input.UUID
	if res := s.userModel.Db().Where(&where).Updates(input).Where(&where).First(&output); res.Error != nil {
		return nil, res.Error
	}
	return output, nil
}

func (s *PortfolioService) Remove(input *user_portfolio.Input) (output *user_portfolio.Output, err error) {
	var where user_portfolio.Input
	where.UUID = input.UUID
	if output, err = s.FindOne(&where); err != nil {
		return nil, err
	} else if dRes := s.userModel.Db().Delete(&where); dRes.Error != nil {
		return nil, err
	}
	return
}
