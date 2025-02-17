package Controllers

import (
	"encoding/json"
	"example.com/m/Models/Task1Part1"
	"example.com/m/Models/Task1Part2"
	"example.com/m/Services"
	"io/ioutil"
	"net/http"
)

// Controller for Task1 Part1. Function for POST request
func PostTask1Part1Controller(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Unmarshal the request body
	var req Task1Part1.Task1Part1RequestStruct
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the service with the request data
	result := Services.Task1Part1Service(Task1Part1.Task1Part1FuelComponentsModel{
		Hydrogen: req.Hydrogen,
		Carbon:   req.Carbon,
		Sulfur:   req.Sulfur,
		Nitrogen: req.Nitrogen,
		Oxygen:   req.Oxygen,
		Moister:  req.Moister,
		Ash:      req.Ash,
	})

	// Call the service
	response, err := json.Marshal(result)

	// Check for error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Controller for Task1 Part2. Function for POST request
func PostTask1Part2Controller(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Unmarshal the request body
	var req Task1Part2.Task1Part2RequestStruct
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the service with the request data
	result := Services.Task1Part2Service(Task1Part2.Task1Part2LowerHeatingComponents{
		Carbon:            req.Carbon,
		Hydrogen:          req.Hydrogen,
		Oxygen:            req.Oxygen,
		Sulfur:            req.Sulfur,
		LowerHeatingValue: req.LowerHeatingValue,
		Moisture:          req.Moisture,
		Ash:               req.Ash,
		Vanadium:          req.Vanadium,
	})

	// Call the service
	response, err := json.Marshal(result)

	// Check for error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
