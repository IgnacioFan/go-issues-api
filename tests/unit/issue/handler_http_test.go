package issue

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_handler "go-issues-api/internal/issue/handler"
	"go-issues-api/internal/issue/mocks"
	"go-issues-api/internal/model"

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
	usecase := new(mocks.Usecase)
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

	usecase := new(mocks.Usecase)
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

	usecase := new(mocks.Usecase)
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
	usecase := new(mocks.Usecase)
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
