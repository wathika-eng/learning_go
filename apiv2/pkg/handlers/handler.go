package handlers

import (
	"apiv2/pkg/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetEvents(c *gin.Context) {
	events, errEvents := models.GetAllEvents()
	if errEvents != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch events"})
		return
		// log.Fatal(err)
	}
	// returned in context
	c.JSON(http.StatusOK, events)
}

// extract incoming data
func CreateEvents(c *gin.Context) {
	var event models.Event
	// works like scan function in fmt
	// pointer passed, body with model structure, empty structs will be null but can enforce
	errJson := c.ShouldBindJSON(&event)
	// handle error
	// preallocate
	event.Id = 1
	event.UserID = 1
	errSave := event.Save()
	if errJson != nil || errSave != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errJson.Error()})
		return
		// log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event '" + event.Name + "' created successfully",
		"event":   event,
	})
}

func GetEvent(c *gin.Context) {
	// get actual ID value in int64
	// convert to int (what, base, base)
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Couln't pasrse"})
	}
	fmt.Println(eventId)

}
