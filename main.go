package main

import (
	"github.com/riwandylbs/go-learning-gin-gonic/config"
	"github.com/riwandylbs/go-learning-gin-gonic/routes"
)

func main() {

	db := config.DatabaseConnections()
	routes.SetupRoutes(db)

}
