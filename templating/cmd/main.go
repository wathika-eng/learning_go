package main

import (
	"fmt"
	"log"
	"os"
	"templating/internals/db"
	"templating/internals/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadDotenv() string {
	err := godotenv.Load()
	if err != nil {
		log.Printf("failed to load .env file: %v", err)
		return ""
	}
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080" // Default port
	}
	return PORT
}

func main() {
	PORT := loadDotenv()
	db.InitDB()
	defer db.DB.Close()
	server := gin.New()
	server.Static("/views", "./views")
	server.GET("/health", CheckAPIHealth)
	server.GET("/drivers", routes.GetDrivers)
	// instance of drivers
	fmt.Printf("Server running on http://localhost:%v", PORT)
	if err := server.Run(":" + PORT); err != nil {
		log.Fatal(err)
	}
}

func CheckAPIHealth(c *gin.Context) {
	c.JSON(200, gin.H{"message": "API is live"})
}
