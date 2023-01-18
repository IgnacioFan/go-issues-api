package database

import (
	"fmt"
	"go-issues-api/config"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		log.Fatal(err)
	}

	m := Migrate(mode)
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func Disconnect(mode string) {
	db, _ := DB.DB()
	defer db.Close()

	m := Migrate(mode)
	if err := m.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
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

func Migrate(mode string) *migrate.Migrate {
	dbConfig := config.PostgresConn(mode)
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Dbname,
		dbConfig.Ssl,
	)

	m, err := migrate.New(config.MigrationsPath(), dbUrl)

	if err != nil {
		log.Fatal(err)
	}
	return m
}
