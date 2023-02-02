package server

import (
	_issueHanlderHttp "go-issues-api/internal/issue/handler"
	_issueRepository "go-issues-api/internal/issue/repository"
	_issueUsecase "go-issues-api/internal/issue/usecase"
	_userHanlderHttp "go-issues-api/internal/user/handler"
	_userRepository "go-issues-api/internal/user/repository"
	_userUsecase "go-issues-api/internal/user/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DBConn *gorm.DB
}

func (s *Server) Start() {
	router := gin.Default()
	v1 := router.Group("api/v1")

	issueRepo := _issueRepository.NewIssueRepository(s.DBConn)
	userRepo := _userRepository.NewUserRepository(s.DBConn)
	issueUsecase := _issueUsecase.NewIssueUsecase(userRepo, issueRepo)

	issueHandler := _issueHanlderHttp.NewIssueHttp(issueUsecase)
	v1.GET("issues", issueHandler.GetIssues)
	v1.POST("issues", issueHandler.CreateIssue)
	v1.GET("issues/:id", issueHandler.GetIssue)
	v1.PUT("issues/:id", issueHandler.UpdateIssue)
	v1.DELETE("issues/:id", issueHandler.DeleteIssue)

	userUsecase := _userUsecase.NewUserUsecase(userRepo)

	userHandler := _userHanlderHttp.NewUserHttp(userUsecase)
	v1.POST("users", userHandler.CreateUser)

	v1.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, "success")
	})

	router.Run(":3000")
}
