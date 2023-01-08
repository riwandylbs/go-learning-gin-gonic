package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/riwandylbs/go-learning-gin-gonic/service"
	"github.com/riwandylbs/go-learning-gin-gonic/utils"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get tokenString
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):] // get string without bearer string

		token, err := service.NewJWTService().ValidateToken(tokenString)

		if err != nil || token == nil {

			c.AbortWithStatusJSON(http.StatusBadRequest, utils.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: "Unauthorized",
			})
			return
		}

		// validate the token
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			expired := claims["exp"]

			fmt.Println("Expired token at : ", expired)

			// Next
			c.Next()
		}
	}
}
