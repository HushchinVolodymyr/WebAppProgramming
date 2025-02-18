package Task3

type Task3ResponseModel struct {
	ProbabilityImbalance float64 `json:"probabilityImbalance"`
	ProfitImbalance      float64 `json:"profitImbalance"`
	PeniyaImbalance      float64 `json:"peniyaImbalance"`
	TotalImbalance       float64 `json:"totalImbalance"`
	ProbabilityBalance   float64 `json:"probabilityBalance"`
	ProfitBalance        float64 `json:"profitBalance"`
	PeniyaBalance        float64 `json:"peniyaBalance"`
	TotalBalance         float64 `json:"totalBalance"`
}
