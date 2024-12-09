package routes

import (
	"apiv2/pkg/models"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// fetch all events in DB and return in JSON format
func GetEvents(c *gin.Context) {
	var message string
	events, errEvents := models.GetAllEvents()
	if errEvents != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to fetch events"})
		return
		// log.Fatal(err)
	}
	if len(events) == 0 {
		message = "NO EVENTS"
	}
	// returned in context
	c.JSON(http.StatusOK, gin.H{"events": events, "message": message})
}

// create a new event on the DB
func CreateEvents(c *gin.Context) {
	var event models.Event

	// Parse the incoming JSON body into the event struct
	if err := c.ShouldBindJSON(&event); err != nil {
		// Return a bad request error with specific message
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Pre-allocate event fields if required (consider user authentication for dynamic fields)
	event.UserID = 1 // This should ideally come from authenticated user
	event.Id = 1     // Assuming ID is generated or set elsewhere (like DB)

	// // Validate the event data
	// if err := event.CheckEvent(); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// Save the event to the database
	if err := event.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with a success message and the created event details
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event '" + event.Name + "' created successfully",
		"event":   event,
	})
}

// fetch a single event based on ID in DB and return in JSON format
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

// update a certain event on the DB
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

// delete a certain event on the DB
func DeleteEvent(c *gin.Context) {
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
	var deleteEvent models.Event
	deleteEvent.Id = eventId
	err = deleteEvent.Delete()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		// log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted", "event": eventId})
}
