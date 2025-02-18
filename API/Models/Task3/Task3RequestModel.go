package Task3

// Task3RequestModel struct represent the request model for Task3 controller
type Task3RequestModel struct {
	EnergyPower       float64 `json:"energyPower"`
	OldErrorDeviation float64 `json:"oldErrorDeviation"`
	NewErrorDeviation float64 `json:"newErrorDeviation"`
	EnergyPrice       float64 `json:"energyPrice"`
}
