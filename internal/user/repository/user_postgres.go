package repository

import (
	"go-issues-api/internal/model"

	"gorm.io/gorm"
)

type User struct {
	ID   uint   `gorm:"column:id;primary_key"`
	Name string `gorm:"column:name"`
}

func (u *User) toModel() *model.User {
	return &model.User{
		ID:   u.ID,
		Name: u.Name,
	}
}

func toGorm(u *model.User) *User {
	return &User{
		Name: u.Name,
	}
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(conn *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: conn,
	}
}

func (repo *UserRepository) Get(id int) (model.User, error) {
	var dbuser User
	res := repo.DB.First(&dbuser, id)

	return *dbuser.toModel(), res.Error
}

func (repo *UserRepository) Create(user *model.User) error {
	dbuser := toGorm(user)
	res := repo.DB.Create(dbuser)
	return res.Error
}
