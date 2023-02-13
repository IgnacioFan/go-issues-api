package usercase

import (
	"go-issues-api/internal/model"
)

type UserUsecase struct {
	UserRepository model.UserRepository
}

func NewUserUsecase(userRepo model.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepo,
	}
}

func (this *UserUsecase) Create(name string) (*model.User, error) {
	user := &model.User{
		Name: name,
	}
	err := this.UserRepository.Create(user)
	return user, err
}
