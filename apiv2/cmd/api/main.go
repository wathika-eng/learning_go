package main

import (
	"apiv2/pkg/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const port = 8080

func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvents)

	fmt.Printf("server listening on http://localhost:%v\n", port)
	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
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
	event.Id = 1
	event.UserID = 1
	c.JSON(http.StatusCreated, gin.H{"message": "Created successfully"})
}
