package notification

import (
	"car-rent-platform/backend/common/src/lib/wql_lib"
	"car-rent-platform/backend/common/src/repository"
	"car-rent-platform/backend/common/src/repository/common"
	user_notification "car-rent-platform/backend/common/src/repository/user-notification"
)

type (
	UserNotificationInterface interface {
		FindAll(input *wql_lib.FilterInput) (output []*user_notification.Output, err error)
		FindOne(input *user_notification.Input) (output *user_notification.Output, err error)
		Create(input *user_notification.Input) (output *user_notification.Output, err error)
		Update(input *user_notification.Input) (output *user_notification.Output, err error)
		Remove(input *user_notification.Input) (output *user_notification.Output, err error)
	}
	UserNotificationService struct {
		repo      *repository.Repository
		userModel *common.Repo[user_notification.Input, user_notification.Output]
	}
)

func NewUserNotificationService(repo *repository.Repository) UserNotificationInterface {
	var i = UserNotificationService{repo: repo}
	i.userModel = i.repo.Model("userNotification").(*common.Repo[user_notification.Input, user_notification.Output])
	return &i
}

func (s UserNotificationService) FindAll(input *wql_lib.FilterInput) (output []*user_notification.Output, err error) {
	if res := s.userModel.Scan(input). /*.Joins("Profile") Where(input).*/ Find(&output); res.Error != nil {
		return nil, res.Error
	} else {
		return s.userModel.OutputListValidate(&output)
	}
	//return nil, err
}

func (s UserNotificationService) FindOne(input *user_notification.Input) (output *user_notification.Output, err error) {
	if res := s.userModel.Db().Where(input).First(&output); res.Error != nil {
		return nil, res.Error
	} else {
		return s.userModel.OutputValidate(output)
	}
}

func (s UserNotificationService) Create(input *user_notification.Input) (output *user_notification.Output, err error) {
	if res := s.userModel.Db().Create(input); res.Error != nil {
		return nil, res.Error
	} else {
		dest := res.Statement.Dest.(*user_notification.Input)
		return s.FindOne(dest)
	}
}

func (s UserNotificationService) Update(input *user_notification.Input) (output *user_notification.Output, err error) {
	var where user_notification.Input
	where.UUID = input.UUID
	if res := s.userModel.Db().Where(&where).Updates(input).Where(&where).First(&output); res.Error != nil {
		return nil, res.Error
	}
	return output, nil
}

func (s UserNotificationService) Remove(input *user_notification.Input) (output *user_notification.Output, err error) {
	var where user_notification.Input
	where.UUID = input.UUID
	if output, err = s.FindOne(&where); err != nil {
		return nil, err
	} else if dRes := s.userModel.Db().Delete(&where); dRes.Error != nil {
		return nil, err
	}
	return
}
