package user

import (
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/common/src/repository/common"
	"car-rent-platform/backend/common/src/repository/user"
)

type (
	UserInterface interface {
		FindAll(input *wql_lib.FilterInput) (output []*user.Output, err error)
		FindOne(input *user.Input) (output *user.Output, err error)
		Create(input *user.Input) (output *user.Output, err error)
		Update(input *user.Input) (output *user.Output, err error)
		Remove(input *user.Input) (output *user.Output, err error)
	}
	UserService struct {
		repo      *repository.Repository
		userModel *common.Repo[user.Input, user.Output]
	}
)

func NewUserService(repo *repository.Repository) UserInterface {
	var i = UserService{repo: repo}
	i.userModel = i.repo.Model("user").(*common.Repo[user.Input, user.Output])
	return &i
}

func (s UserService) FindAll(input *wql_lib.FilterInput) (output []*user.Output, err error) {
	if res := s.userModel.Scan(input). /*.Joins("Profile") Where(input).*/ Find(&output); res.Error != nil {
		return nil, res.Error
	} else {
		return s.userModel.OutputListValidate(&output)
	}
	//return nil, err
}

func (s UserService) FindOne(input *user.Input) (output *user.Output, err error) {
	if res := s.userModel.Db().Where(input).First(&output); res.Error != nil {
		return nil, res.Error
	} else {
		return s.userModel.OutputValidate(output)
	}
}

func (s UserService) Create(input *user.Input) (output *user.Output, err error) {
	if res := s.userModel.Db().Create(input); res.Error != nil {
		return nil, res.Error
	} else {
		dest := res.Statement.Dest.(*user.Input)
		return s.FindOne(dest)
	}
}

func (s UserService) Update(input *user.Input) (output *user.Output, err error) {
	var where user.Input
	where.UUID = input.UUID
	if res := s.userModel.Db().Where(&where).Updates(input).Where(&where).First(&output); res.Error != nil {
		return nil, res.Error
	}
	return output, nil
}

func (s UserService) Remove(input *user.Input) (output *user.Output, err error) {
	var where user.Input
	where.UUID = input.UUID
	if output, err = s.FindOne(&where); err != nil {
		return nil, err
	} else if dRes := s.userModel.Db().Delete(&where); dRes.Error != nil {
		return nil, err
	}
	return
}
