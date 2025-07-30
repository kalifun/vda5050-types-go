package connection

import "github.com/kalifun/vda5050-types-go"

// ConnectionState represents the current connection state
type ConnectionState string

const (
	Online           ConnectionState = "ONLINE"
	Offline          ConnectionState = "OFFLINE"
	ConnectionBroken ConnectionState = "CONNECTIONBROKEN"
)

// Connection represents the connection status message
type Connection struct {
	vda5050.ProtocolHeader
	ConnectionState   ConnectionState `json:"connectionState"`   // Current connection state
	Version           string          `json:"version"`           // Protocol version
	SupportedVersions []string        `json:"supportedVersions"` // List of supported versions
}

// LastWill represents the last will message for MQTT
type LastWill struct {
	vda5050.ProtocolHeader
	ConnectionState ConnectionState `json:"connectionState"` // Set to OFFLINE
}
