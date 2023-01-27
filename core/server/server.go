package server

import (
	_issueHanlderRest "go-issues-api/core/issue/handler"
	_issueRepository "go-issues-api/core/issue/repository"
	_issueUsecase "go-issues-api/core/issue/usercase"
	_userRepository "go-issues-api/core/user/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DBConn *gorm.DB
}

func (s *Server) Start() {
	issueRepo := _issueRepository.NewIssueRepository(s.DBConn)
	userRepo := _userRepository.NewUserRepository(s.DBConn)
	issueUsecase := _issueUsecase.NewIssueUsercase(userRepo, issueRepo)
	issueHandler := _issueHanlderRest.NewIssueHttp(issueUsecase)

	router := gin.Default()

	v1 := router.Group("api/v1")
	v1.GET("issues", issueHandler.GetIssues)
	v1.POST("issues", issueHandler.CreateIssue)
	v1.GET("issues/:id", issueHandler.GetIssue)
	v1.PUT("issues/:id", issueHandler.UpdateIssue)
	v1.DELETE("issues/:id", issueHandler.DeleteIssue)

	v1.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, "success")
	})

	router.Run(":3000")
}
