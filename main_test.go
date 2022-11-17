package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type RequestParams struct {
	Action string
	Url    string
	Body   io.Reader
}

func RequestHelper(router *gin.Engine, params *RequestParams) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(params.Action, params.Url, params.Body)
	if params.Body != nil {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	router.ServeHTTP(response, request)
	return response
}

func TestIssuesRoute(t *testing.T) {
	router := SetupRouter()

	res := RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodGet,
			Url:    "/api/v1/issues",
			Body:   nil,
		},
	)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "{\"issues\":[{\"id\":1,\"title\":\"issue 1\",\"description\":\"This is issue 1\"},{\"id\":2,\"title\":\"issue 2\",\"description\":\"This is issue 2\"}]}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodPost,
			Url:    "/api/v1/issues",
			Body:   strings.NewReader("title=test&description=test test test"),
		},
	)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "{\"issue\":{\"id\":3,\"title\":\"test\",\"description\":\"test test test\"}}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodGet,
			Url:    "/api/v1/issues/2",
			Body:   nil,
		},
	)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "{\"issue\":{\"id\":2,\"title\":\"issue 2\",\"description\":\"This is issue 2\"}}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodGet,
			Url:    "/api/v1/issues/4",
			Body:   nil,
		},
	)
	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.Equal(t, "{\"message\":\"id 4 is not found\"}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodPut,
			Url:    "/api/v1/issues/1",
			Body:   strings.NewReader("title=test&description=test test test"),
		},
	)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "{\"issue\":{\"id\":1,\"title\":\"test\",\"description\":\"test test test\"}}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodPut,
			Url:    "/api/v1/issues/4",
			Body:   strings.NewReader("title=test&description=test test test"),
		},
	)
	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.Equal(t, "{\"message\":\"id 4 is not found\"}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodDelete,
			Url:    "/api/v1/issues/2",
			Body:   nil,
		},
	)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "{\"message\":\"id 2 is removed\"}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodDelete,
			Url:    "/api/v1/issues/4",
			Body:   nil,
		},
	)
	assert.Equal(t, http.StatusNotFound, res.Code)
	assert.Equal(t, "{\"message\":\"id 4 is not found\"}", res.Body.String())
}
