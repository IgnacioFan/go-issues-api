package repository

import (
	"go-issues-api/internal/model"
	_author "go-issues-api/internal/user/repository"

	"gorm.io/gorm"
)

type Issue struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
	AuthorID    uint
	Author      _author.User `gorm:"foreignKey:AuthorID"`
}

func (i *Issue) toModel() *model.Issue {
	return &model.Issue{
		ID:          i.ID,
		Title:       i.Title,
		Description: i.Description,
		Author:      model.User(i.Author),
	}
}

func toGorm(i *model.Issue) *Issue {
	return &Issue{
		Title:       i.Title,
		Description: i.Description,
		Author:      _author.User(i.Author),
	}
}

type IssueRepository struct {
	DB *gorm.DB
}

func NewIssueRepository(conn *gorm.DB) *IssueRepository {
	return &IssueRepository{
		DB: conn,
	}
}

func (this *IssueRepository) FindAll() ([]*model.Issue, error) {
	var dbissues []*Issue
	res := this.DB.Joins("Author").Find(&dbissues)

	issues := make([]*model.Issue, 0, len(dbissues))
	for _, dbi := range dbissues {
		issues = append(issues, &model.Issue{
			ID:          dbi.ID,
			Title:       dbi.Title,
			Description: dbi.Description,
			Author: model.User{
				ID:   dbi.Author.ID,
				Name: dbi.Author.Name,
			},
		})
	}

	return issues, res.Error
}

func (this *IssueRepository) Create(issue *model.Issue) error {
	dbissue := toGorm(issue)
	res := this.DB.Create(dbissue)

	return res.Error
}

func (this *IssueRepository) FindBy(id int) (*model.Issue, error) {
	var dbissue *Issue
	res := this.DB.Joins("Author").First(&dbissue, id)

	return dbissue.toModel(), res.Error
}

// TODO: doesn't really update
func (this *IssueRepository) Update(issue *model.Issue) (*model.Issue, error) {
	dbissue := toGorm(issue)
	res := this.DB.Save(&dbissue)

	return dbissue.toModel(), res.Error
}

func (this *IssueRepository) Delete(id int) (int64, error) {
	var dbissue Issue
	res := this.DB.Delete(&dbissue, id)

	return res.RowsAffected, res.Error
}
