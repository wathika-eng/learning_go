package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// logger and recovery middleware, avoids crashes
	server := gin.Default()
	// handlers are functions when url is called
	server.GET("/events", getEvents)
	server.Run(":8080")
	log.Println("running on http://localhost:8080")
}

// context is usually a pointer to gin context if registered as a handler
func getEvents(context *gin.Context) {
	// status code, data
	context.JSON(http.StatusOK, gin.H{"Hello": "World"})
}
