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

	router.GET("/user", handler.GetUser)
    router.GET("/user/:id", handler.GetById)
	router.POST("/user", handler.Store)
}

func (u *UserHandler) GetUser(c *gin.Context) {
	var user *model.User
	var err error
	name := c.Query("name")
	if name != "" {
		user, err = u.UserUsecase.GetByName(c, name)
	}

	phone := c.Query("phone")
	if phone != "" {
		user, err = u.UserUsecase.GetByPhone(c, phone)
	}
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

func (u *UserHandler) GetById(c *gin.Context) {
	id := c.Param("id")
	user, err := u.UserUsecase.GetById(c, id)
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

