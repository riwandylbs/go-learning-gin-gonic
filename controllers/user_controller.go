package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riwandylbs/go-learning-gin-gonic/service"
)

type UserController interface {
	GetAllUser(*gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(s service.UserService) *userController {
	return &userController{
		userService: s,
	}
}

func (u *userController) GetAllUser(c *gin.Context) {
	log.Print("Request to service")
	users, err := u.userService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}
