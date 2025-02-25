package Task4

type ResponseModel struct {
	CableOptions       []CableOption
	SubstationStates   Substation
	ThreePhaseCurrent  float64
	SinglePhaseCurrent float64
}
