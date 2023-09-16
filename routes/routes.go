package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maoudev/veterinaya-go/handlers"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Bienvenido a mi api",
		})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/signin", handlers.SignIn)
	}
}
