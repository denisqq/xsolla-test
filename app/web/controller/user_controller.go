package controller

import (
	"github.com/denisqq/xsolla-test/app/model"
	"github.com/denisqq/xsolla-test/app/service"
	"github.com/denisqq/xsolla-test/app/web/dto"
	"github.com/denisqq/xsolla-test/app/web/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{UserService: userService}
}

func (uc *UserController) Register(c *gin.Context) {
	var input request.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := model.User{Username: input.Username, Password: input.Password}
	user, err := uc.UserService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": dto.UserDto{ID: user.ID, Username: user.Username}})
}

func (uc *UserController) Self(c *gin.Context) {
	username := c.GetString(gin.AuthUserKey)
	user, err := uc.UserService.FindByUsername(username)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": dto.UserDto{ID: user.ID, Username: user.Username}})
}
