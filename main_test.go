package main

import (
	"go-issues-api/database"
	"go-issues-api/routes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
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

type SuiteTest struct {
	suite.Suite
	db *gorm.DB
}

var router *gin.Engine

func (test *SuiteTest) SetupSuite() {
	database.ConnectDatabase("test")
	database.SeedIssues()
	router = routes.SetupRouter()
}

func (test *SuiteTest) TearDownSuite() {
	database.DisconnectDatabase()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SuiteTest))
}

func (t *SuiteTest) TestIssuesRoute() {
	res := RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodGet,
			Url:    "/api/v1/issues",
			Body:   nil,
		},
	)
	assert.Equal(t.T(), http.StatusOK, res.Code)
	assert.Equal(t.T(), "{\"issues\":[{\"ID\":1,\"Title\":\"issue 1\",\"Description\":\"This is issue 1\"},{\"ID\":2,\"Title\":\"issue 2\",\"Description\":\"This is issue 2\"}]}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodPost,
			Url:    "/api/v1/issues",
			Body:   strings.NewReader("title=test&description=test test test"),
		},
	)
	assert.Equal(t.T(), http.StatusOK, res.Code)
	assert.Equal(t.T(), "{\"issue\":{\"ID\":3,\"Title\":\"test\",\"Description\":\"test test test\"}}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodGet,
			Url:    "/api/v1/issues/2",
			Body:   nil,
		},
	)
	assert.Equal(t.T(), http.StatusOK, res.Code)
	assert.Equal(t.T(), "{\"issue\":{\"ID\":2,\"Title\":\"issue 2\",\"Description\":\"This is issue 2\"}}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodGet,
			Url:    "/api/v1/issues/4",
			Body:   nil,
		},
	)
	assert.Equal(t.T(), http.StatusNotFound, res.Code)
	assert.Equal(t.T(), "{\"message\":\"id 4 is not found\"}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodPut,
			Url:    "/api/v1/issues/1",
			Body:   strings.NewReader("title=test&description=test test test"),
		},
	)
	assert.Equal(t.T(), http.StatusOK, res.Code)
	assert.Equal(t.T(), "{\"issue\":{\"ID\":1,\"Title\":\"test\",\"Description\":\"test test test\"}}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodPut,
			Url:    "/api/v1/issues/4",
			Body:   strings.NewReader("title=test&description=test test test"),
		},
	)
	assert.Equal(t.T(), http.StatusNotFound, res.Code)
	assert.Equal(t.T(), "{\"message\":\"id 4 is not found\"}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodDelete,
			Url:    "/api/v1/issues/2",
			Body:   nil,
		},
	)
	assert.Equal(t.T(), http.StatusOK, res.Code)
	assert.Equal(t.T(), "{\"message\":\"id 2 is removed\"}", res.Body.String())

	res = RequestHelper(
		router,
		&RequestParams{
			Action: http.MethodDelete,
			Url:    "/api/v1/issues/4",
			Body:   nil,
		},
	)
	assert.Equal(t.T(), http.StatusNotFound, res.Code)
	assert.Equal(t.T(), "{\"message\":\"id 4 is not found\"}", res.Body.String())
}
