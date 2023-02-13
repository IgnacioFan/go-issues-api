package model

type Issue struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      User   `gorm:"foreignKey:ID" json:"author"`
}

type IssueRepository interface {
	FindAll() ([]*Issue, error)
	Create(issue *Issue) error
	FindBy(id int) (*Issue, error)
	Update(*Issue) (*Issue, error)
	Delete(id int) (int64, error)
}
