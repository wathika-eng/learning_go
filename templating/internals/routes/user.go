package routes

import (
	"templating/internals/handler"

	"github.com/gin-gonic/gin"
)

func GetDrivers(c *gin.Context) {
	// Call the function to get drivers
	drivers, err := handler.GetDriversDB()

	if err != nil {
		// If there is an error, return the error message
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	// If no error, return the list of drivers
	c.JSON(200, gin.H{"drivers": drivers})
}
