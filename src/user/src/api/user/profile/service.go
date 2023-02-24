package profile

import (
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/common/src/repository/common"
	"car-rent-platform/backend/common/src/repository/user-profile"
)

type (
	UserProfileInterface interface {
		FindAll(input *wql_lib.FilterInput) (output []*user_profile.Output, err error)
		FindOne(input *user_profile.Input) (output *user_profile.Output, err error)
		Create(input *user_profile.Input) (output *user_profile.Output, err error)
		Update(input *user_profile.Input) (output *user_profile.Output, err error)
		Remove(input *user_profile.Input) (output *user_profile.Output, err error)
	}
	UserProfileService struct {
		repo      *repository.Repository
		userModel *common.Repo[user_profile.Input, user_profile.Output]
	}
)

func NewUserService(repo *repository.Repository) UserProfileInterface {
	var i = UserProfileService{repo: repo}
	i.userModel = i.repo.Model("userProfile").(*common.Repo[user_profile.Input, user_profile.Output])
	return &i
}

func (s UserProfileService) FindAll(input *wql_lib.FilterInput) (output []*user_profile.Output, err error) {
	if res := s.userModel.Scan(input). /*.Joins("Profile") Where(input).*/ Find(&output); res.Error != nil {
		return nil, res.Error
	} else {
		return s.userModel.OutputListValidate(&output)
	}
	//return nil, err
}

func (s UserProfileService) FindOne(input *user_profile.Input) (output *user_profile.Output, err error) {
	if res := s.userModel.Db().Where(input).First(&output); res.Error != nil {
		return nil, res.Error
	} else {
		return s.userModel.OutputValidate(output)
	}
}

func (s UserProfileService) Create(input *user_profile.Input) (output *user_profile.Output, err error) {
	if res := s.userModel.Db().Create(input); res.Error != nil {
		return nil, res.Error
	} else {
		dest := res.Statement.Dest.(*user_profile.Input)
		return s.FindOne(dest)
	}
}

func (s UserProfileService) Update(input *user_profile.Input) (output *user_profile.Output, err error) {
	var where user_profile.Input
	where.UUID = input.UUID
	if res := s.userModel.Db().Where(&where).Updates(input).Where(&where).First(&output); res.Error != nil {
		return nil, res.Error
	}
	return output, nil
}

func (s UserProfileService) Remove(input *user_profile.Input) (output *user_profile.Output, err error) {
	var where user_profile.Input
	where.UUID = input.UUID
	if output, err = s.FindOne(&where); err != nil {
		return nil, err
	} else if dRes := s.userModel.Db().Delete(&where); dRes.Error != nil {
		return nil, err
	}
	return
}
