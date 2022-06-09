package commons

import (
	"database/sql/driver"
	"gorm.io/gorm"
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
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" gorm:"index;"`
	Label    string `json:"label"`
	Icon     string `json:"icon"`
	IsActive bool   `json:"is_active" gorm:"index;default:true;"`

	ParentID  uint           `json:"parent_id" gorm:"index"`
	Children  []ComMenu      `gorm:"foreignkey:ParentID"`
	Type      Type           `json:"type_" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	UiPath  string `json:"ui_path"`
	ApiPath string `json:"api_path"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedBy uint      `json:"created_by"`
	UpdatedBy uint      `json:"updated_by"`
	DeletedBy uint      `json:"deleted_by"`
}
