package Controllers

import (
	"encoding/json"
	"example.com/m/Models/Task4"
	"example.com/m/Services"
	"io/ioutil"
	"net/http"
)

func PostTask4Controller(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Unmarshal the request body
	var req Task4.RequestModel
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cableOption := Services.SelectCable(req.TransformerPowerCable, req.VoltageCable, req.DistanceCable, req.MaxVoltDropCable)

	threePahseCurrent := Services.ThreePhaseCurrent(req.CircuitVoltage, req.CircuitImpedance)
	singlePhaseCurrent := Services.SinglePhaseCurrent(req.CircuitVoltage, req.CircuitImpedance)

	substationState := Services.CalculateCurrents(req.VoltageSubstation, req.RNormal, req.XNormal, req.RMin, req.XMin)

	result := Task4.ResponseModel{
		CableOptions:       cableOption,
		SubstationStates:   substationState,
		ThreePhaseCurrent:  threePahseCurrent,
		SinglePhaseCurrent: singlePhaseCurrent,
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
