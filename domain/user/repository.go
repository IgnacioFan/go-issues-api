package user

import "go-issues-api/domain/model"

type Repository interface {
	Get(id int) (model.User, error)
}
