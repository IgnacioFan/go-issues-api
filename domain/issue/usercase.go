package issue

import "go-issues-api/domain/model"

type Usecase interface {
	GetAll() ([]*model.Issue, error)
	Create(userId int, title string, description string) error
	FindBy(id int) (*model.Issue, error)
	FindAndUpdate(id int, title, description string) (*model.Issue, error)
	DeleteBy(id int) (int64, error)
}
