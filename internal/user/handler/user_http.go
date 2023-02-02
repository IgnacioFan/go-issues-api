package handler

import (
	"go-issues-api/internal/user"

	"github.com/gin-gonic/gin"
)

type UserHttp struct {
	Usecase user.Usecase
}

func NewUserHttp(usecase user.Usecase) *UserHttp {
	return &UserHttp{
		Usecase: usecase,
	}
}

func (this *UserHttp) CreateUser(c *gin.Context) {
	name := c.PostForm("name")

	user, err := this.Usecase.Create(name)

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}
