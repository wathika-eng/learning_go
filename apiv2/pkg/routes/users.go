package routes

import (
	"apiv2/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// signup users functionality/endpoint
func CreateUsers(c *gin.Context) {
	var users models.User

	// Parse the incoming JSON body into the users struct
	if err := c.ShouldBindJSON(&users); err != nil {
		// Return a bad request error with specific message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users.ID = 1

	// // Validate the users data
	// if err := users.Checkusers(); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// Save the users to the database
	if err := users.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message and the created users details
	c.JSON(http.StatusCreated, gin.H{
		"message": "users '" + users.Name + "' created successfully",
		"users":   users.Email,
	})
}
