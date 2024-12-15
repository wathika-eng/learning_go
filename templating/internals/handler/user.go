package handler

import (
	"fmt"
	"templating/internals/db"
)

type DriverHandler struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Status     bool   `json:"status"`
	CarOwnerID int64  // Foreign key field
	CarOwner   Cars
}

func GetDriversDB() ([]DriverHandler, error) {
	if db.DB == nil {
		return nil, fmt.Errorf("database connection is nil")
	}
	drivers := make([]DriverHandler, 0, 10)

	query := `SELECT * FROM driver_handlers`
	result, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var driver DriverHandler
		if err := result.Scan(&driver.ID, &driver.Name, &driver.Age, &driver.Status, &driver.CarOwnerID, &driver.CarOwner); err != nil {
			return nil, err
		}
		drivers = append(drivers, driver)
	}

	if err := result.Err(); err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return drivers, nil
}

// func
