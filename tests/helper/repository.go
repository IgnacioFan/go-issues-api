package helper

import (
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupGormMock(t *testing.T) (*gorm.DB, *sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to setup sqlmock: %s", err)
	}

	orm, err := gorm.Open(postgres.New(postgres.Config{DriverName: "postgres", Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to use gorm to open DB connection: %s", err)
	}
	return orm, &mock
}
