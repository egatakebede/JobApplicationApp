package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // create a Gin router with default middleware (logger + recovery)

	// Example route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8080") // run server on localhost:8080
}
