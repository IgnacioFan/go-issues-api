package repository

import (
	"go-issues-api/internal/model"

	"gorm.io/gorm"
)

type IssueRepository struct {
	DB *gorm.DB
}

func NewIssueRepository(db *gorm.DB) *IssueRepository {
	return &IssueRepository{
		DB: db,
	}
}

func (this *IssueRepository) FindAll() ([]*model.Issue, error) {
	var issues []*model.Issue
	res := this.DB.Joins("Author").Find(&issues)

	return issues, res.Error
}

func (this *IssueRepository) Create(issue *model.Issue) error {
	res := this.DB.Create(issue)
	return res.Error
}

func (this *IssueRepository) FindBy(id int) (*model.Issue, error) {
	var issue *model.Issue
	res := this.DB.Joins("Author").First(&issue, id)

	return issue, res.Error
}

func (this *IssueRepository) Update(issue *model.Issue) (*model.Issue, error) {
	res := this.DB.Save(&issue)

	return issue, res.Error
}

func (this *IssueRepository) Delete(id int) (int64, error) {
	res := this.DB.Delete(&model.Issue{}, id)
	return res.RowsAffected, res.Error
}
