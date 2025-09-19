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
	ConnectionState ConnectionState `json:"connectionState"` // Current connection state
}

// LastWill represents the last will message for MQTT
type LastWill struct {
    vda5050.ProtocolHeader
    ConnectionState ConnectionState `json:"connectionState"` // Set to OFFLINE
}

// NewConnection creates a connection message with the given state
func NewConnection(header vda5050.ProtocolHeader, state ConnectionState) *Connection {
    return &Connection{
        ProtocolHeader: header,
        ConnectionState: state,
    }
}

// NewLastWillOffline creates a last-will message preset to OFFLINE
func NewLastWillOffline(header vda5050.ProtocolHeader) *LastWill {
    return &LastWill{
        ProtocolHeader: header,
        ConnectionState: Offline,
    }
}
