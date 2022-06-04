package models_pg

import "time"

func (ComUserRole) TableName() string {
	return "com_user_role"
}

type ComUserRole struct {
	UserId   uint    `json:"user_id"`
	RoleId   uint    `json:"role_id"`
	ComUser  ComUser `gorm:"foreignKey:UserId;"`
	ComeRole ComRole `gorm:"foreignKey:RoleId;"`

	CreatedBy uint      `json:"created_by"`
	UpdatedBy uint      `json:"updated_by"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
