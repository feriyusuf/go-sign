package models_pg

import "time"

func (ComRole) TableName() string {
	return "com_role"
}

type ComRole struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"index"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedBy uint      `json:"created_by"`
	UpdatedBy uint      `json:"updated_by"`
}
