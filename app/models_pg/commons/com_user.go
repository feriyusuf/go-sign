package commons

import "time"

func (ComUser) TableName() string {
	return "com_user"
}

type ComUser struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Username string `json:"username" gorm:"index,unique"`
	Password string `json:"password"`

	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	CreatedBy uint `json:"created_by"`
	UpdatedBy uint `json:"updated_by"`
}
