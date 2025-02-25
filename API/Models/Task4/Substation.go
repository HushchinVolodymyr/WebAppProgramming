package Task4

type Substation struct {
	NormalState  float64 `json:"normalState"`
	MinimalState float64 `json:"minimalState"`
	FaultState   string  `json:"faultState"`
}
