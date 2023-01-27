package user

import "go-issues-api/core/model"

type Repository interface {
	Get(id int) (model.User, error)
	Create(user *model.User) error
}
