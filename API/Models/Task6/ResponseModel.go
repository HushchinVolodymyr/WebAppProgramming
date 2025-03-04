package Task6

type ResponseModel struct {
	Grinding    DeviceRespModel `json:"grinding"`
	Drilled     DeviceRespModel `json:"drilled"`
	Jointing    DeviceRespModel `json:"jointing"`
	CircularSaw DeviceRespModel `json:"circularSaw"`
	Press       DeviceRespModel `json:"press"`
	Polishing   DeviceRespModel `json:"polishing"`
	Shaper      DeviceRespModel `json:"shaper"`
	Fan         DeviceRespModel `json:"fan"`

	Shr1 DeviceRespModel `json:"shr1"`
	Shr2 DeviceRespModel `json:"shr2"`
	Shr3 DeviceRespModel `json:"shr3"`

	Welding DeviceRespModel `json:"welding"`
	Dryer   DeviceRespModel `json:"dryer"`

	FullLoad DeviceRespModel `json:"fullLoad"`
}
