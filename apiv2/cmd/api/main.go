package main

import (
	"apiv2/pkg/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	server := gin.Default()
	// url and handler
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)
	server.GET("/hello", Hello)
	// Print the port without the colon
	fmt.Printf("server listening on http://localhost%v\n", PORT)

	err := server.Run(PORT)
	if err != nil {
		log.Fatal(err)
	}
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello")
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	// returned in context
	c.JSON(http.StatusOK, events)
}

// extract incoming data
func createEvents(c *gin.Context) {
	var event models.Event
	// works like scan function in fmt
	// pointer passed, body with model structure, empty structs will be null but can enforce
	err := c.ShouldBindJSON(&event)
	// handle error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed"})
		log.Fatal(err)
	}
	// preallocate
	event.Id = 1
	event.UserID = 1
	event.Save()
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event '" + event.Name + "' created successfully",
		"event":   event,
	})
}
