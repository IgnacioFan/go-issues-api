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

func Connect(mode string) {
	postgresConfig := config.NewPostgresConfig(mode)
	DB, err = gorm.Open(postgres.Open(postgresConfig.Dsn()), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if mode != "test" {
		dbConfig := config.NewPostgresConfig("test")
		m := NewMigrate("database/migrations", dbConfig.Url())
		m.Up()
	}
}

func Disconnect(mode string) {
	db, _ := DB.DB()
	defer db.Close()

	if mode != "test" {
		dbConfig := config.NewPostgresConfig("test")
		m := NewMigrate("database/migrations", dbConfig.Url())
		m.Down()
	}
}
