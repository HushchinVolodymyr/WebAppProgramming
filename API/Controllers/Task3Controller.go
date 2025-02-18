package Controllers

import (
	"encoding/json"
	"example.com/m/Models/Task3"
	"example.com/m/Services"
	"io/ioutil"
	"net/http"
)

func PostTask3Controller(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Unmarshal the request body
	var req Task3.Task3RequestModel
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Calculate the values of imbalance
	ProbabilityImbalance := Services.CalculateProbability(req.EnergyPower-req.OldErrorDeviation, req.EnergyPower+req.OldErrorDeviation, 0.01)
	ProfitImbalance := Services.CalculateProfit(req.EnergyPower, ProbabilityImbalance/100, req.EnergyPrice)
	PeniyaImbalance := Services.CalculateProfit(req.EnergyPower, (100-ProbabilityImbalance)/100, req.EnergyPrice)

	// Calculate the values of balance
	ProbabilityBalance := Services.CalculateProbability(req.EnergyPower-req.NewErrorDeviation, req.EnergyPower+req.NewErrorDeviation, 0.01)
	ProfitBalance := Services.CalculateProfit(req.EnergyPower, ProbabilityBalance/100, req.EnergyPrice)
	PeniyaBalance := Services.CalculateProfit(req.EnergyPower, (100-ProbabilityBalance)/100, req.EnergyPrice)

	// Calculate Total Balance
	TotalImbalance := ProfitImbalance - PeniyaImbalance
	TotalBalance := ProfitBalance - PeniyaBalance

	// Create the response model
	result := Task3.Task3ResponseModel{
		ProbabilityImbalance: ProbabilityImbalance,
		ProfitImbalance:      ProfitImbalance,
		PeniyaImbalance:      PeniyaImbalance,
		TotalImbalance:       TotalImbalance,
		ProbabilityBalance:   ProbabilityBalance,
		ProfitBalance:        ProfitBalance,
		PeniyaBalance:        PeniyaBalance,
		TotalBalance:         TotalBalance,
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
