package handlers

import (
	"apiv2/pkg/models"
	"database/sql"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": errJson})
		return
		// log.Fatal(err)
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event '" + event.Name + "' created successfully",
		"event":   event,
	})
}

func GetEvent(c *gin.Context) {
	// Parse the event ID from the URL parameter
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		// Client error - invalid ID
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid event ID format"})
		return
	}

	// Fetch the event from the database using the parsed ID
	event, err := models.GetEventByID(eventId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Event not found - client error
			c.JSON(http.StatusNotFound, gin.H{"Message": "Event not found"})
		} else {
			// Server error
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to fetch event"})
		}
		return
	}

	// Return the event as JSON
	c.JSON(http.StatusOK, event)
}

func UpdateEvent(c *gin.Context) {
	// Parse the event ID from the URL parameter
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		// Client error - invalid ID
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid event ID format"})
		return
	}

	// Fetch the event from the database using the parsed ID
	_, err = models.GetEventByID(eventId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Event not found - client error
			c.JSON(http.StatusNotFound, gin.H{"Message": "Event not found"})
		} else {
			// Server error
			c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to fetch event"})
		}
		return
	}
	// create a new event populated with incoming data
	var newEvent models.Event
	err = c.ShouldBindJSON(&newEvent)
	if err != nil {
		// Client error - invalid ID
		c.JSON(http.StatusBadRequest, gin.H{"Message": "Invalid data submitted"})
		return
	}
	// write the new event to the DB
	// use the already existing ID, don't create new one
	newEvent.Id = eventId
	err = newEvent.Update()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		// log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event with ID %v updated", "event": newEvent.Name})
}
