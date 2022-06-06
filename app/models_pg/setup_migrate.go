package models_pg

import "gorm.io/gorm"

func PostgresAutoMigrate(database *gorm.DB) {
	database.AutoMigrate(
		&ComMenu{},
		&ComMenuRole{},
		&ComRole{},
		&ComRoleUser{},
		&ComUser{},
	)
}
