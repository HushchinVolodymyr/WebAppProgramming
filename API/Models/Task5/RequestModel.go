package Task5

type RequestModel struct {
	EGActivity      float64 `json:"egActivity"`
	Pl              float64 `json:"pl"`
	PlLong          float64 `json:"plLong"`
	Transmission    float64 `json:"transmission"`
	Activator       float64 `json:"activator"`
	Connection      float64 `json:"connection"`
	ConnectionTimes float64 `json:"connectionTimes"`
	CostAvaria      float64 `json:"costAvaria"`
	CostPlan        float64 `json:"costPlan"`
}
