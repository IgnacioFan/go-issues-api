package handler

import (
	"go-issues-api/core/model"
	"go-issues-api/core/user/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestCreateIssue(t *testing.T) {
	res := &model.User{
		ID:   1,
		Name: "Foo Bar",
	}
	usecase := new(mocks.Usecase)
	usecase.On("Create", "Foo Bar").Return(res, nil)
	handler := NewUserHttp(usecase)

	r := gin.Default()
	r.POST("api/v1/users", handler.CreateUser)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader("name=Foo Bar"))
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
	assert.Equal(t, "{\"data\":{\"id\":1,\"name\":\"Foo Bar\"}}", w.Body.String())
}
