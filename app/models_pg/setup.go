package models_pg

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var PGDB *gorm.DB

type Tabler interface {
	TableName() string
}

func ConnectDatabase() {
	sqlTimeZone := os.Getenv("TIMEZONE")
	sqlUser := os.Getenv("DB_USER_PG")
	sqlPassword := os.Getenv("DB_PASSWORD_PG")
	sqlPort := os.Getenv("DB_PORT_PG")
	sqlDbName := os.Getenv("DB_NAME_PG")
	sqlHost := os.Getenv("DB_HOST_PG")
	autoMigrate := os.Getenv("DB_AUTO_MIGRATE")
	autoPopulate := os.Getenv("DB_AUTO_POPULATE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s TimeZone=Asia/Jakarta", sqlHost, sqlPort, sqlUser, sqlDbName, sqlPassword /*, sqlTimeZone*/)
	database, err := gorm.Open(postgres.New(
		postgres.Config{DSN: dsn, PreferSimpleProtocol: true}),
		&gorm.Config{
			NowFunc: func() time.Time {
				ti, _ := time.LoadLocation(sqlTimeZone)
				return time.Now().In(ti)
			},
		},
	)

	if err != nil {
		panic("Failed to connect to postgres database at " + dsn)
	}

	if err == nil {
		log.Printf("Application connected to %v", dsn)
	}

	if autoMigrate == "TRUE" {
		PostgresAutoMigrate(database)
	}

	PGDB = database

	if autoPopulate == "TRUE" {
		PostgresAutoPopulate()
	}
}
