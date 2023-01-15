package routes

import (
	"go-issues-api/issue"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("api/v1")

	v1.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, "success")
	})

	v1.GET("issues", issue.GetIssues)
	v1.POST("issues", issue.CreateIssue)
	v1.GET("issues/:id", issue.GetIssue)
	v1.PUT("issues/:id", issue.UpdateIssue)
	v1.DELETE("issues/:id", issue.DeleteIssue)

	return router
}
