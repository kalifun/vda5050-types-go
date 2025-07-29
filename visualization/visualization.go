package visualization

import vda5050 "github.com/kalifun/vda5050-types-go"

// AGVPosition represents the current position and orientation of an AGV
type AGVPosition struct {
	X                   float64 `json:"x"`                   // X coordinate in meters
	Y                   float64 `json:"y"`                   // Y coordinate in meters
	Theta               float64 `json:"theta"`               // Orientation in radians (-π to π)
	MapId               string  `json:"mapId"`               // ID of the current map
	PositionInitialized bool    `json:"positionInitialized"` // Position is initialized
}

// Velocity represents the current velocity of an AGV
type Velocity struct {
	Vx    float64 `json:"vx"`    // Velocity in x direction (m/s)
	Vy    float64 `json:"vy"`    // Velocity in y direction (m/s)
	Omega float64 `json:"omega"` // Angular velocity (rad/s)
}

// Visualization represents real-time visualization data for an AGV
type Visualization struct {
	vda5050.ProtocolHeader
	AgvPosition AGVPosition `json:"agvPosition"` // Current position
	Velocity    Velocity    `json:"velocity"`    // Current velocity
}
