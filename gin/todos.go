package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, Gin!"})
	})

	r.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a Gin."})
	})

	r.GET("/todos", func(c *gin.Context) {
		todos := []gin.H{
			{"id": 1, "task": "Learn Go", "done": false},
			{"id": 2, "task": "Build a web app", "done": true},
		}
		c.JSON(http.StatusOK, todos)
	})

	r.Run(":8080")
}
