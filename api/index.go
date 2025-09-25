package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()
	// Define a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		// Return JSON response
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}

func Handler(r *gin.Context) {

	// Create a Gin router with default middleware (logger and recovery)
	app := setupRouter()
	app.ServeHTTP(r.Writer, r.Request)
}
