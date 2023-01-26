package issue

import "go-issues-api/model"

type Usecase interface {
	GetAll() ([]*model.Issue, error)
	Create(userId int, title string, description string) error
}
