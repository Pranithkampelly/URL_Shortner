package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H {
				"message": "pong",
			})
		})
	}
	api.GET("/jokes", JokeHandler)

	// Start and run the server
	router.Run(":3000")
}
func JokeHandler(c *gin.Context) {
	c.Header("Content-Type", "template/welcome-template.html")

}

// LikeJoke increments the likes of a particular joke Item
