package issue

import "go-issues-api/domain/model"

type Repository interface {
	GetAll() ([]*model.Issue, error)
	Create(issue *model.Issue) error
}
