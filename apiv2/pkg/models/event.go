package models

import (
	"apiv2/pkg/db"
	"fmt"
	"time"
)

// shape of event, required fields marked
type Event struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	Time        time.Time `json:"time" binding:"required"`
	UserID      int       `json:"userID"`
}

// Save method to interact with Events in the DB
func (e *Event) Save() error {
	// Define the query to insert the event, using placeholders for user input (prevents SQL injection)
	query := `
	INSERT INTO events (name, description, location, time, user_id)
	VALUES (?, ?, ?, ?, ?)
	`

	// Prepare the SQL statement for execution
	sqlstmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing the SQL query: %v", err)
	}
	defer sqlstmt.Close() // Ensure the statement is closed after execution

	// Execute the statement with the event data
	result, err := sqlstmt.Exec(e.Name, e.Description, e.Location, e.Time, e.UserID)
	if err != nil {
		return fmt.Errorf("error executing the SQL query: %v", err)
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error retrieving last inserted ID: %v", err)
	}

	// Assign the new ID to the event struct
	e.Id = id

	return nil
}

// get all events, not a method
func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM EVENTS`
	// query is simple and easy
	// use query not exec, to get back rows
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// a slice of events to be appended
	var events = make([]Event, 0, 10)
	// loop runs as long as we have rows with data using a bool from NEXT
	for rows.Next() {
		var event Event
		// read contents of row, returns pointer, can return error
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.Name, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

// func GetEventByID(id int64){
// 	query := "SELECT * FROM events WHERE id = ?"
// 	db.DB.QueryRow()
// }
