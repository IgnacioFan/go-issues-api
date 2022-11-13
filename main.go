package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Issue struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func BuildIssue(title string, description string) (issue Issue) {
	return Issue{
		ID:          len(issesTable) + 1,
		Title:       title,
		Description: description,
	}
}

func FindIssueIndex(id int) int {
	for index, issue := range issesTable {
		if issue.ID == id {
			return index
		}
	}
	return -1
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
		ID:          1,
		Title:       "issue 1",
		Description: "This is issue 1",
	},
	{
		ID:          2,
		Title:       "issue 2",
		Description: "This is issue 2",
	},
	{
		ID:          3,
		Title:       "issue 3",
		Description: "This is issue 3",
	},
	{
		ID:          4,
		Title:       "issue 4",
		Description: "This is issue 4",
	},
	{
		ID:          5,
		Title:       "issue 5",
		Description: "This is issue 5",
	},
}

func main() {
	router := gin.Default()

	router.GET("/api/v1/issues", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"issues": RenderIssues(issesTable),
		})
	})

	router.POST("/api/v1/issues", func(c *gin.Context) {
		title := c.PostForm("title")
		description := c.PostForm("description")

		item := BuildIssue(title, description)
		issesTable = append(issesTable, item)

		c.JSON(200, gin.H{
			"issue": RenderIssue(item),
		})
	})

	router.GET("/api/v1/issues/:id", func(c *gin.Context) {
		issueId, _ := strconv.Atoi(c.Param("id"))
		index := FindIssueIndex(issueId)

		if index == -1 {
			c.JSON(404, gin.H{
				"message": fmt.Sprintf("id %v is not found", issueId),
			})
		} else {
			c.JSON(200, gin.H{
				"issue": RenderIssue(issesTable[index]),
			})
		}
	})

	router.PUT("/api/v1/issues/:id", func(c *gin.Context) {
		issueId, _ := strconv.Atoi(c.Param("id"))
		index := FindIssueIndex(issueId)

		if index == -1 {
			c.JSON(404, gin.H{
				"message": fmt.Sprintf("id %v is not found", issueId),
			})
		} else {
			issue := &issesTable[index]
			issue.Title = c.PostForm("title")
			issue.Description = c.PostForm("description")

			c.JSON(200, gin.H{
				"issue": RenderIssue(issesTable[index]),
			})
		}
	})

	router.DELETE("/api/v1/issues/:id", func(c *gin.Context) {
		issueId, _ := strconv.Atoi(c.Param("id"))
		index := FindIssueIndex(issueId)

		if index == -1 {
			c.JSON(404, gin.H{
				"message": fmt.Sprintf("id %v is not found", issueId),
			})
		} else {
			issesTable = append(issesTable[:index], issesTable[index+1:]...)

			c.JSON(200, gin.H{
				"message": fmt.Sprintf("id %v is removed", issueId),
			})
		}
	})

	router.Run()
}
