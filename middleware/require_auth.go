package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/riwandylbs/go-learning-gin-gonic/service"
	"github.com/riwandylbs/go-learning-gin-gonic/utils"
)

type TokenID struct {
	Token string
}

func validateToken(c *gin.Context, tokenString string) {
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

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get tokenString
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):] // get string without bearer string

		validateToken(c, tokenString)
	}
}

func AuthorizeHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerKey := c.GetHeader("API-KEY")
		if headerKey != os.Getenv("API-KEY") {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.ApiResponse{
				Code:    http.StatusBadRequest,
				Message: "Unknown client request",
			})
			return
		}

		// Validate a token post from client
		var tokenString TokenID
		if err := c.ShouldBindJSON(&tokenString); err != nil {
			log.Fatalf("Error binding to json")
		}
		validateToken(c, tokenString.Token)
	}

}
