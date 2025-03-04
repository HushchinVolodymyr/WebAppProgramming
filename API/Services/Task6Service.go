package Services

import (
	"example.com/m/Models/Task6"
	"math"
)

func CalcDevice(n, cos, u, count, ph, kv, tg float64) Task6.DeviceRespModel {
	np := count * ph

	npk := count * ph * kv

	npktg := count * ph * kv * tg

	np2 := count * math.Pow(ph, 2)

	i := (count * ph) / (math.Sqrt(3) * u * cos * n)

	return Task6.DeviceRespModel{
		N:     &n,
		Cos:   &cos,
		U:     &u,
		Count: &count,
		Ph:    &ph,
		Kv:    &kv,
		NP:    &np,
		Tg:    &tg,
		NPK:   &npk,
		NPKTg: &npktg,
		NP2:   &np2,
		I:     &i,
	}
}

func CalcShr(grinding, drilled, jointing, circularSaw, press, polishing, shaper, fan Task6.DeviceRespModel) Task6.DeviceRespModel {
	devices := []Task6.DeviceRespModel{grinding, drilled, jointing, circularSaw, press, polishing, shaper, fan}

	var devicesCount float64 = 0
	var npCount float64 = 0
	var npkCount float64 = 0
	var npktgCount float64 = 0
	var np2Count float64 = 0

	for _, device := range devices {
		devicesCount += *device.Count
		npCount += *device.NP
		npkCount += *device.NPK
		npktgCount += *device.NPKTg
		np2Count += *device.NP2
	}

	var kvCount float64 = npkCount / npCount
	var efficiencyQuantityCount float64 = math.Pow(npCount, 2) / np2Count
	activPowerCoefficient := 1.25
	estimatedActiveLoadCount := activPowerCoefficient * npkCount
	estimatedReactiveLoad := 1.0 * npktgCount
	fullPowerCount := math.Sqrt(math.Pow(estimatedActiveLoadCount, 2) + math.Pow(estimatedReactiveLoad, 2))
	iCount := estimatedActiveLoadCount / 0.38

	return Task6.DeviceRespModel{
		N:                     nil,
		Cos:                   nil,
		U:                     nil,
		Count:                 &devicesCount,
		Ph:                    nil,
		NP:                    &npCount,
		Kv:                    &kvCount,
		Tg:                    nil,
		NPK:                   &npkCount,
		NPKTg:                 &npktgCount,
		NP2:                   &np2Count,
		EfficiencyQuantity:    &efficiencyQuantityCount,
		ActivPowerCoefficient: &activPowerCoefficient,
		EstimatedActiveLoad:   &estimatedActiveLoadCount,
		EstimatedReactiveLoad: &estimatedReactiveLoad,
		FullPower:             &fullPowerCount,
		I:                     &iCount,
	}
}

func CalcFullLoad() Task6.DeviceRespModel {

	var deviceCount float64 = 81
	var npCount float64 = 2330
	var npkCount float64 = 752
	var npktgCount float64 = 657
	var np2Count float64 = 96388

	kvCount := npkCount / npCount
	efficiencyQuantityCount := math.Pow(npCount, 2) / np2Count
	activPowerCoefficient := 0.7
	estimatedActiveLoadCount := activPowerCoefficient * npkCount
	estimatedReactiveLoad := activPowerCoefficient * npktgCount
	fullPower := math.Sqrt(math.Pow(estimatedActiveLoadCount, 2) + math.Pow(estimatedReactiveLoad, 2))
	i := estimatedActiveLoadCount / 0.38

	return Task6.DeviceRespModel{
		N:                     nil,
		Cos:                   nil,
		U:                     nil,
		Count:                 &deviceCount,
		Ph:                    nil,
		NP:                    &npCount,
		Kv:                    &kvCount,
		Tg:                    nil,
		NPK:                   &npkCount,
		NPKTg:                 &npktgCount,
		NP2:                   &np2Count,
		EfficiencyQuantity:    &efficiencyQuantityCount,
		ActivPowerCoefficient: &activPowerCoefficient,
		EstimatedActiveLoad:   &estimatedActiveLoadCount,
		EstimatedReactiveLoad: &estimatedReactiveLoad,
		FullPower:             &fullPower,
		I:                     &i,
	}
}
