package user

import "go-issues-api/internal/model"

type Repository interface {
	Get(id int) (model.User, error)
	Create(user *model.User) error
}
