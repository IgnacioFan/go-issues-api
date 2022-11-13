package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/api/v1/issues", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "list all issues",
		})
	})

	router.POST("/api/v1/issues", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Create a issue",
		})
	})

	router.GET("/api/v1/issues/:id", func(c *gin.Context) {
		issueId := c.Param("id")
		c.JSON(200, gin.H{
			"id":      issueId,
			"message": "Get the issue",
		})
	})

	router.PUT("/api/v1/issues/:id", func(c *gin.Context) {
		issueId := c.Param("id")
		c.JSON(200, gin.H{
			"id":      issueId,
			"message": "Update the issue",
		})
	})

	router.DELETE("/api/v1/issues/:id", func(c *gin.Context) {
		issueId := c.Param("id")
		c.JSON(200, gin.H{
			"id":      issueId,
			"message": "Delete the issue",
		})
	})

	router.Run()
}
