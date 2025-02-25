package Services

import (
	"example.com/m/Models/Task4"
	"math"
)

var cableOptions = []Task4.CableOption{
	{50.0, "Copper", 150.0, 0.386, 0.08},
	{70.0, "Copper", 190.0, 0.272, 0.075},
	{95.0, "Copper", 230.0, 0.206, 0.07},
	{120.0, "Copper", 260.0, 0.161, 0.065},
	{150.0, "Copper", 300.0, 0.129, 0.06},
	{185.0, "Copper", 340.0, 0.104, 0.055},
	{240.0, "Copper", 400.0, 0.078, 0.05},
	{300.0, "Copper", 450.0, 0.062, 0.045},
}

func SelectCable(transformerPower, voltage, distance, maxVoltageDropPercent float64) []Task4.CableOption {
	nominalCurrent := transformerPower * 1000 / (math.Sqrt(3) * voltage * 1000)
	var suitableCables []Task4.CableOption

	for _, cable := range cableOptions {
		voltageDrop := nominalCurrent * (cable.Resistance + cable.Reactance) * distance
		isCurrentOk := cable.MaxCurrent >= nominalCurrent
		isVoltageDropOk := (voltageDrop/(voltage*1000))*100 <= maxVoltageDropPercent
		if isCurrentOk && isVoltageDropOk {
			suitableCables = append(suitableCables, cable)
		}
	}
	return suitableCables
}

func calculateShortCircuitCurrent(U float64, R float64, X float64) float64 {
	Z := math.Sqrt(R*R + X*X)
	I_k := U / Z
	return I_k
}

func CalculateCurrents(voltage, rNormal, xNormal, rMin, xMin float64) Task4.Substation {
	normalState := calculateShortCircuitCurrent(voltage, rNormal, xNormal)
	minimalState := calculateShortCircuitCurrent(voltage, rMin, xMin)

	return Task4.Substation{
		NormalState:  normalState,
		MinimalState: minimalState,
		FaultState:   "Data not provided from the task",
	}
}

func ThreePhaseCurrent(voltage, impedance float64) float64 {
	return (voltage * 1000) / (math.Sqrt(3) * impedance)
}

func SinglePhaseCurrent(voltage, impedance float64) float64 {
	return (voltage * 1000) / impedance
}
