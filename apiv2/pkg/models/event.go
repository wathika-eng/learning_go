package models

import (
	"apiv2/pkg/db"
	"errors"
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

// methods to interact with Events in the DB
func (e Event) Save() error {
	// prevent sql injections by using ?
	query := `
	INSERT INTO events (
	name, description, location, time, user_id
	)
	VALUES (?,?,?,?,?)
	`
	// prepare sql statement and handle any error
	// reusable, in memory, complex
	sqlstmt, errSql := db.DB.Prepare(query)
	defer sqlstmt.Close()
	// exec is used when inserting, updating data
	result, errResult := sqlstmt.Exec(e.Name, e.Description, e.Location, e.Time, e.UserID)
	// get the last ID inserted
	id, errID := result.LastInsertId()
	e.Id = id
	if errSql != nil || errResult != nil || errID != nil {
		combinedErr := errors.Join(errSql, errResult, errID)
		return combinedErr
	}
	return nil
	// events = append(events, e)
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

func GetEventByID(id int64){
	query := "SELECT * FROM events WHERE id = ?"
	db.DB.QueryRow()
}