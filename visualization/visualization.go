package visualization

import (
	vda5050 "github.com/kalifun/vda5050-types-go"
	"github.com/kalifun/vda5050-types-go/state"
)

// Visualization represents real-time visualization data for an AGV.
// Per docs/visualization.md, it reuses the structures from the state topic.
type Visualization struct {
	vda5050.ProtocolHeader
	AgvPosition state.AgvPosition `json:"agvPosition"` // Current position (same as state)
	Velocity    state.Velocity    `json:"velocity"`    // Current velocity (same as state)
}
