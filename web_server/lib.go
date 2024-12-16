package main

import (
	"strconv"

	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
)

// GenData method to populate the Drivers struct with random data
func (d *Drivers) GenData() {
	// Generate random age between 22 and 65
	age, _ := faker.RandomInt(22, 65, 1)

	// Generate a fake DL class (A to E)
	classInt, _ := faker.RandomInt(65, 69, 1) // ASCII codes for A to E
	classString := string(rune(classInt[0]))

	// Generate a year of manufacture
	year := faker.YearString()
	YOM, _ := strconv.Atoi(year)

	// Populate the Drivers struct fields
	d.ID = uuid.New().String()
	d.Name = faker.Name()
	d.Age = age[0]
	d.DLClass = classString
	d.Matatu = Car{
		ID:          uuid.New().String(),
		NumberPlate: generateNumberPlate(),
		Engine:      "Toyota",
		YOM:         YOM,
	}
}

func PopulateCache(n int) {
	for i := 0; i < n; i++ {
		var driver Drivers
		driver.GenData()

		cacheMutex.Lock()
		DriverCache[len(DriverCache)+1] = driver
		cacheMutex.Unlock()
	}
}
