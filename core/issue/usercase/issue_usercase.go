package usercase

import (
	_issueRepository "go-issues-api/core/issue"
	"go-issues-api/core/model"
	_userRepository "go-issues-api/core/user"
)

type IssueUsercase struct {
	userRepository  _userRepository.Repository
	IssueRepository _issueRepository.Repository
}

func NewIssueUsercase(user _userRepository.Repository, issue _issueRepository.Repository) *IssueUsercase {
	return &IssueUsercase{
		userRepository:  user,
		IssueRepository: issue,
	}
}

func (this *IssueUsercase) GetAll() ([]*model.Issue, error) {
	return this.IssueRepository.GetAll()
}

func (this *IssueUsercase) Create(userId int, title string, description string) error {
	var err error

	user, err := this.userRepository.Get(userId)
	if err != nil {
		panic(err)
	}

	var issue = &model.Issue{
		Title:       title,
		Description: description,
		Author:      user,
	}
	err = this.IssueRepository.Create(issue)

	return err
}

func (this *IssueUsercase) FindBy(id int) (*model.Issue, error) {
	issue, err := this.IssueRepository.FindBy(id)
	return issue, err
}

func (this *IssueUsercase) FindAndUpdate(id int, title, description string) (*model.Issue, error) {
	var err error
	issue, err := this.IssueRepository.FindBy(id)

	if err != nil {
		return issue, err
	}

	issue.Title = title
	issue.Description = description
	issue, err = this.IssueRepository.Update(issue)

	return issue, err
}

func (this *IssueUsercase) DeleteBy(id int) (int64, error) {
	var err error
	issue, err := this.IssueRepository.FindBy(id)

	if err != nil {
		return int64(0), err
	}

	affected, err := this.IssueRepository.Delete(int(issue.ID))

	return affected, err
}
