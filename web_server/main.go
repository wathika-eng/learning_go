package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Car struct {
	ID          string `json:"ID"`
	NumberPlate string `json:"Number Plate"`
	Engine      string `json:"Engine"`
	YOM         int    `json:"YOM"`
}

type Drivers struct {
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Age     int    `json:"Age"`
	DLClass string `json:"Dl Class"`
	Matatu  Car    `json:"Matatu"`
}

// cache data
var DriverCache = make(map[int]Drivers)

// lock read/write
var cacheMutex sync.RWMutex

func main() {
	// driver1 := Drivers{}
	// driver1.GenData()
	// fmt.Printf("Driver Details:\n%+v\n", driver1)
	PopulateCache(10)
	// server setup
	router := http.NewServeMux()
	// routes
	router.HandleFunc("GET /drivers", handleGetDrivers)
	router.HandleFunc("POST /drivers", handleCreateDrivers)
	// get user by ID
	router.HandleFunc("GET /drivers/{id}", getDriver)
	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
