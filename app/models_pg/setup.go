package models_pg

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var PGDB *gorm.DB

type Tabler interface {
	TableName() string
}

func ConnectDatabase(DbUser, DbPassword, DbPort, DbHost, DbName string) {

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	database, err := gorm.Open("postgres", DBURL)

	if err != nil {
		panic("Failed to connect to postgres database at " + DBURL)
	}

	autoMigrate := os.Getenv("DB_AUTO_MIGRATE")
	autoPopulate := os.Getenv("DB_AUTO_POPULATE")

	if autoMigrate == "TRUE" {
		PostgresAutoMigrate(database)
	}

	PGDB = database

	if autoPopulate == "TRUE" {
		PostgresAutoPopulate()
	}
}
