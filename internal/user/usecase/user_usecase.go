package usercase

import (
	"go-issues-api/internal/model"
	_userRepository "go-issues-api/internal/user"
)

type UserUsecase struct {
	UserRepository _userRepository.Repository
}

func NewUserUsecase(userRepo _userRepository.Repository) *UserUsecase {
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
