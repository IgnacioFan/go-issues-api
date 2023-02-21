package issue

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_handler "go-issues-api/internal/issue/handler"
	"go-issues-api/internal/model"
	mocks "go-issues-api/tests/mocks/issue"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

var (
	router           = gin.Default()
	issueUsecaseMock = new(mocks.IssueUsecase)
)

func TestGetIssues(t *testing.T) {
	auther := &model.User{
		ID:   1,
		Name: "Foo Bar",
	}
	result := []*model.Issue{
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
	}

	usecase := new(mocks.IssueUsecase)
	usecase.On("FindAll").Return(result, nil)
	handler := _handler.NewIssueHttp(usecase)

	r := gin.Default()
	r.GET("api/v1/issues", handler.GetIssues)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/issues", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"data\":[{\"id\":1,\"title\":\"issue 1\",\"description\":\"This is issue 1\",\"author\":{\"id\":1,\"name\":\"Foo Bar\"}},{\"id\":2,\"title\":\"issue 2\",\"description\":\"This is issue 2\",\"author\":{\"id\":1,\"name\":\"Foo Bar\"}}]}", w.Body.String())
}

func TestCreateIssue(t *testing.T) {
	usecase := new(mocks.IssueUsecase)
	usecase.On("Create", 1, "test", "test test test").Return(nil, nil)
	handler := _handler.NewIssueHttp(usecase)

	r := gin.Default()
	r.POST("api/v1/issues", handler.CreateIssue)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/issues", strings.NewReader("id=1&title=test&description=test test test"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"data\":\"succussfully created an issue!\"}", w.Body.String())
}

func TestGetIssue(t *testing.T) {
	auther := &model.User{
		ID:   1,
		Name: "Foo Bar",
	}
	result := &model.Issue{
		ID:          1,
		Title:       "issue",
		Description: "Hello world",
		Author:      *auther,
	}

	usecase := new(mocks.IssueUsecase)
	usecase.On("FindBy", 1).Return(result, nil)
	handler := _handler.NewIssueHttp(usecase)

	r := gin.Default()
	r.GET("api/v1/issues/:id", handler.GetIssue)

	req, err := http.NewRequest(http.MethodGet, "/api/v1/issues/1", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"data\":{\"id\":1,\"title\":\"issue\",\"description\":\"Hello world\",\"author\":{\"id\":1,\"name\":\"Foo Bar\"}}}", w.Body.String())

	usecase.On("FindBy", 2).Return(nil, errors.New("record not found"))

	req, err = http.NewRequest(http.MethodGet, "/api/v1/issues/2", nil)
	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "{\"data\":\"record not found\"}", w.Body.String())
}

func TestUpdateIssue(t *testing.T) {
	auther := &model.User{
		ID:   1,
		Name: "Foo Bar",
	}
	result := &model.Issue{
		ID:          1,
		Title:       "updated title",
		Description: "updated text",
		Author:      *auther,
	}

	usecase := new(mocks.IssueUsecase)
	usecase.On("FindAndUpdate", 1, "updated title", "updated text").Return(result, nil)
	handler := _handler.NewIssueHttp(usecase)

	r := gin.Default()
	r.PUT("api/v1/issues/:id", handler.UpdateIssue)

	req, err := http.NewRequest(http.MethodPut, "/api/v1/issues/1", strings.NewReader("title=updated title&description=updated text"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"data\":{\"id\":1,\"title\":\"updated title\",\"description\":\"updated text\",\"author\":{\"id\":1,\"name\":\"Foo Bar\"}}}", w.Body.String())

	usecase.On("FindAndUpdate", 2, "updated title", "updated text").Return(nil, errors.New("record not found"))

	req, err = http.NewRequest(http.MethodPut, "/api/v1/issues/2", strings.NewReader("title=updated title&description=updated text"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "{\"data\":\"record not found\"}", w.Body.String())
}

func TestDeleteIssue(t *testing.T) {
	var affected int64 = 1
	usecase := new(mocks.IssueUsecase)
	usecase.On("DeleteBy", 1).Return(affected, nil)
	handler := _handler.NewIssueHttp(usecase)

	r := gin.Default()
	r.DELETE("api/v1/issues/:id", handler.DeleteIssue)

	req, err := http.NewRequest(http.MethodDelete, "/api/v1/issues/1", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"data\":\"id 1 is removed\"}", w.Body.String())

	affected = 0
	usecase.On("DeleteBy", 2).Return(affected, errors.New("record not found"))

	req, err = http.NewRequest(http.MethodDelete, "/api/v1/issues/2", nil)

	if err != nil {
		t.Fatalf("Couldn't create request: %v\n", err)
	}

	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "{\"data\":\"record not found\"}", w.Body.String())
}

func TestVoteIssue(t *testing.T) {
	upvoteIssue := &model.VoteIssue{ID: 1, IssueId: 1, UserId: 1, Vote: 1}
	downvoteIssue := &model.VoteIssue{ID: 1, IssueId: 1, UserId: 1, Vote: -1}
	tests := []struct {
		name        string
		usecase     func(u *mocks.IssueUsecase)
		args        [2]int
		expected    *model.VoteIssue
		expectedErr error
	}{
		{
			"when user upvote",
			func(usecase *mocks.IssueUsecase) {
				usecase.On("Vote", 1, 1, 1).Return(upvoteIssue, nil)
			},
			[2]int{1, 1},
			upvoteIssue,
			nil,
		},
		{
			"when user downvote",
			func(usecase *mocks.IssueUsecase) {
				usecase.On("Vote", 1, 1, -1).Return(downvoteIssue, nil)
			},
			[2]int{1, -1},
			downvoteIssue,
			nil,
		},
	}

	handler := _handler.NewIssueHttp(issueUsecaseMock)
	router.POST("api/v1/issues/:id/vote", handler.VoteIssue)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.usecase(issueUsecaseMock)

			// set request body
			req := fmt.Sprintf("user_id=%v&vote=%v", test.args[0], test.args[1])

			// get response
			w := performRequest(router, http.MethodPost, "/api/v1/issues/1/vote", strings.NewReader(req))
			assert.Equal(t, http.StatusOK, w.Code)

			// verify response status
			var response map[string]model.VoteIssue
			err := json.Unmarshal([]byte(w.Body.String()), &response)
			assert.Equal(t, test.expectedErr, err)

			// verify response body
			values, ok := response["data"]
			assert.Equal(t, true, ok)
			assert.Equal(t, test.expected, values)
		})
	}
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
