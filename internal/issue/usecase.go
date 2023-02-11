package issue

import "go-issues-api/internal/model"

type Usecase interface {
	FindAll() ([]*model.Issue, error)
	Create(userId int, title string, description string) (*model.Issue, error)
	FindBy(id int) (*model.Issue, error)
	FindAndUpdate(id int, title, description string) (*model.Issue, error)
	DeleteBy(id int) (int64, error)
}
