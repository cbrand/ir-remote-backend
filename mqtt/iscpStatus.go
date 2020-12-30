package mqtt

// IscpStatus includes the status of eiscp devices on the network which the remote is connected to
type IscpStatus struct {
	// Identifier is the identifier which can be used to send data to the remote
	Identifier string `json:"identifier"`
	// ModelName is the self advertised model name of the device
	ModelName string `json:"model_name"`
	// AreaCode represents the self advertised area code of the device
	AreaCode string `json:"area_code"`
	// DeviceCategory represents the category which the devices advertises itself as
	DeviceCategory string `json:"device_category"`
	// IscpPort is the TCP port which the device is expecting ISCP messages to be sent to
	IscpPort string `json:"iscp_port"`
}
