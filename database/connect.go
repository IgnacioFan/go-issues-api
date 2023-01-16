package database

import (
	"fmt"
	"go-issues-api/config"
	"go-issues-api/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect(mode string) {
	DB, err = gorm.Open(postgres.Open(dsn(mode)), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	fmt.Println("connect to database")

	err := DB.AutoMigrate(&model.Issue{}) // TODO: consider using goose

	if err != nil {
		panic("failed to run data migration")
	}
}

func Disconnect() {
	db, _ := DB.DB()
	defer db.Close()
	DB.Migrator().DropTable(&model.Issue{}) // TODO: consider using goose
}

func dsn(mode string) string {
	conn := config.PostgresConn(mode)

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		conn.Host,
		conn.User,
		conn.Password,
		conn.Dbname,
		conn.Port,
		conn.Ssl,
		conn.Timezone,
	)
}
