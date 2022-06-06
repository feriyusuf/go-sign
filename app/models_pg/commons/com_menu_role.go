package commons

import (
	"gorm.io/gorm"
	"time"
)

func (ComMenuRole) TableName() string {
	return "com_menu_role"
}

// ComMenuRole Only one role has one menu
type ComMenuRole struct {
	MenuId   uint `gorm:"index;primaryKey;autoIncrement:false" json:"menu_id"`
	RoleId   uint `gorm:"index;primaryKey;autoIncrement:false" json:"role_id"`
	Read     bool `json:"read_"`
	Write    bool `json:"write_"`
	IsActive bool `json:"is_active" gorm:"index;default:true;"`

	ComMenu  ComMenu `gorm:"foreignKey:MenuId;"`
	ComeRole ComRole `gorm:"foreignKey:RoleId;"`

	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
	DeletedBy uint `json:"deleted_by"`
}
