package issue

import "go-issues-api/domain/model"

type Usecase interface {
	GetAll() ([]*model.Issue, error)
	Create(userId int, title string, description string) error
}
