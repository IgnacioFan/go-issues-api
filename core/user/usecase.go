package user

import "go-issues-api/core/model"

type Usecase interface {
	Create(name string) (*model.User, error)
}
