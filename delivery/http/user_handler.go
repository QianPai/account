package http

import (
	"github.com/QianPai/account/model"
	"github.com/QianPai/account/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUsecase usecase.User
}

func NewUserHandler(router gin.Engine, userUsecase usecase.User)  {
	handler := &UserHandler{
		UserUsecase: userUsecase,
	}

    router.GET("/user/:name", handler.GetByName)
	router.POST("/user", handler.Store)
}

func (u *UserHandler) GetByName(c *gin.Context) {
	name := c.Param("name")
	user, err := u.UserUsecase.GetByName(c, name)
	if err != nil {
		c.JSON(http.StatusOK ,gin.H{
			"status": "fail",
			"data" : nil,
		})
	}
	c.JSON(http.StatusOK ,gin.H{
		"status": "success",
		"data" : user,
	})
}

func (u *UserHandler) Store(c *gin.Context) {
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")

	user := &model.User{
		Name:name,
		Phone:phone,
		Password:password,
		CreatedAt:time.Now().String(),
		UpdatedAt:time.Now().String(),
	}

	err := u.UserUsecase.Store(c, user)

	if err != nil {
		c.JSON(http.StatusOK ,gin.H{
			"status": "fail",
			"data" : nil,
		})
	}
	c.JSON(http.StatusOK ,gin.H{
		"status": "success",
		"data" : user,
	})
}

