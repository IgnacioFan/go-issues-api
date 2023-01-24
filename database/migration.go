package database

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration struct {
	client *migrate.Migrate
}

func NewMigrate(filepath, url string) *Migration {
	m := &Migration{}
	m.client, err = migrate.New("file://"+filepath, url)

	if err != nil {
		log.Fatal(err)
	}
	return m
}

func (this *Migration) Up() {
	if err := this.client.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

func (this *Migration) Down() {
	if err := this.client.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
