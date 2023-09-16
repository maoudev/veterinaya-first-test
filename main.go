package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya-go/database"
	"github.com/maoudev/veterinaya-go/routes"
)

func main() {
	err := database.ConnectDB()
	if err != nil {
		panic(err.Error())
	}

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":3000")
}
