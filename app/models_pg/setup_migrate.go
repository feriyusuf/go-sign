package models_pg

import "github.com/jinzhu/gorm"

func PostgresAutoMigrate(database *gorm.DB) {
	database.AutoMigrate(
		&ComMenu{},
		&ComMenuRole{},
		&ComRole{},
		&ComRoleUser{},
		&ComUser{},
	)
}
