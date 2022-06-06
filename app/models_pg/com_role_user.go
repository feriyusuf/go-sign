package models_pg

import (
	"gorm.io/gorm"
	"time"
)

func (ComRoleUser) TableName() string {
	return "com_role_user"
}

// ComRoleUser Only one use has one role
type ComRoleUser struct {
	UserId   uint `gorm:"index;primaryKey;autoIncrement:false" json:"user_id"`
	RoleId   uint `gorm:"index;primaryKey;autoIncrement:false" json:"role_id"`
	IsActive bool `json:"is_active" gorm:"index;default:true;"`

	ComUser  ComUser `gorm:"foreignKey:UserId;"`
	ComeRole ComRole `gorm:"foreignKey:RoleId;"`

	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
	DeletedBy uint `json:"deleted_by"`
}
