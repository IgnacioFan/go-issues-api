package issue

import "go-issues-api/domain/model"

type Repository interface {
	GetAll() ([]*model.Issue, error)
	Create(issue *model.Issue) error
	FindBy(id int) (*model.Issue, error)
	Update(*model.Issue) (*model.Issue, error)
	Delete(id int) (int64, error)
}
