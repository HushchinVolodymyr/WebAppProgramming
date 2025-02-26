package Controllers

import (
	"encoding/json"
	"example.com/m/Models/Task5"
	"example.com/m/Services"
	"io/ioutil"
	"net/http"
)

func PostTask5Controller(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Unmarshal the request body
	var req Task5.RequestModel
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	circuitReliability := Services.CalcReliable(req.EGActivity, req.Pl, req.PlLong, req.Transmission, req.Activator, req.Connection, req.ConnectionTimes)
	losses := Services.CalcLosses(req.CostAvaria, req.CostPlan)

	result := Task5.ResponseModel{
		CircuitReliability: circuitReliability,
		Losses:             losses,
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
