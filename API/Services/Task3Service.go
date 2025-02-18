package Services

import (
	"math"
)

// CalculateImbalance function calculates the imbalance values
func CalculateProfit(averagePower, probability, energyPrice float64) float64 {
	return (averagePower * 24 * probability) * energyPrice
}

// Normal distribution function
func NormalDistribution(x, mean, stdDev float64) float64 {
	return (1 / (stdDev * math.Sqrt(2*math.Pi))) * math.Exp(-math.Pow(x-mean, 2)/(2*math.Pow(stdDev, 2)))
}

// Calculates the probability
func CalculateProbability(lowerBound, upperBound, step float64) float64 {
	integral := 0.0
	x := lowerBound

	for x < upperBound {
		y1 := NormalDistribution(x, 5.0, 1.0)
		y2 := NormalDistribution(x+step, 5.0, 1.0)
		integral += (y1 + y2) * step / 2
		x += step
	}

	return integral * 100
}
