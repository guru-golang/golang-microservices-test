package schema

import (
	"car-rent-platform/backend/common/src/repository/common"
)

type (
	User struct {
		common.Model[User]

		Email *string `gorm:"column:email;type:text;index;check:(email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$')" json:"email" validate:"required"`
		Phone *string `gorm:"column:phone;type:text;index" json:"phone" validate:"required"`
	}
)

func (User) TableName() string {
	return "users"
}
