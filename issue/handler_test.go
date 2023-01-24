package issue

import (
	"go-issues-api/config"
	"go-issues-api/database"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/suite"
)

type SuiteTest struct {
	suite.Suite
	Router *gin.Engine
}

func (test *SuiteTest) SetupSuite() {
	database.Connect("test")
	test.Router = gin.Default()
	test.RegisterHanlder()
}

func (test *SuiteTest) RegisterHanlder() {
	NewIssueHandler(test.Router.Group("api/v1"))
}

func (test *SuiteTest) TearDownSuite() {
	database.Disconnect("test")
}

func (test *SuiteTest) SetupTest() {
	dbConfig := config.NewPostgresConfig("test")
	m := database.NewMigrate("../database/migrations", dbConfig.Url())
	m.Up()
	Seed()
}

func (test *SuiteTest) TearDownTest() {
	dbConfig := config.NewPostgresConfig("test")
	m := database.NewMigrate("../database/migrations", dbConfig.Url())
	m.Down()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SuiteTest))
}

type RequestParams struct {
	Action string
	Url    string
	Body   io.Reader
}

func (test *SuiteTest) RequestHelper(params *RequestParams) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()
	request, _ := http.NewRequest(params.Action, params.Url, params.Body)
	if params.Body != nil {
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	test.Router.ServeHTTP(response, request)
	return response
}

func (t *SuiteTest) TestGetIssues() {
	res := t.RequestHelper(
		&RequestParams{
			Action: http.MethodGet,
			Url:    "/api/v1/issues",
			Body:   nil,
		},
	)
	assert.Equal(t.T(), http.StatusOK, res.Code)
	assert.Equal(t.T(), "{\"issues\":[{\"ID\":1,\"Title\":\"issue 1\",\"Description\":\"This is issue 1\"},{\"ID\":2,\"Title\":\"issue 2\",\"Description\":\"This is issue 2\"}]}", res.Body.String())
}

func (t *SuiteTest) TestCreateIssue() {
	res := t.RequestHelper(
		&RequestParams{
			Action: http.MethodPost,
			Url:    "/api/v1/issues",
			Body:   strings.NewReader("title=test&description=test test test"),
		},
	)
	assert.Equal(t.T(), http.StatusOK, res.Code)
	assert.Equal(t.T(), "{\"issue\":{\"ID\":3,\"Title\":\"test\",\"Description\":\"test test test\"}}", res.Body.String())
}

func (t *SuiteTest) TestGetIssue() {
	res := t.RequestHelper(
		&RequestParams{
			Action: http.MethodGet,
			Url:    "/api/v1/issues/2",
			Body:   nil,
		},
	)
	assert.Equal(t.T(), http.StatusOK, res.Code)
	assert.Equal(t.T(), "{\"issue\":{\"ID\":2,\"Title\":\"issue 2\",\"Description\":\"This is issue 2\"}}", res.Body.String())

	res = t.RequestHelper(
		&RequestParams{
			Action: http.MethodGet,
			Url:    "/api/v1/issues/4",
			Body:   nil,
		},
	)
	assert.Equal(t.T(), http.StatusNotFound, res.Code)
	assert.Equal(t.T(), "{\"message\":\"id 4 is not found\"}", res.Body.String())
}

func (t *SuiteTest) TestUpdateIssue() {
	res := t.RequestHelper(
		&RequestParams{
			Action: http.MethodPut,
			Url:    "/api/v1/issues/1",
			Body:   strings.NewReader("title=test&description=test test test"),
		},
	)
	assert.Equal(t.T(), http.StatusOK, res.Code)
	assert.Equal(t.T(), "{\"issue\":{\"ID\":1,\"Title\":\"test\",\"Description\":\"test test test\"}}", res.Body.String())

	res = t.RequestHelper(
		&RequestParams{
			Action: http.MethodPut,
			Url:    "/api/v1/issues/4",
			Body:   strings.NewReader("title=test&description=test test test"),
		},
	)
	assert.Equal(t.T(), http.StatusNotFound, res.Code)
	assert.Equal(t.T(), "{\"message\":\"id 4 is not found\"}", res.Body.String())
}

func (t *SuiteTest) TestDeleteIssue() {
	res := t.RequestHelper(
		&RequestParams{
			Action: http.MethodDelete,
			Url:    "/api/v1/issues/2",
			Body:   nil,
		},
	)
	assert.Equal(t.T(), http.StatusOK, res.Code)
	assert.Equal(t.T(), "{\"message\":\"id 2 is removed\"}", res.Body.String())

	res = t.RequestHelper(
		&RequestParams{
			Action: http.MethodDelete,
			Url:    "/api/v1/issues/4",
			Body:   nil,
		},
	)
	assert.Equal(t.T(), http.StatusNotFound, res.Code)
	assert.Equal(t.T(), "{\"message\":\"id 4 is not found\"}", res.Body.String())
}
