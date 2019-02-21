package http

import (
	"github.com/QianPai/account/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
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

