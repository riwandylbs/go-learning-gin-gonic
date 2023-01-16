package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

type loginForm struct {
	username string
	password string
}

func NewLoginController(l service.LoginService, jwtSrv service.JWTService) *loginController {
	return &loginController{
		loginService: l,
		jwtService:   jwtSrv,
	}
}

func (l *loginController) Login(c *gin.Context) {
	var username = c.Param("username")
	var pwd = c.Param("password")

	if ok := l.loginService.Login(username, pwd); !ok {
		c.JSON(http.StatusBadRequest, utils.ApiResponse{
			Code:    http.StatusBadRequest,
			Message: "User or password is wrong",
		})
		return
	}
	token := l.jwtService.GenerateToken(username)
	var res = map[string]string{"token": token}
	c.JSON(http.StatusAccepted, utils.ApiResponse{
		Code:    http.StatusAccepted,
		Message: "Login Success",
		Data:    res,
	})
}
