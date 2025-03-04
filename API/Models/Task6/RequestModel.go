package Task6

type RequestModel struct {
	Grinding    DeviceReqModel `json:"grinding"`
	Drilled     DeviceReqModel `json:"drilled"`
	Jointing    DeviceReqModel `json:"jointing"`
	CircularSaw DeviceReqModel `json:"circularSaw"`
	Press       DeviceReqModel `json:"press"`
	Polishing   DeviceReqModel `json:"polishing"`
	Shaper      DeviceReqModel `json:"shaper"`
	Fan         DeviceReqModel `json:"fan"`
	Welding     DeviceReqModel `json:"welding"`
	Dryer       DeviceReqModel `json:"dryer"`
}
