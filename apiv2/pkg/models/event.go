package models

import (
	"apiv2/pkg/db"
	"apiv2/pkg/utils"
	"database/sql"
	"errors"
	"fmt"
)

// shape of event, required fields marked
type Event struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Location    string `json:"location" binding:"required"`
	Time        string `json:"time"`
	UserID      int    `json:"userID"`
}

// Save method to interact with Events in the DB
func (e *Event) Save() error {
	e.Time = utils.DBTime()
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
	defer func(sqlstmt *sql.Stmt) {
		err := sqlstmt.Close()
		if err != nil {

		}
	}(sqlstmt) // Ensure the statement is closed after execution

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
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
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

// fetch a single event in DB
func GetEventByID(id int64) (*Event, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid event ID: %d", id)
	}

	query := "SELECT id, name, description, location, user_id FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.UserID)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("event with ID %d not found", id)
	} else if err != nil {
		return nil, err
	}

	return &event, nil
}

// update Event
func (e *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, time = ?
	WHERE id=?
	`
	_, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.Time, e.Id)
	if err != nil {
		return err
	}
	return nil
}

// delete an event from DB
func (e *Event) Delete() error {
	query := `
	DELETE FROM events
	WHERE id=?
	`
	_, err := db.DB.Exec(query, e.Id)
	if err != nil {
		return err
	}
	return nil
}

// check if a certain event exists
func (e Event) CheckEvent() error {
	query := `
    SELECT name FROM events
    WHERE name = ?
`
	var existingEventName string
	err := db.DB.QueryRow(query, e.Name).Scan(&existingEventName)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
			// No event found, proceed with inserting the new event
		}
		return errors.New("Event with a similar name already exists")
	}
	return nil
}
