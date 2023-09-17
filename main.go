package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya-go/internal/database"
	"github.com/maoudev/veterinaya-go/internal/routes"
)

func main() {
	// Connection with database
	err := database.ConnectDB()
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":3000")
}
