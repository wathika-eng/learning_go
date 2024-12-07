package models

import "time"

// shape of event, required fields marked
type Event struct {
	Id          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Time        time.Time `binding:"required"`
	UserID      int
}

// store events
var events = make([]Event, 0, 10)

// methods to interact with Events in the DB
func (e *Event) Save() {
	events = append(events, *e)
}

// get all events, not a method
func GetAllEvents() *[]Event {
	return &events
}
