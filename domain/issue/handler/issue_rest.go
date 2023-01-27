package handler

import (
	"go-issues-api/domain/issue"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IssueRest struct {
	Usecase issue.Usecase
}

func NewIssueRest(usecase issue.Usecase) *IssueRest {
	return &IssueRest{Usecase: usecase}
}

func (this *IssueRest) GetIssues(c *gin.Context) {
	issues, err := this.Usecase.GetAll()

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": issues,
	})
}

func (this *IssueRest) CreateIssue(c *gin.Context) {
	userId, _ := strconv.Atoi(c.PostForm("id"))
	title := c.PostForm("title")
	description := c.PostForm("description")

	err := this.Usecase.Create(userId, title, description)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": "succussfully created an issue!",
	})
}

func (this *IssueRest) GetIssue(c *gin.Context) {
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

func (this *IssueRest) UpdateIssue(c *gin.Context) {
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

// func DeleteIssue(c *gin.Context) {
// 	id, _ := strconv.Atoi(c.Param("id"))
// 	affected, err := Delete(id)

// 	if err == nil && affected == 1 {
// 		c.JSON(200, gin.H{
// 			"message": fmt.Sprintf("id %v is removed", id),
// 		})
// 	} else {
// 		c.JSON(404, gin.H{
// 			"message": fmt.Sprintf("id %v is not found", id),
// 		})
// 	}
// }
