package issue

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewIssueHandler(router *gin.RouterGroup) {
	router.GET("issues", GetIssues)
	router.POST("issues", CreateIssue)
	router.GET("issues/:id", GetIssue)
	router.PUT("issues/:id", UpdateIssue)
	router.DELETE("issues/:id", DeleteIssue)
}

func GetIssues(c *gin.Context) {
	issues, err := FindAll()

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"issues": issues,
	})
}

func CreateIssue(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	issue, err := Create(title, description)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"issue": issue,
	})
}

func GetIssue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	issue, err := Find(id)

	if err == nil {
		c.JSON(200, gin.H{
			"issue": issue,
		})
	} else {
		c.JSON(404, gin.H{
			"message": fmt.Sprintf("id %v is not found", id),
		})
	}
}

func UpdateIssue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.PostForm("title")
	description := c.PostForm("description")
	issue, err := FindAndUpdate(id, title, description)

	if err == nil {
		c.JSON(200, gin.H{
			"issue": issue,
		})
	} else {
		c.JSON(404, gin.H{
			"message": fmt.Sprintf("id %v is not found", id),
		})
	}
}

func DeleteIssue(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	affected, err := Delete(id)

	if err == nil && affected == 1 {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("id %v is removed", id),
		})
	} else {
		c.JSON(404, gin.H{
			"message": fmt.Sprintf("id %v is not found", id),
		})
	}
}
