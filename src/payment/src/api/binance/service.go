package binance

import (
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/common/src/repository/common"
	"car-rent-platform/backend/common/src/repository/configuration"
)

type (
	BinanceInterface interface {
		FindAll(input *wql_lib.FilterInput) (output []*configuration.Output, err error)
		FindOne(input *configuration.Input) (output *configuration.Output, err error)
		Create(input *configuration.Input) (output *configuration.Output, err error)
		Update(input *configuration.Input) (output *configuration.Output, err error)
		Remove(input *configuration.Input) (output *configuration.Output, err error)
	}
	BinanceService struct {
		repo      *repository.Repository
		userModel *common.Repo[configuration.Input, configuration.Output]
	}
)

func NewBinanceService(repo *repository.Repository) BinanceInterface {
	var i = BinanceService{repo: repo}
	i.userModel = i.repo.Model("configuration").(*common.Repo[configuration.Input, configuration.Output])
	return &i
}

func (s *BinanceService) FindAll(input *wql_lib.FilterInput) (output []*configuration.Output, err error) {
	if res := s.userModel.Scan(input).Find(&output); res.Error != nil {
		return nil, res.Error
	} else {
		return s.userModel.OutputListValidate(&output)
	}
}

func (s *BinanceService) FindOne(input *configuration.Input) (output *configuration.Output, err error) {
	if res := s.userModel.Db().Where(input).First(&output); res.Error != nil {
		return nil, res.Error
	} else {
		return s.userModel.OutputValidate(output)
	}
}

func (s *BinanceService) Create(input *configuration.Input) (output *configuration.Output, err error) {
	if res := s.userModel.Db().Create(input); res.Error != nil {
		return nil, res.Error
	} else {
		dest := res.Statement.Dest.(*configuration.Input)
		return s.FindOne(dest)
	}
}

func (s *BinanceService) Update(input *configuration.Input) (output *configuration.Output, err error) {
	var where configuration.Input
	where.UUID = input.UUID
	if res := s.userModel.Db().Where(&where).Updates(input).Where(&where).First(&output); res.Error != nil {
		return nil, res.Error
	}
	return output, nil
}

func (s *BinanceService) Remove(input *configuration.Input) (output *configuration.Output, err error) {
	var where configuration.Input
	where.UUID = input.UUID
	if output, err = s.FindOne(&where); err != nil {
		return nil, err
	} else if dRes := s.userModel.Db().Delete(&where); dRes.Error != nil {
		return nil, err
	}
	return
}
