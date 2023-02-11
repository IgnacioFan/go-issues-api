package issue

import (
	_issueRepo "go-issues-api/internal/issue/mocks"
	_usecase "go-issues-api/internal/issue/usecase"
	"go-issues-api/internal/model"
	_userRepo "go-issues-api/internal/user/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
)

type args struct {
	userId      int
	title       string
	description string
}

var (
	issueMock = new(_issueRepo.Repository)
	userMock  = new(_userRepo.Repository)
	auther    = &model.User{
		ID:   1,
		Name: "Foo Bar",
	}
)

func TestFindAll(t *testing.T) {
	testCases := []struct {
		name          string
		input         interface{}
		expected      []*model.Issue
		expectedError error
	}{
		{
			"Get all issues",
			nil,
			[]*model.Issue{
				{
					ID:          1,
					Title:       "issue 1",
					Description: "This is issue 1",
					Author:      *auther,
				},
				{
					ID:          2,
					Title:       "issue 2",
					Description: "This is issue 2",
					Author:      *auther,
				},
			},
			nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			issueMock.On("FindAll").Return(testCase.expected, nil)
			usecase := _usecase.NewIssueUsecase(userMock, issueMock)

			res, err := usecase.FindAll()
			if testCase.expectedError != nil {
				assert.Equal(t, testCase.expected, res)
				assert.Equal(t, nil, err)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	testCases := []struct {
		name          string
		args          *args
		expected      *model.Issue
		expectedError error
	}{
		{
			"Create issue",
			&args{
				userId:      1,
				title:       "Foo",
				description: "Bar",
			},
			&model.Issue{
				ID:          1,
				Title:       "Foo",
				Description: "Bar",
				Author:      *auther,
			},
			nil,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			userId, title, description := test.args.userId, test.args.title, test.args.description

			userMock.On("Get", userId).Return(*auther, nil)
			issueMock.On("Create", &model.Issue{Title: title, Description: description, Author: *auther}).Return(nil)
			usecase := _usecase.NewIssueUsecase(userMock, issueMock)

			res, err := usecase.Create(userId, title, description)
			if test.expectedError != nil {
				assert.Equal(t, test.expected, res)
				assert.Equal(t, nil, err)
			}
		})
	}
}
