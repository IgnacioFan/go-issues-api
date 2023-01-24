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

	issue.NewIssueHandler(v1)

	return router
}
