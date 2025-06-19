package order

// EdgePosition represents a position on an edge
type EdgePosition struct {
	X               float64 `json:"x"`               // X coordinate in meters
	Y               float64 `json:"y"`               // Y coordinate in meters
	Orientation     float64 `json:"orientation"`     // Orientation in radians (-π to π)
	OrientationType string  `json:"orientationType"` // GLOBAL or TANGENTIAL
}

// Trajectory defines a path between two nodes
type Trajectory struct {
	Degree        uint32         `json:"degree"`        // Degree of the polynomial
	Knots         []float64      `json:"knots"`         // Vector containing the knots
	ControlPoints []EdgePosition `json:"controlPoints"` // Vector containing the control points
}

// Edge represents a connection between two nodes
type Edge struct {
	EdgeId      string      `json:"edgeId"`               // Unique identifier of the edge
	SequenceId  uint32      `json:"sequenceId"`           // Sequence number in the order
	StartNodeId string      `json:"startNodeId"`          // Start node of the edge
	EndNodeId   string      `json:"endNodeId"`            // End node of the edge
	MaxSpeed    *float64    `json:"maxSpeed,omitempty"`   // Maximum permitted speed (m/s)
	MaxHeight   *float64    `json:"maxHeight,omitempty"`  // Maximum permitted height (m)
	MinHeight   *float64    `json:"minHeight,omitempty"`  // Minimum permitted height (m)
	Trajectory  *Trajectory `json:"trajectory,omitempty"` // Optional trajectory
	Actions     []Action    `json:"actions,omitempty"`    // Actions to be executed on edge
	Released    bool        `json:"released"`             // Indicates if edge is part of base
}
