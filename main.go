package main

import (
	"github.com/gin-gonic/gin"

	"github.com/riwandylbs/go-learning-gin-gonic/config"
)

type Product struct {
	Name  string
	Price int
}

func main() {
	r := gin.Default()

	config.DatabaseConnections()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    200,
			"message": "berhasil",
			"data": &Product{
				Name:  "Product Pertama",
				Price: 1,
			},
		})
	})

	r.Run(":8100")
}
