package main

import (
	"fmt"
	"go-issues-api/model"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BuildIssue(title string, description string) *model.Issue {
	return &model.Issue{
		Title:       title,
		Description: description,
	}
}

func RenderIssue(item model.Issue) model.Issue {
	return model.Issue{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
	}
}

func RenderIssues(items []model.Issue) (issues []model.Issue) {
	for _, item := range items {
		issue := RenderIssue(item)
		issues = append(issues, issue)
	}
	return issues
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	var db = model.DB

	router.GET("/api/v1/issues", func(c *gin.Context) {
		var issues []model.Issue

		result := db.Find(&issues) // result.RowsAffected, result.Error

		if result.Error == nil {
			c.JSON(200, gin.H{
				"issues": RenderIssues(issues),
			})
		} else {
			c.JSON(404, gin.H{
				"error": "something wrong!",
			})
		}
	})

	router.POST("/api/v1/issues", func(c *gin.Context) {
		title := c.PostForm("title")
		description := c.PostForm("description")

		issue := BuildIssue(title, description)
		result := db.Create(issue)

		if result.Error == nil {
			c.JSON(200, gin.H{
				"issue": RenderIssue(*issue),
			})
		} else {
			c.JSON(404, gin.H{
				"error": "something wrong!",
			})
		}
	})

	router.GET("/api/v1/issues/:id", func(c *gin.Context) {
		issueId, _ := strconv.Atoi(c.Param("id"))

		var issue model.Issue
		res := db.First(&issue, issueId)

		if res.Error == nil {
			c.JSON(200, gin.H{
				"issue": RenderIssue(issue),
			})
		} else {
			c.JSON(404, gin.H{
				"message": fmt.Sprintf("id %v is not found", issueId),
			})
		}
	})

	router.PUT("/api/v1/issues/:id", func(c *gin.Context) {
		issueId, _ := strconv.Atoi(c.Param("id"))

		var issue model.Issue
		res := db.First(&issue, issueId)

		if res.Error == nil {
			issue.Title = c.PostForm("title")
			issue.Description = c.PostForm("description")
			db.Save(&issue)

			c.JSON(200, gin.H{
				"issue": RenderIssue(issue),
			})
		} else {
			c.JSON(404, gin.H{
				"message": fmt.Sprintf("id %v is not found", issueId),
			})
		}
	})

	router.DELETE("/api/v1/issues/:id", func(c *gin.Context) {
		issueId, _ := strconv.Atoi(c.Param("id"))
		res := db.Delete(&model.Issue{}, issueId)

		if res.Error == nil && res.RowsAffected == 1 {
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("id %v is removed", issueId),
			})
		} else {
			c.JSON(404, gin.H{
				"message": fmt.Sprintf("id %v is not found", issueId),
			})
		}
	})

	return router
}

func add(a, b int) int {
	return a + b
}

func main() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Taipei",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME_DEV"),
	)
	model.SetupDatabase(dsn)
	router := SetupRouter()
	router.Run(":3000")
}
