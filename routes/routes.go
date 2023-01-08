package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/riwandylbs/go-learning-gin-gonic/controllers"
	"github.com/riwandylbs/go-learning-gin-gonic/repository"
	"github.com/riwandylbs/go-learning-gin-gonic/service"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to gin server",
		})
	})

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	usercontroller := controllers.NewUserController(userService)

	r.GET("/all-users", usercontroller.GetAllUser)

	r.Run(":8080")
}
