package schema

import (
	"car-rent-platform/backend/common/src/repository/common"
	"time"
)

type (
	UserProfile struct {
		common.Model[UserProfile]

		UserUUID  *string    `gorm:"column:userUUID;type:uuid;index;not null" json:"userUUID" validate:"required"`
		FirstName *string    `gorm:"column:firstName;type:text;index;not null" json:"firstName" validate:"required"`
		LastName  *string    `gorm:"column:lastName;type:text;index;not null" json:"lastName" validate:"required"`
		Gender    *string    `gorm:"column:gender;type:text;index;check:(gender IN ('male', 'female', 'other'))" json:"gender" validate:"omitempty"`
		BirthDate *time.Time `gorm:"column:birthDate;type:timestamp;index" json:"birthDate" validate:"required"`
	}
)

func (UserProfile) TableName() string {
	return "usersProfiles"
}
