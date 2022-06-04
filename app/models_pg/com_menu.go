package models_pg

import (
	"database/sql/driver"
	"time"
)

type Type string

const (
	SCREEN Type = "SCREEN"
	ACTION Type = "ACTION"
)

func (e *Type) Scan(value interface{}) error {
	*e = Type(value.([]byte))
	return nil
}

func (e Type) Value() (driver.Value, error) {
	return string(e), nil
}

func (ComMenu) TableName() string {
	return "com_menu"
}

type ComMenu struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Icon string `json:"icon"`

	ParentID uint      `json:"parent_id" gorm:"index"`
	Children []ComMenu `gorm:"foreignkey:ParentID"`
	Type     Type      `json:"type_" gorm:"index"`

	UiPath  string `json:"ui_path"`
	ApiPath string `json:"api_path"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedBy uint      `json:"created_by"`
	UpdatedBy uint      `json:"updated_by"`
}
