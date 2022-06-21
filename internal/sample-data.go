package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"time"
)

func main() {
	r := gin.Default()
	currentTime := time.Now()
	r.GET("/twitter-post", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"title":   "Vykas Release",
			"message": "Vykas has been delayed",
			"date":    currentTime.Format("01-02-2006"),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
