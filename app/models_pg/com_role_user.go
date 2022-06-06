package models_pg

import "time"

func (ComRoleUser) TableName() string {
	return "com_role_user"
}

// ComRoleUser Only one use has one role
type ComRoleUser struct {
	UserId uint `gorm:"index;primaryKey;autoIncrement:false" json:"user_id"`
	RoleId uint `gorm:"index;primaryKey;autoIncrement:false" json:"role_id"`

	ComUser  ComUser `gorm:"foreignKey:UserId;"`
	ComeRole ComRole `gorm:"foreignKey:RoleId;"`

	CreatedBy uint      `json:"created_by"`
	UpdatedBy uint      `json:"updated_by"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
