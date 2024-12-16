package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// send response and write it, body, headers of request
func handleGetDrivers(w http.ResponseWriter, r *http.Request) {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(DriverCache)
}

func handleCreateDrivers(w http.ResponseWriter, r *http.Request) {
	var driver Drivers

	// Attempt to decode the incoming JSON request
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		if err.Error() == "EOF" {
			// If the body is empty, proceed with generating all data
			driver.GenData()
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		// If decoding succeeds, fill in any missing fields using GenData
		driver.GenData()
	}

	// Ensure the Name field is populated
	if driver.Name == "" {
		http.Error(w, "name required", http.StatusBadRequest)
		return
	}

	// Lock the mutex before updating the cache
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	// Add driver to the cache
	DriverCache[len(DriverCache)+1] = driver
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(driver)
	// Respond with success
	w.WriteHeader(http.StatusOK)
}

func getDriver(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	cacheMutex.RLock()
	user, ok := DriverCache[id]
	cacheMutex.RLock()
	if !ok {
		http.Error(w, "driver not found", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	driver, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	w.Write(driver)
}
