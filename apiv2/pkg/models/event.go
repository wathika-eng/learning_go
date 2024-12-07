package models

import "time"

// shape of event, required fields marked
type Event struct {
	Id          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Time        time.Time `json:"time" binding:"required"`
	UserID      int       `json:"userID"`
}

// store events
var events = make([]Event, 0, 10)

// methods to interact with Events in the DB
func (e Event) Save() {
	events = append(events, e)
}

// get all events, not a method
func GetAllEvents() []Event {
	return events
}
