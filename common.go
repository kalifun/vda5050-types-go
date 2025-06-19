package vda5050

// ProtocolHeader is the common header structure for all messages in the VDA5050 protocol
// It contains basic identification information such as message ID, timestamp, and version
type ProtocolHeader struct {
	// HeaderId is a unique identifier for the message, must be incremented for each new message
	HeaderId uint32 `json:"headerId"`
	// Timestamp represents the message creation time in ISO8601 UTC format
	Timestamp string `json:"timestamp"`
	// Version indicates the VDA5050 protocol version being used
	Version string `json:"version"`
	// Manufacturer identifies the AGV manufacturer
	Manufacturer string `json:"manufacturer"`
	// SerialNumber is the unique serial number of the AGV
	SerialNumber string `json:"serialNumber"`
}

// BlockingType represents how an action blocks other actions
type BlockingType string

const (
	None BlockingType = "NONE" // Action doesn't block other actions
	Soft BlockingType = "SOFT" // Action blocks only other soft and hard blocking actions
	Hard BlockingType = "HARD" // Action blocks all other actions
)

type ActionParameter struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
