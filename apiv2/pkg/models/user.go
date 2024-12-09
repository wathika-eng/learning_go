package models

import (
	"apiv2/pkg/db"
	"apiv2/pkg/utilis"
	"database/sql"
	"fmt"
)

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	CreatedAt string `json:"time" binding:"required"`
}

// Save method to interact with Users in the DB
func (e *User) Save() error {
	e.CreatedAt = utilis.DBTime()
	// Define the query to insert the event, using placeholders for user input (prevents SQL injection)
	query := `
	INSERT INTO users (name, email, email, password, createdAt)
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
	result, err := sqlstmt.Exec(e.Name, e.Email, e.Password, e.CreatedAt)
	if err != nil {
		return fmt.Errorf("error executing the SQL query: %v", err)
	}

	// Get the last inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error retrieving last inserted ID: %v", err)
	}

	// Assign the new ID to the event struct
	e.ID = id

	return nil
}
