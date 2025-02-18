package Task2

type Task2RequestModel struct {
	CoalWeight    float64 `json:"coalWeight"`
	CoalHeatValue float64 `json:"coalHeatValue"`
	OilWeight     float64 `json:"oilWeight"`
	OilHeatValue  float64 `json:"oilHeatValue"`
	GasWeight     float64 `json:"gasWeight"`
	GasHeatValue  float64 `json:"gasHeatValue"`
}
