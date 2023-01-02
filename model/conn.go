package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db
	migration()
}

func TearDownDatabase() {
	db, _ := DB.DB()
	defer db.Close()
	DB.Migrator().DropTable(&Issue{})
}
