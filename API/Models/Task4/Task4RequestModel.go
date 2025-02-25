package Task4

type RequestModel struct {
	TransformerPowerCable float64 `json:"transformerPowerCable"`
	VoltageCable          float64 `json:"voltageCable"`
	DistanceCable         float64 `json:"distanceCable"`
	MaxVoltDropCable      float64 `json:"maxVoltDropCable"`

	CircuitVoltage   float64 `json:"circuitVoltage"`
	CircuitImpedance float64 `json:"circuitImpedance"`

	VoltageSubstation float64 `json:"voltageSubstation"`
	RNormal           float64 `json:"rNormal"`
	XNormal           float64 `json:"xNormal"`
	RMin              float64 `json:"rMin"`
	XMin              float64 `json:"xMin"`
}
