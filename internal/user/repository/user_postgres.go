package repository

import (
	"go-issues-api/internal/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (this *UserRepository) Get(id int) (model.User, error) {
	var user model.User
	res := this.DB.First(&user, id)

	return user, res.Error
}

func (this *UserRepository) Create(user *model.User) error {
	res := this.DB.Create(user)
	return res.Error
}
