package commons

import (
	"gorm.io/gorm"
	"time"
)

func (ComRole) TableName() string {
	return "com_role"
}

type ComRole struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"index"`
	IsActive bool   `json:"is_active" gorm:"index;default:true;"`

	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
	DeletedBy uint `json:"deleted_by"`
}
