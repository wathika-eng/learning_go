package main

import (
	"apiv2/pkg/db"
	"apiv2/pkg/handlers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	server := gin.Default()
	db.InitDB()
	// url and handler
	server.GET("/events", handlers.GetEvents)
	server.POST("/events", handlers.CreateEvents)
	// get event by ID
	server.GET("/events:id", handlers.GetEvents)
	// Print the port without the colon
	fmt.Printf("server listening on http://localhost%v\n", PORT)

	err := server.Run(PORT)
	if err != nil {
		log.Fatal(err)
	}
}
