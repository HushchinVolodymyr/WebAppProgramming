package Services

import (
	"example.com/m/Models/Task1Part1"
	"example.com/m/Models/Task1Part2"
	"fmt"
)

// Task1Part1Service - function for calculating the results of the first part of the first task
func Task1Part1Service(data Task1Part1.Task1Part1FuelComponentsModel) Task1Part1.Task1Part1ResponseModel {
	// Calculate the coefficients
	kRS := 100 / (100 - data.Moister)
	kRG := 100 / (100 - data.Moister - data.Ash)

	// Calculate the dry part
	hydrogenDry := data.Hydrogen * kRS
	carbonDry := data.Carbon * kRS
	sulfurDry := data.Sulfur * kRS
	nitrogenDry := data.Nitrogen * kRS
	oxygenDry := data.Oxygen * kRS
	ashDry := data.Ash * kRS

	// Calculate the combustible part
	hydrogenCombustible := data.Hydrogen * kRG
	carbonCombustible := data.Carbon * kRG
	sulfurCombustible := data.Sulfur * kRG
	nitrogenCombustible := data.Nitrogen * kRG
	oxygenCombustible := data.Oxygen * kRG

	// Calculate the lower heating value
	lowerHeatingValueWorking := (339*data.Carbon + 1030*data.Hydrogen - 108.8*(data.Oxygen-data.Sulfur) - 25*data.Moister) / 1000
	lowerHeatingValueDry := (lowerHeatingValueWorking + 0.025*data.Moister) * kRS
	lowerHeatingValueCombustible := (lowerHeatingValueWorking + 0.025*data.Moister) * kRG

	// Return the result
	return Task1Part1.Task1Part1ResponseModel{
		KRS:                          kRS,
		KRG:                          kRG,
		HydrogenDry:                  formatPercentage(hydrogenDry),
		CarbonDry:                    formatPercentage(carbonDry),
		SulfurDry:                    formatPercentage(sulfurDry),
		NitrogenDry:                  nitrogenDry,
		OxygenDry:                    formatPercentage(oxygenDry),
		AshDry:                       formatPercentage(ashDry),
		HydrogenCombustible:          formatPercentage(hydrogenCombustible),
		CarbonCombustible:            formatPercentage(carbonCombustible),
		SulfurCombustible:            formatPercentage(sulfurCombustible),
		NitrogenCombustible:          nitrogenCombustible,
		OxygenCombustible:            formatPercentage(oxygenCombustible),
		LowerHeatingValueWorking:     formatMJ(lowerHeatingValueWorking),
		LowerHeatingValueDry:         formatMJ(lowerHeatingValueDry),
		LowerHeatingValueCombustible: formatMJ(lowerHeatingValueCombustible),
	}
}

// formatPercentage - function for formatting the percentage
func formatPercentage(value float64) string {
	return fmt.Sprintf("%.2f%%", value)
}

// formatMJ - function for formatting the value in MJ/kg
func formatMJ(value float64) string {
	return fmt.Sprintf("%.2f МДж/кг", value)
}

// workingCalc - function for calculating the working mass of the element
func workingCalc(element float64, components Task1Part2.Task1Part2LowerHeatingComponents) float64 {
	return element * (100 - components.Moisture - components.Ash) / 100
}

// Task1Part2Service - function for calculating the results of the second part of the first task
func Task1Part2Service(data Task1Part2.Task1Part2LowerHeatingComponents) Task1Part2.Task1Part2ResponseModel {
	moistureFraction := data.Moisture / 100
	ashFraction := data.Ash / 100

	carbonWorking := workingCalc(data.Carbon, data)
	hydrogenWorking := workingCalc(data.Hydrogen, data)
	oxygenWorking := workingCalc(data.Oxygen, data)
	sulfurWorking := workingCalc(data.Sulfur, data)
	ashWorking := workingCalc(data.Ash, data)
	vanadiumWorking := workingCalc(data.Vanadium, data)

	heatingValueWorkingMass := data.LowerHeatingValue * (1 - moistureFraction) * (1 - ashFraction)

	return Task1Part2.Task1Part2ResponseModel{
		CarbonWorkingMass:            carbonWorking,
		HydrogenWorkingMass:          hydrogenWorking,
		OxygenWorkingMass:            oxygenWorking,
		SulfurWorkingMass:            sulfurWorking,
		AshWorkingMass:               ashWorking,
		VanadiumWorkingMass:          vanadiumWorking,
		LowerHeatingValueWorkingMass: heatingValueWorkingMass,
	}
}
