package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"go-issues-api/domain/issue/mocks"
	"go-issues-api/domain/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
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

	usecase := new(mocks.Usecase)
	usecase.On("GetAll").Return(result, nil)
	handler := NewIssueRest(usecase)

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
	usecase := new(mocks.Usecase)
	usecase.On("Create", 1, "test", "test test test").Return(nil)
	handler := NewIssueRest(usecase)

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

	usecase := new(mocks.Usecase)
	usecase.On("FindBy", 1).Return(result, nil)
	handler := NewIssueRest(usecase)

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

// func (t *SuiteTest) TestUpdateIssue() {
// 	res := t.RequestHelper(
// 		&RequestParams{
// 			Action: http.MethodPut,
// 			Url:    "/api/v1/issues/1",
// 			Body:   strings.NewReader("title=test&description=test test test"),
// 		},
// 	)
// 	assert.Equal(t.T(), http.StatusOK, res.Code)
// 	assert.Equal(t.T(), "{\"issue\":{\"ID\":1,\"Title\":\"test\",\"Description\":\"test test test\"}}", res.Body.String())

// 	res = t.RequestHelper(
// 		&RequestParams{
// 			Action: http.MethodPut,
// 			Url:    "/api/v1/issues/4",
// 			Body:   strings.NewReader("title=test&description=test test test"),
// 		},
// 	)
// 	assert.Equal(t.T(), http.StatusNotFound, res.Code)
// 	assert.Equal(t.T(), "{\"message\":\"id 4 is not found\"}", res.Body.String())
// }

// func (t *SuiteTest) TestDeleteIssue() {
// 	res := t.RequestHelper(
// 		&RequestParams{
// 			Action: http.MethodDelete,
// 			Url:    "/api/v1/issues/2",
// 			Body:   nil,
// 		},
// 	)
// 	assert.Equal(t.T(), http.StatusOK, res.Code)
// 	assert.Equal(t.T(), "{\"message\":\"id 2 is removed\"}", res.Body.String())

// 	res = t.RequestHelper(
// 		&RequestParams{
// 			Action: http.MethodDelete,
// 			Url:    "/api/v1/issues/4",
// 			Body:   nil,
// 		},
// 	)
// 	assert.Equal(t.T(), http.StatusNotFound, res.Code)
// 	assert.Equal(t.T(), "{\"message\":\"id 4 is not found\"}", res.Body.String())
// }
