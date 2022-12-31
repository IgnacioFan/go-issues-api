package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Issue struct {
	ID          uint `gorm:"primaryKey"`
	Title       string
	Description string
}

func BuildIssue(title string, description string) *Issue {
	return &Issue{
		Title:       title,
		Description: description,
	}
}

func RenderIssue(item Issue) Issue {
	return Issue{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
	}
}

func RenderIssues(items []Issue) (issues []Issue) {
	for _, item := range items {
		issue := RenderIssue(item)
		issues = append(issues, issue)
	}
	return issues
}

var issesTable = []Issue{
	{
		Title:       "issue 1",
		Description: "This is issue 1",
	},
	{
		Title:       "issue 2",
		Description: "This is issue 2",
	},
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	dsn := "host=localhost user=postgres password=postgres dbname=issues_hub port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&issesTable)

	router.GET("/api/v1/issues", func(c *gin.Context) {
		var issues []Issue

		db.Find(&issues) // result.RowsAffected, result.Error

		c.JSON(200, gin.H{
			"issues": RenderIssues(issues),
		})
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

		var issue Issue
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

		var issue Issue
		db.First(&issue, issueId)
		issue.Title = c.PostForm("title")
		issue.Description = c.PostForm("description")
		res := db.Save(&issue)

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

	router.DELETE("/api/v1/issues/:id", func(c *gin.Context) {
		issueId, _ := strconv.Atoi(c.Param("id"))
		res := db.Delete(&Issue{}, issueId)

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

func main() {
	router := SetupRouter()
	router.Run(":8080")
}
