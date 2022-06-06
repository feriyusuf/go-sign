package models_pg

import (
	"github.com/feriyusuf/go-sign/app/models_pg/commons"
	"gorm.io/gorm"
)

func PostgresAutoMigrate(database *gorm.DB) {
	database.AutoMigrate(
		&commons.ComMenu{},
		&commons.ComMenuRole{},
		&commons.ComRole{},
		&commons.ComRoleUser{},
		&commons.ComUser{},
	)
}
