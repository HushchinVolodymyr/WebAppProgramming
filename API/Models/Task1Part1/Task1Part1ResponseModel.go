package Task1Part1

// Task1Part1ResponseModel - model for Task1 Part1 Response
type Task1Part1ResponseModel struct {
	KRS                          float64 `json:"kRS"`
	KRG                          float64 `json:"kRG"`
	HydrogenDry                  string  `json:"hydrogenDry"`
	CarbonDry                    string  `json:"carbonDry"`
	SulfurDry                    string  `json:"sulfurDry"`
	NitrogenDry                  float64 `json:"nitrogenDry"`
	OxygenDry                    string  `json:"oxygenDry"`
	AshDry                       string  `json:"ashDry"`
	HydrogenCombustible          string  `json:"hydrogenCombustible"`
	CarbonCombustible            string  `json:"carbonCombustible"`
	SulfurCombustible            string  `json:"sulfurCombustible"`
	NitrogenCombustible          float64 `json:"nitrogenCombustible"`
	OxygenCombustible            string  `json:"oxygenCombustible"`
	LowerHeatingValueWorking     string  `json:"lowerHeatingValueWorking"`
	LowerHeatingValueDry         string  `json:"lowerHeatingValueDry"`
	LowerHeatingValueCombustible string  `json:"lowerHeatingValueCombustible"`
}
