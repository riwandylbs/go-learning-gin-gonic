package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/riwandylbs/go-learning-gin-gonic/config"
	"github.com/riwandylbs/go-learning-gin-gonic/controllers"
	"github.com/riwandylbs/go-learning-gin-gonic/middleware"
	"github.com/riwandylbs/go-learning-gin-gonic/repository"
	"github.com/riwandylbs/go-learning-gin-gonic/service"
)

var (
	db                                        = config.DatabaseConnections()
	userRepository repository.UserRepository  = repository.NewUserRepository(db)
	userService    service.UserService        = service.NewUserService(userRepository)
	jwtService     service.JWTService         = service.NewJWTService()
	userController controllers.UserController = controllers.NewUserController(userService, jwtService)
)

func SetupRoutes() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to gin server",
		})
	})

	r.GET("/all-users", middleware.AuthorizeJWT(), userController.GetAllUser)

	r.Run(":8080")
}
