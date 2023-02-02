package user

import "go-issues-api/internal/model"

type Usecase interface {
	Create(name string) (*model.User, error)
}
