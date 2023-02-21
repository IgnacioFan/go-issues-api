package vote_issue

import (
	"database/sql/driver"
	"errors"
	"go-issues-api/internal/model"
	_repo "go-issues-api/internal/vote_issue/repository"
	"go-issues-api/tests/helper"
	"regexp"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	voteIssueId uint = 1
	userId      uint = 1
	issueId     uint = 1
	upvote      int  = 1
	downvote    int  = -1
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestCreate(t *testing.T) {
	tests := []struct {
		name   string
		input  *model.VoteIssue
		runSQL func(mock sqlmock.Sqlmock)
	}{
		{
			name:  "Create success",
			input: &model.VoteIssue{IssueId: issueId, UserId: userId, Vote: upvote},
			runSQL: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "vote_issues"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "issue_id", "user_id", "vote"}).
						AddRow(voteIssueId, issueId, userId, upvote))
				mock.ExpectCommit()
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			orm, mock := helper.SetupGormMock(t)
			test.runSQL(*mock)

			repo := _repo.NewVoteIssueRepository(orm)
			vote_issue, err := repo.Create(test.input)

			if err != nil {
				assert.Equal(t, errors.New("record not found"), err)
			} else {
				assert.Equal(t, voteIssueId, vote_issue.ID)
				assert.Equal(t, test.input.IssueId, vote_issue.IssueId)
				assert.Equal(t, test.input.UserId, vote_issue.UserId)
				assert.Equal(t, test.input.Vote, vote_issue.Vote)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		name   string
		input  *model.VoteIssue
		runSQL func(mock sqlmock.Sqlmock)
	}{
		{
			name:  "from upvote to downvote",
			input: &model.VoteIssue{IssueId: issueId, UserId: userId, Vote: downvote},
			runSQL: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "vote_issues"`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "issue_id", "user_id", "vote"}).
						AddRow(voteIssueId, issueId, userId, upvote))
				mock.ExpectBegin()
				mock.ExpectExec(regexp.QuoteMeta(`UPDATE "vote_issues" SET "vote"=$1,"updated_at"=$2 WHERE "id" = $3`)).
					WithArgs(downvote, AnyTime{}, voteIssueId).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
		},
		{
			name:  "from upvote to upvote",
			input: &model.VoteIssue{IssueId: issueId, UserId: userId, Vote: upvote},
			runSQL: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "vote_issues" WHERE`)).
					WithArgs(issueId, userId).
					WillReturnRows(sqlmock.NewRows([]string{"id", "issue_id", "user_id", "vote"}).
						AddRow(voteIssueId, issueId, userId, upvote))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			orm, mock := helper.SetupGormMock(t)
			test.runSQL(*mock)

			repo := _repo.NewVoteIssueRepository(orm)
			vote_issue, err := repo.Update(test.input)

			if err != nil {
				assert.Equal(t, errors.New("record not found"), err)
			} else {
				assert.Equal(t, voteIssueId, vote_issue.ID)
				assert.Equal(t, test.input.IssueId, vote_issue.IssueId)
				assert.Equal(t, test.input.UserId, vote_issue.UserId)
				assert.Equal(t, test.input.Vote, vote_issue.Vote)
			}
		})
	}
}
