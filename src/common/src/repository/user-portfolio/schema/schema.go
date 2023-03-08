package schema

import (
	"car-rent-platform/backend/common/src/repository/common"
)

type (
	UserPortfolio struct {
		common.Model[UserPortfolio]

		UserUUID *string `gorm:"column:userUUID; type:uuid; index; not null" json:"userUUID" validate:"required"`
		Name     *string `gorm:"column:name; type:text; index; not null" json:"name" validate:"required"`
		Type     *string `gorm:"column:type; type:text; index; not null" json:"type" validate:"required"`
		Secrets  *string `gorm:"column:secrets; type:text; index; not null" json:"secrets" validate:"required"`
		Status   *string `gorm:"column:status; type:text; index; not null" json:"status" validate:"required"`
		Source   *string `gorm:"column:source; type:text; index; not null" json:"source" validate:"required"`
		Account  *string `gorm:"column:account; type:json; not null" json:"account" validate:"required"`
		Assets   *string `gorm:"column:assets; type:json; not null" json:"assets" validate:"required"`
	}
)

func (UserPortfolio) TableName() string {
	return "usersProfiles"
}
