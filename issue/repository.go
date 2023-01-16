package issue

import (
	"go-issues-api/database"
	"go-issues-api/model"
)

func Seed() {
	var seedIssues = []model.Issue{
		{
			Title:       "issue 1",
			Description: "This is issue 1",
		},
		{
			Title:       "issue 2",
			Description: "This is issue 2",
		},
	}
	database.DB.Create(&seedIssues)
}

func Create(title, description string) (model.Issue, error) {
	issue := &model.Issue{
		Title:       title,
		Description: description,
	}

	res := database.DB.Create(issue)

	return *issue, res.Error
}

func Find(id int) (model.Issue, error) {
	var issue model.Issue
	res := database.DB.First(&issue, id)

	return issue, res.Error
}

func FindAll() ([]model.Issue, error) {
	var issues []model.Issue
	res := database.DB.Find(&issues)

	return issues, res.Error
}

func FindAndUpdate(id int, title, description string) (model.Issue, error) {
	issue, err := Find(id)

	if err != nil {
		return issue, err
	}

	issue.Title = title
	issue.Description = description
	res := database.DB.Save(&issue)

	return issue, res.Error
}

func Delete(id int) (int64, error) {
	res := database.DB.Delete(&model.Issue{}, id)

	return res.RowsAffected, res.Error
}
