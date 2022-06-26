package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setRouter() *gin.Engine {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Redirect if route cannot be handled
	router.RedirectTrailingSlash = true

	// Setup route group for the API
	api := router.Group("/api")
	{
		// POST
		api.POST("/sign-up", signUp)
		//api.POST("/login", loginIn)

		// GET
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	}
	router.NoRoute(func(c *gin.Context) { c.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
