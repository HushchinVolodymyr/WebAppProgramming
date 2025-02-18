package Services

// CalculateCoalEmission calculates the emission of coal
func CalculateCoalEmission(kgCoal, heatValueCoal float64) float64 {
	return 1e-6 * GetCoalEmissionFactor() * heatValueCoal * kgCoal
}

// CalculateOilEmission calculates the emission of oil
func CalculateOilEmission(kgOil, heatValueOil float64) float64 {
	return 1e-6 * GetOilEmissionFactor() * heatValueOil * kgOil
}

// CalculateGasEmission calculates the emission of gas
func CalculateGasEmission(kgGas float64, heatValueGas float64) float64 {
	return 1e-6 * GetGasEmissionFactor() * heatValueGas * kgGas
}

// CalculateTotalEmission calculates the total emission
func CalculateTotalEmission(coalEmission, oilEmission, gasEmission float64) float64 {
	return coalEmission + oilEmission + gasEmission
}

// GetCoalEmissionFactor returns the emission factor of coal
func GetCoalEmissionFactor() float64 {
	return 150.0
}

// GetOilEmissionFactor returns the emission factor of oil
func GetOilEmissionFactor() float64 {
	return 0.57
}

// GetGasEmissionFactor returns the emission factor of gas
func GetGasEmissionFactor() float64 {
	return 0.0
}
