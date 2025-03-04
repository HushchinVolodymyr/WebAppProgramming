package Task6

type DeviceRespModel struct {
	N                     *float64 `json:"n"`
	Cos                   *float64 `json:"cos"`
	U                     *float64 `json:"u"`
	Count                 *float64 `json:"count"`
	Ph                    *float64 `json:"ph"`
	NP                    *float64 `json:"np"`
	Kv                    *float64 `json:"kv"`
	Tg                    *float64 `json:"tg"`
	NPK                   *float64 `json:"npk"`
	NPKTg                 *float64 `json:"npktg"`
	NP2                   *float64 `json:"np2"`
	EfficiencyQuantity    *float64 `json:"efficiencyQuantity"`
	ActivPowerCoefficient *float64 `json:"activPowerCoefficient"`
	EstimatedActiveLoad   *float64 `json:"estimatedActiveLoad"`
	EstimatedReactiveLoad *float64 `json:"estimatedReactiveLoad"`
	FullPower             *float64 `json:"fullPower"`
	I                     *float64 `json:"i"`
}
