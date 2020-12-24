package mqtt

// LiveSignMessage is a message specifying a last lifetime information of
type LiveSignMessage struct {
	// TicksMs returns the ticks which have been reported by the underlying config.
	TicksMs int `json:"ticks_ms"`
	// Datetime includes the last reported date time of the livesign message.
	Datetime Time `json:"datetime"`
}
