package Task2

type Task2ResponseModel struct {
	CoalEmissionFactor float64 `json:"coalEmissionFactor"`
	CoalEmission       float64 `json:"coalEmission"`
	OilEmisionFactor   float64 `json:"oilEmisionFactor"`
	OilEmission        float64 `json:"oilEmission"`
	GasEmisionFactor   float64 `json:"gasEmisionFactor"`
	GasEmission        float64 `json:"gasEmission"`
	TotalEmission      float64 `json:"totalEmission"`
}
