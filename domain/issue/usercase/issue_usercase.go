package usercase

import (
	_issueRepository "go-issues-api/domain/issue"
	"go-issues-api/domain/model"
	_userRepository "go-issues-api/domain/user"
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
