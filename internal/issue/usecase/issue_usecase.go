package usecase

import (
	"go-issues-api/internal/model"
)

type IssueUsecase struct {
	userRepository  model.UserRepository
	IssueRepository model.IssueRepository
}

func NewIssueUsecase(user model.UserRepository, issue model.IssueRepository) *IssueUsecase {
	return &IssueUsecase{
		userRepository:  user,
		IssueRepository: issue,
	}
}

func (this *IssueUsecase) FindAll() ([]*model.Issue, error) {
	return this.IssueRepository.FindAll()
}

func (this *IssueUsecase) Create(userId int, title string, description string) (*model.Issue, error) {
	var err error

	user, err := this.userRepository.Get(userId)
	if err != nil {
		panic(err)
	}

	issue := &model.Issue{
		Title:       title,
		Description: description,
		Author:      user,
	}
	err = this.IssueRepository.Create(issue)

	return issue, err
}

func (this *IssueUsecase) FindBy(id int) (*model.Issue, error) {
	issue, err := this.IssueRepository.FindBy(id)
	return issue, err
}

func (this *IssueUsecase) FindAndUpdate(id int, title, description string) (*model.Issue, error) {
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

func (this *IssueUsecase) DeleteBy(id int) (int64, error) {
	var err error
	issue, err := this.IssueRepository.FindBy(id)

	if err != nil {
		return int64(0), err
	}

	affected, err := this.IssueRepository.Delete(int(issue.ID))

	return affected, err
}
