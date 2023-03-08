package schema

import (
	"car-rent-platform/backend/common/src/repository/common"
)

type (
	UserNotification struct {
		common.Model[UserNotification]

		UserUuid    string `gorm:"column:userUuid; type:uuid;index" json:"userUuid" validate:"omitempty"`
		AppUuid     string `gorm:"column:appUuid; type:uuid;index" json:"appUuid" validate:"omitempty"`
		Body        string `gorm:"column:body; type:json" json:"body" validate:"omitempty"`
		InAppStatus string `gorm:"column:inAppStatus; type:text;index" json:"inAppStatus" validate:"omitempty"`
		PushStatus  string `gorm:"column:pushStatus; type:text;index" json:"pushStatus" validate:"omitempty"`
	}
)

func (UserNotification) TableName() string {
	return "usersNotifications"
}
