package user

import "go-issues-api/model"

type Repository interface {
	Get(id int) (model.User, error)
}
