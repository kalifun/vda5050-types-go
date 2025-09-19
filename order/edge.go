package order

// Trajectory defines a path between two nodes
// ControlPoint defines a NURBS control point
type ControlPoint struct {
	X      float64  `json:"x"`
	Y      float64  `json:"y"`
	Weight *float64 `json:"weight,omitempty"`
}

// Trajectory defines a path between two nodes as NURBS
type Trajectory struct {
	Degree        float64        `json:"degree"`        // Degree of the NURBS curve (spec: float64)
	KnotVector    []float64      `json:"knotVector"`    // Knot vector
	ControlPoints []ControlPoint `json:"controlPoints"` // Control points including start and end
}

// Edge represents a connection between two nodes
type Edge struct {
	EdgeId           string           `json:"edgeId"`                     // Unique identifier of the edge
	SequenceId       uint32           `json:"sequenceId"`                 // Sequence number in the order
	StartNodeId      string           `json:"startNodeId"`                // Start node of the edge
	EndNodeId        string           `json:"endNodeId"`                  // End node of the edge
	MaxSpeed         *float64         `json:"maxSpeed,omitempty"`         // Maximum permitted speed (m/s)
	MaxHeight        *float64         `json:"maxHeight,omitempty"`        // Maximum permitted height (m)
	MinHeight        *float64         `json:"minHeight,omitempty"`        // Minimum permitted height (m)
	Orientation      *float64         `json:"orientation,omitempty"`      // Desired orientation on the edge (rad)
	OrientationType  *OrientationType `json:"orientationType,omitempty"`  // GLOBAL or TANGENTIAL
	RotationAllowed  *bool            `json:"rotationAllowed,omitempty"`  // Whether rotation is allowed on the edge
	MaxRotationSpeed *float64         `json:"maxRotationSpeed,omitempty"` // Max rotation speed (rad/s)
	Trajectory       *Trajectory      `json:"trajectory,omitempty"`       // Optional trajectory
	Corridor         *Corridor        `json:"corridor,omitempty"`         // Optional corridor bounds
	Actions          []Action         `json:"actions,omitempty"`          // Actions to be executed on edge
	Released         bool             `json:"released"`                   // Indicates if edge is part of base
}

// OrientationType defines the frame in which orientation is interpreted
type OrientationType string

const (
	OrientationGlobal     OrientationType = "GLOBAL"
	OrientationTangential OrientationType = "TANGENTIAL"
)

// Corridor defines lateral deviation bounds relative to the trajectory
type Corridor struct {
    LeftWidth  float64 `json:"leftWidth"`
    RightWidth float64 `json:"rightWidth"`
    CorridorRefPoint *CorridorRefPoint `json:"corridorRefPoint,omitempty"`
}

// CorridorRefPoint defines whether corridor bounds relate to kinematic center or contour
type CorridorRefPoint string

const (
    CorridorRefKinematicCenter CorridorRefPoint = "KINEMATICCENTER"
    CorridorRefContour         CorridorRefPoint = "CONTOUR"
)
