package database

import (
	"go-issues-api/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect(mode string) *gorm.DB {
	postgresConfig := config.NewPostgresConfig(mode)
	DB, err = gorm.Open(postgres.Open(postgresConfig.Dsn()), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if mode == "dev" {
		dbConfig := config.NewPostgresConfig("dev")
		m := NewMigrate("../database/migrations", dbConfig.Url())
		m.Up()
	}

	return DB
}

func Disconnect(mode string) {
	db, _ := DB.DB()
	defer db.Close()

	if mode == "dev" {
		dbConfig := config.NewPostgresConfig("dev")
		m := NewMigrate("../database/migrations", dbConfig.Url())
		m.Down()
	}
}
