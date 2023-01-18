package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/riwandylbs/go-learning-gin-gonic/dto"
	"github.com/riwandylbs/go-learning-gin-gonic/service"
	"github.com/riwandylbs/go-learning-gin-gonic/utils"
)

type LoginController interface {
	Login(*gin.Context)
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func NewLoginController(l service.LoginService, jwtSrv service.JWTService) *loginController {
	return &loginController{
		loginService: l,
		jwtService:   jwtSrv,
	}
}

func (l *loginController) Login(c *gin.Context) {

	var loginForm dto.LoginForm
	if error := c.ShouldBindJSON(&loginForm); error != nil {
		log.Fatalf("Error binding to json")
	}

	user, err := l.loginService.Login(loginForm.Username, loginForm.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	token := l.jwtService.GenerateToken(loginForm.Username)

	var userDTO []*dto.UserDTO
	for _, a := range user {
		userDTO = append(userDTO, &dto.UserDTO{
			Id:      a.Id,
			Name:    a.Name,
			Email:   a.Email,
			Address: a.Address,
			Token:   token,
		})
	}

	c.JSON(http.StatusAccepted, utils.ApiResponse{
		Code:    http.StatusAccepted,
		Message: "Login Success",
		Data:    userDTO,
	})
}
