package Task1Part2

// Task1Part2ResponseModel - model for Task1 Part2 Response
type Task1Part2ResponseModel struct {
	CarbonWorkingMass            float64 `json:"carbonWorkingMassFraction"`
	HydrogenWorkingMass          float64 `json:"hydrogenWorkingMassFraction"`
	OxygenWorkingMass            float64 `json:"oxygenWorkingMassFraction"`
	SulfurWorkingMass            float64 `json:"sulfurWorkingMassFraction"`
	AshWorkingMass               float64 `json:"ashWorkingMassFraction"`
	VanadiumWorkingMass          float64 `json:"vanadiumWorkingMassFraction"`
	LowerHeatingValueWorkingMass float64 `json:"lowerHeatingValueWorkingMassFraction"`
}
