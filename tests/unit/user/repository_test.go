package user

import (
	"go-issues-api/internal/model"
	"regexp"
	"testing"

	_userRepository "go-issues-api/internal/user/repository"

	"github.com/go-playground/assert/v2"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	tests := []struct {
		name          string
		mockDB        func(mock sqlmock.Sqlmock)
		expectedError error
		expectedLen   int
	}{
		{
			name: "Greate User Successfully",
			mockDB: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(
					regexp.QuoteMeta(
						`INSERT INTO "users" ("name") VALUES ($1) RETURNING "id"`)).
					WithArgs("Foo").
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Foo"))
				mock.ExpectCommit()
			},
			expectedError: nil,
			expectedLen:   1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			test.mockDB(mock)

			orm, _ := gorm.Open(postgres.New(postgres.Config{DriverName: "postgres", Conn: db}), &gorm.Config{})
			repo := _userRepository.NewUserRepository(orm)
			err = repo.Create(&model.User{Name: "Foo"})
			assert.Equal(t, test.expectedError, err)
		})
	}
}
