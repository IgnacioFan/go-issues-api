package handler

import (
	"fmt"
	"go-issues-api/internal/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IssueHttp struct {
	Usecase model.IssueUsecase
}

func NewIssueHttp(u model.IssueUsecase) *IssueHttp {
	return &IssueHttp{Usecase: u}
}

func (this *IssueHttp) GetIssues(c *gin.Context) {
	issues, err := this.Usecase.FindAll()

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": issues,
	})
}

func (this *IssueHttp) CreateIssue(c *gin.Context) {
	userId, _ := strconv.Atoi(c.PostForm("id"))
	title := c.PostForm("title")
	description := c.PostForm("description")

	_, err := this.Usecase.Create(userId, title, description)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": "succussfully created an issue!",
	})
}

func (this *IssueHttp) GetIssue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	issue, err := this.Usecase.FindBy(id)

	if err == nil {
		c.JSON(200, gin.H{
			"data": issue,
		})
	} else {
		c.JSON(404, gin.H{
			"data": err.Error(),
		})
	}
}

func (this *IssueHttp) UpdateIssue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.PostForm("title")
	description := c.PostForm("description")

	issue, err := this.Usecase.FindAndUpdate(id, title, description)

	if err == nil {
		c.JSON(200, gin.H{
			"data": issue,
		})
	} else {
		c.JSON(404, gin.H{
			"data": err.Error(),
		})
	}
}

func (this *IssueHttp) DeleteIssue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	affected, err := this.Usecase.DeleteBy(id)

	if err == nil && affected == 1 {
		c.JSON(200, gin.H{
			"data": fmt.Sprintf("id %v is removed", id),
		})
	} else {
		c.JSON(404, gin.H{
			"data": err.Error(),
		})
	}
}
