package state

import (
    "github.com/kalifun/vda5050-types-go"
)

// BatteryState represents the current battery status of the AGV (see state.md)
type BatteryState struct {
    BatteryCharge  float64  `json:"batteryCharge"`            // State of charge in percent
    BatteryVoltage *float64 `json:"batteryVoltage,omitempty"` // Battery voltage (optional)
    BatteryHealth  *int8    `json:"batteryHealth,omitempty"`  // Battery health [0..100] (optional)
    Charging       bool     `json:"charging"`                 // Charging in progress
    Reach          *uint32  `json:"reach,omitempty"`          // Estimated reach in meters (optional)
}

// SafetyState represents the safety-related status of the AGV
type SafetyState struct {
    EStop          EStop `json:"eStop"`          // Emergency stop status
    FieldViolation bool  `json:"fieldViolation"` // Protective field violation
}

// EStop represents the emergency stop status of the AGV
type EStop string

const (
	EStopAutoAck EStop = "AUTOACK" // auto-acknowledgeable e-stop is activated, e.g., by bumper or protective field.
	EStopManual  EStop = "MANUAL"  // e-stop hast to be acknowledged manually at the vehicle.
	EStopRemote  EStop = "REMOTE"  // facility e-stop has to be acknowledged remotely.
	EStopNone    EStop = "NONE"    // no e-stop activated.
)

// ErrorLevel represents the severity of an error (per state.md)
type ErrorLevel string

const (
    ErrorLevelWarning ErrorLevel = "WARNING"
    ErrorLevelFatal   ErrorLevel = "FATAL"
)

// ErrorReference provides context for an error (e.g., nodeId, edgeId)
type ErrorReference struct {
    ReferenceKey   string `json:"referenceKey"`
    ReferenceValue string `json:"referenceValue"`
}

// Error represents an error message
type Error struct {
    ErrorType        string           `json:"errorType"`
    ErrorReferences  []ErrorReference `json:"errorReferences,omitempty"`
    ErrorDescription *string          `json:"errorDescription,omitempty"`
    ErrorHint        *string          `json:"errorHint,omitempty"`
    ErrorLevel       ErrorLevel       `json:"errorLevel"`
}

// InfoLevel represents the severity/type of info
type InfoLevel string

const (
    InfoLevelDebug InfoLevel = "DEBUG"
    InfoLevelInfo  InfoLevel = "INFO"
)

// InfoReference provides context for an info entry
type InfoReference struct {
    ReferenceKey   string `json:"referenceKey"`
    ReferenceValue string `json:"referenceValue"`
}

// State represents the current state of an AGV
type State struct {
    vda5050.ProtocolHeader
    Maps                  []Map         `json:"maps"`                  // Array of map objects that are currently stored on the vehicle.
    OrderId               string        `json:"orderId"`               // Unique order identification of the current order or the previously finished order.
    OrderUpdateId         uint32        `json:"orderUpdateId"`         // Order update identification to identify,that an order update has been accepted by the AGV.
    ZoneSetId             string        `json:"zoneSetId"`             // Unique ID of the zone set, that the AGV currently uses for path planning. Shall be the same as the one used inthe order.
    LastNodeId            string        `json:"lastNodeId"`            // ID of last reached node
    LastNodeSequenceId    uint32        `json:"lastNodeSequenceId"`    // Sequence ID of last node
    NodeStates            []NodeState   `json:"nodeStates"`            // Array of nodeState objects that need to be traversed for fulfilling the order
    EdgeStates            []EdgeState   `json:"edgeStates"`            // Array of edgeState objects that need to be traversed for fulfilling the order
    AgvPosition           AgvPosition   `json:"agvPosition"`           // Current position of the AGV on the map.
    Velocity              Velocity      `json:"velocity"`              // The AGV velocity in vehicle coordinates.
    Loads                 []Load        `json:"loads,omitempty"`       // Loads currently handled by the AGV.
    Driving               bool          `json:"driving"`               // Indicates if AGV is driving
    Paused                bool          `json:"paused"`                // Indicates if AGV is paused
    NewBaseRequest        bool          `json:"newBaseRequest"`        // Request new base from master
    DistanceSinceLastNode float64       `json:"distanceSinceLastNode"` // Distance in meters
    ActionStates          []ActionState `json:"actionStates"`          // States of all actions
    BatteryState          BatteryState  `json:"batteryState"`          // Current battery state
    OperatingMode         OperatingMode `json:"operatingMode"`         // Current operating mode
    Errors                []Error       `json:"errors"`                // Current errors
    Information           []Info        `json:"information"`           // Array of info objects. An empty array indicates, that the AGV has no information.
    SafetyState           SafetyState   `json:"safetyState"`           // Current safety state
}

// NodeState represents the state of a node to traverse
type NodeState struct {
    NodeId          string        `json:"nodeId"`                    // Unique node identification
    SequenceId      uint32        `json:"sequenceId"`                // Sequence ID for duplicate nodeIds
    NodeDescription *string       `json:"nodeDescription,omitempty"` // Additional information on the node
    Released        bool          `json:"released"`                  // true: base, false: horizon
    NodePosition    *NodePosition `json:"nodePosition,omitempty"`    // Optional position information (from order spec)
}

// EdgeState represents the state of an edge to traverse
type EdgeState struct {
    EdgeId          string      `json:"edgeId"`                    // Unique edge identification
    SequenceId      uint32      `json:"sequenceId"`                // Sequence ID for duplicate edgeIds
    EdgeDescription *string     `json:"edgeDescription,omitempty"` // Additional information on the edge
    Released        bool        `json:"released"`                  // true: base, false: horizon
    Trajectory      *Trajectory `json:"trajectory,omitempty"`      // Optional trajectory (NURBS), see order spec
}

// NodePosition mirrors the position structure used in the order topic
type NodePosition struct {
    X               float64 `json:"x"`
    Y               float64 `json:"y"`
    Theta           float64 `json:"theta"`
    MapId           string  `json:"mapId"`
    MapDescription  *string `json:"mapDescription,omitempty"`
}

// Trajectory is a placeholder for the NURBS trajectory from the order spec
type Trajectory struct{}

// AgvPosition describes the current position of the AGV
type AgvPosition struct {
    PositionInitialized bool     `json:"positionInitialized"`           // true if position is initialized
    LocalizationScore   *float64 `json:"localizationScore,omitempty"`   // [0.0..1.0], optional
    DeviationRange      *float64 `json:"deviationRange,omitempty"`      // meters, optional
    X                   float64  `json:"x"`                             // X position in meters
    Y                   float64  `json:"y"`                             // Y position in meters
    Theta               float64  `json:"theta"`                         // Orientation in radians
    MapId               string   `json:"mapId"`                         // Map identifier
    MapDescription      *string  `json:"mapDescription,omitempty"`      // Additional info on the map
}

// Velocity describes the current velocities in vehicle coordinates
type Velocity struct {
    Vx    *float64 `json:"vx,omitempty"`    // Velocity in X (m/s), optional
    Vy    *float64 `json:"vy,omitempty"`    // Velocity in Y (m/s), optional
    Omega *float64 `json:"omega,omitempty"` // Angular velocity (rad/s), optional
}

// Load represents a currently handled load
type Load struct {
    LoadId               *string               `json:"loadId,omitempty"`               // Unique identification of the load
    LoadType             *string               `json:"loadType,omitempty"`             // Type of the load
    LoadPosition         *string               `json:"loadPosition,omitempty"`         // Which carrying unit is used
    BoundingBoxReference *BoundingBoxReference `json:"boundingBoxReference,omitempty"` // Reference point of the bounding box
    LoadDimensions       *LoadDimensions       `json:"loadDimensions,omitempty"`       // Dimensions of the load
    Weight               *float64              `json:"weight,omitempty"`               // Weight in kg
}

// BoundingBoxReference describes the reference point for a load's bounding box
type BoundingBoxReference struct {
    X     float64  `json:"x"`
    Y     float64  `json:"y"`
    Z     float64  `json:"z"`
    Theta *float64 `json:"theta,omitempty"`
}

// LoadDimensions describes the bounding box dimensions of a load
type LoadDimensions struct {
    Length float64  `json:"length"`
    Width  float64  `json:"width"`
    Height *float64 `json:"height,omitempty"`
}

// Info represents an informational message
type Info struct {
    InfoType        string          `json:"infoType"`
    InfoReferences  []InfoReference `json:"infoReferences,omitempty"`
    InfoDescription *string         `json:"infoDescription,omitempty"`
    InfoLevel       InfoLevel       `json:"infoLevel"`
}

// Map represents a map stored on the vehicle
type Map struct {
    MapId          string     `json:"mapId"`          // ID of the map describing a defined area of the vehicle's workspace.
    MapVersion     string     `json:"mapVersion"`     // Version of the map.
    MapDescription *string    `json:"mapDescription,omitempty"` // Additional information on the map.
    MapStatus      MapStatus  `json:"mapStatus"`      // Status of the map on the AGV.
}

// MapStatus defines whether the map is active on the AGV
type MapStatus string

const (
    // MapStatusEnabled indicates this map is currently active/used on the AGV.
    // At most one map with the same mapId can have its status set to 'ENABLED'.
    MapStatusEnabled MapStatus = "ENABLED"
    // MapStatusDisabled indicates this map version is currently not enabled on the AGV
    // and thus could be enabled or deleted by request.
    MapStatusDisabled MapStatus = "DISABLED"
)

// OperatingMode the AGV is in (see state.md)
type OperatingMode string

const (
    OperatingModeAutomatic     OperatingMode = "AUTOMATIC"
    OperatingModeSemiautomatic OperatingMode = "SEMIAUTOMATIC"
    OperatingModeManual        OperatingMode = "MANUAL"
    OperatingModeService       OperatingMode = "SERVICE"
    OperatingModeTeachIn       OperatingMode = "TEACHIN"
)
