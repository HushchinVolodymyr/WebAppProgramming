package Controllers

import (
	"encoding/json"
	"example.com/m/Models/Task2"
	"example.com/m/Services"
	"io/ioutil"
	"net/http"
)

// Controller for Task2. Function for POST request
func PostTask2Controller(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Unmarshal the request body
	var req Task2.Task2RequestModel
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	CoalEmission := Services.CalculateCoalEmission(req.CoalWeight, req.CoalHeatValue)
	OilEmission := Services.CalculateOilEmission(req.OilWeight, req.OilHeatValue)
	GasEmission := Services.CalculateGasEmission(req.GasWeight, req.GasHeatValue)

	// Call the service with the request data
	result := Task2.Task2ResponseModel{
		CoalEmissionFactor: Services.GetCoalEmissionFactor(),
		CoalEmission:       CoalEmission,
		OilEmisionFactor:   Services.GetOilEmissionFactor(),
		OilEmission:        OilEmission,
		GasEmisionFactor:   Services.GetGasEmissionFactor(),
		GasEmission:        GasEmission,
		TotalEmission:      Services.CalculateTotalEmission(CoalEmission, OilEmission, GasEmission),
	}

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
