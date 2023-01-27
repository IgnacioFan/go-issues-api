package routes

import (
	_issueHanlderRest "go-issues-api/domain/issue/handler"
	_issueRepository "go-issues-api/domain/issue/repository"
	_issueUsecase "go-issues-api/domain/issue/usercase"
	_userRepository "go-issues-api/domain/user/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	DBConn *gorm.DB
}

func (s *Router) Start() {
	issueRepo := _issueRepository.NewIssueRepository(s.DBConn)
	userRepo := _userRepository.NewUserRepository(s.DBConn)
	issueUsecase := _issueUsecase.NewIssueUsercase(userRepo, issueRepo)
	issueHandler := _issueHanlderRest.NewIssueRest(issueUsecase)

	router := gin.Default()

	v1 := router.Group("api/v1")
	v1.GET("issues", issueHandler.GetIssues)
	v1.POST("issues", issueHandler.CreateIssue)
	v1.GET("issues/:id", issueHandler.GetIssue)
	v1.PUT("issues/:id", issueHandler.UpdateIssue)
	// v1.DELETE("issues/:id", issueHandler.DeleteIssue)

	v1.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, "success")
	})

	router.Run(":3000")
}
