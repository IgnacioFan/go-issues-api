package issue

import (
	"go-issues-api/model"
)

func Create(title, description string) (model.Issue, error) {
	issue := &model.Issue{
		Title:       title,
		Description: description,
	}

	res := model.DB.Create(issue)

	return *issue, res.Error
}

func Find(id int) (model.Issue, error) {
	var issue model.Issue
	res := model.DB.First(&issue, id)

	return issue, res.Error
}

func FindAll() ([]model.Issue, error) {
	var issues []model.Issue
	res := model.DB.Find(&issues)

	return issues, res.Error
}

func FindAndUpdate(id int, title, description string) (model.Issue, error) {
	issue, err := Find(id)

	if err != nil {
		return issue, err
	}

	issue.Title = title
	issue.Description = description
	res := model.DB.Save(&issue)

	return issue, res.Error
}

func Delete(id int) (int64, error) {
	res := model.DB.Delete(&model.Issue{}, id)

	return res.RowsAffected, res.Error
}
