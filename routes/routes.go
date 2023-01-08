package routes

import (
	"time"

	"github.com/gin-contrib/cors"
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
	r := gin.New()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/validate/me", middleware.AuthorizeHeader(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to gin with group routes",
		})
	})

	// grouping api with middleware authentication
	apiGroup := r.Group("/api", middleware.AuthorizeJWT())
	apiGroup.GET("/all-users", userController.GetAllUser)

	r.Run(":8080")
}
