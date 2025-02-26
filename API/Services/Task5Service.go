package Services

func CalcReliable(egActivity, pl, plLong, transmission, activ, connection, conTimes float64) string {
	omegaOS := egActivity + pl*plLong + transmission + activ + connection*conTimes

	tBsOs := (egActivity*30 + (pl*plLong)*10 + transmission*100 + activ*15 + (connection*conTimes)*2) / omegaOS

	kaOs := (omegaOS * tBsOs) / 8760

	knOs := 1.2 * (43 / 8760)

	omegaDk := 2 * (kaOs*10e-4 + knOs)

	omegaDs := omegaDk + 0.02

	if omegaDs < omegaOS {
		return "Double circuit system is more reliable"
	} else if omegaDs > omegaOS {
		return "Single circuit system is more reliable"
	} else {
		return "Both systems are equally reliable"
	}
}

func CalcLosses(costAvar, costPlan float64) float64 {
	mWnedl := (0.01 * (45 * 0.01) * (5.12 * 0.01) * 6451) * 10000
	mWnedp := ((4 * 0.01) * (5.12 * 0.01) * 6451) * 10000

	loss := costAvar*mWnedl + costPlan*mWnedp

	return loss
}
