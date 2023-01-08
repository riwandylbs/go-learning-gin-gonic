package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// get tokenString

		// validate the token

		// if expired return with abort status

		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		// 	"message": "Unauthorized",
		// })

		// Next
		c.Next()
	}
}
