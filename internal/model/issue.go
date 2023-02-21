package model

type Issue struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      User   `json:"author"`
}

type IssueRepository interface {
	FindAll() ([]*Issue, error)
	Create(issue *Issue) error
	FindBy(id int) (*Issue, error)
	Update(*Issue) (*Issue, error)
	Delete(id int) (int64, error)
}

type IssueUsecase interface {
	FindAll() ([]*Issue, error)
	Create(userId int, title string, description string) (*Issue, error)
	FindBy(id int) (*Issue, error)
	FindAndUpdate(id int, title, description string) (*Issue, error)
	DeleteBy(id int) (int64, error)
	Vote(id, userId, vote int) (*VoteIssue, error)
}
