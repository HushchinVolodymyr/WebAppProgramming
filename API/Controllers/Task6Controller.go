package Controllers

import (
	"encoding/json"
	"example.com/m/Models/Task6"
	"example.com/m/Services"
	"io/ioutil"
	"net/http"
)

func PostTask6Controller(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Unmarshal the request body
	var req Task6.RequestModel
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	grinding := Services.CalcDevice(*req.Grinding.N, *req.Grinding.Cos, *req.Grinding.U, *req.Grinding.Count, *req.Grinding.Ph, *req.Grinding.Kv, *req.Grinding.Tg)
	drilled := Services.CalcDevice(*req.Drilled.N, *req.Drilled.Cos, *req.Drilled.U, *req.Drilled.Count, *req.Drilled.Ph, *req.Drilled.Kv, *req.Drilled.Tg)
	jointing := Services.CalcDevice(*req.Jointing.N, *req.Jointing.Cos, *req.Jointing.U, *req.Jointing.Count, *req.Jointing.Ph, *req.Jointing.Kv, *req.Jointing.Tg)
	circularSaw := Services.CalcDevice(*req.CircularSaw.N, *req.CircularSaw.Cos, *req.CircularSaw.U, *req.CircularSaw.Count, *req.CircularSaw.Ph, *req.CircularSaw.Kv, *req.CircularSaw.Tg)
	press := Services.CalcDevice(*req.Press.N, *req.Press.Cos, *req.Press.U, *req.Press.Count, *req.Press.Ph, *req.Press.Kv, *req.Press.Tg)
	polishing := Services.CalcDevice(*req.Polishing.N, *req.Polishing.Cos, *req.Polishing.U, *req.Polishing.Count, *req.Polishing.Ph, *req.Polishing.Kv, *req.Polishing.Tg)
	shaper := Services.CalcDevice(*req.Shaper.N, *req.Shaper.Cos, *req.Shaper.U, *req.Shaper.Count, *req.Shaper.Ph, *req.Shaper.Kv, *req.Shaper.Tg)
	fan := Services.CalcDevice(*req.Fan.N, *req.Fan.Cos, *req.Fan.U, *req.Fan.Count, *req.Fan.Ph, *req.Fan.Kv, *req.Fan.Tg)

	shr1 := Services.CalcShr(grinding, drilled, jointing, circularSaw, press, polishing, shaper, fan)
	shr2 := &shr1
	shr3 := shr2

	welding := Services.CalcDevice(*req.Welding.N, *req.Welding.Cos, *req.Welding.U, *req.Welding.Count, *req.Welding.Ph, *req.Welding.Kv, *req.Welding.Tg)
	dryer := Services.CalcDevice(*req.Dryer.N, *req.Dryer.Cos, *req.Dryer.U, *req.Dryer.Count, *req.Dryer.Ph, *req.Dryer.Kv, *req.Dryer.Tg)

	fullLoad := Services.CalcFullLoad()

	result := Task6.ResponseModel{
		Grinding:    grinding,
		Drilled:     drilled,
		Jointing:    jointing,
		CircularSaw: circularSaw,
		Press:       press,
		Polishing:   polishing,
		Shaper:      shaper,
		Fan:         fan,

		Shr1: shr1,
		Shr2: *shr2,
		Shr3: *shr3,

		Welding: welding,
		Dryer:   dryer,

		FullLoad: fullLoad,
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
