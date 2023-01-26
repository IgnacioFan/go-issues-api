package repository

import (
	"go-issues-api/domain/model"

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

func (this *IssueRepository) GetAll() ([]*model.Issue, error) {
	var issues []*model.Issue
	// res := this.DB.Find(issues)
	res := this.DB.Joins("Author").Find(&issues)

	return issues, res.Error
}

func (this *IssueRepository) Create(issue *model.Issue) error {
	res := this.DB.Create(issue)

	return res.Error
}

// func Find(id int) (model.Issue, error) {
// 	var issue model.Issue
// 	res := database.DB.First(&issue, id)

// 	return issue, res.Error
// }

// func FindAndUpdate(id int, title, description string) (model.Issue, error) {
// 	issue, err := Find(id)

// 	if err != nil {
// 		return issue, err
// 	}

// 	issue.Title = title
// 	issue.Description = description
// 	res := database.DB.Save(&issue)

// 	return issue, res.Error
// }

// func Delete(id int) (int64, error) {
// 	res := database.DB.Delete(&model.Issue{}, id)

// 	return res.RowsAffected, res.Error
// }
