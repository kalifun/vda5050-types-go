package state

import "github.com/kalifun/vda5050-types-go"

// BatteryState represents the current battery status of the AGV
type BatteryState struct {
	BatteryCharge  float64 `json:"batteryCharge"`  // Current battery charge (0.0-100.0)
	BatteryVoltage float64 `json:"batteryVoltage"` // Current battery voltage
	Charging       bool    `json:"charging"`       // Indicates if battery is being charged
	Ready          bool    `json:"ready"`          // Indicates if battery is ready for use
}

// SafetyState represents the safety-related status of the AGV
type SafetyState struct {
	EStop            EStop `json:"eStop"`            // Emergency stop active
	FieldViolation   bool  `json:"fieldViolation"`   // Protective field violation
	SafetyFieldMuted bool  `json:"safetyFieldMuted"` // Safety fields are muted
}

// EStop represents the emergency stop status of the AGV
type EStop string

const (
	EStopAutoAck EStop = "AUTOACK" // auto-acknowledgeable e-stop is activated, e.g., by bumper or protective field.
	EStopManual  EStop = "MANUAL"  // e-stop hast to be acknowledged manually at the vehicle.
	EStopRemote  EStop = "REMOTE"  // facility e-stop has to be acknowledged remotely.
	EStopNone    EStop = "NONE"    // no e-stop activated.
)

// ErrorLevel represents the severity of an error
type ErrLevel string

const (
	Warning ErrLevel = "warning"
	Error   ErrLevel = "error"
	Fatal   ErrLevel = "fatal"
)

// Error represents an error or warning message
type ErrorInfo struct {
	ErrorType       string   `json:"errorType"`                 // Type of the error
	ErrorLevel      ErrLevel `json:"errorLevel"`                // Severity level
	ErrorDesc       string   `json:"errorDesc"`                 // Description of the error
	ErrorReferences []string `json:"errorReferences,omitempty"` // References to related errors
}

// State represents the current state of an AGV
type State struct {
	vda5050.ProtocolHeader
	ActionStates          []ActionState `json:"actionStates"`          // States of all actions
	BatteryState          BatteryState  `json:"batteryState"`          // Current battery state
	OperatingMode         string        `json:"operatingMode"`         // Current operating mode
	Errors                []ErrorInfo   `json:"errors"`                // Current errors
	SafetyState           SafetyState   `json:"safetyState"`           // Current safety state
	LastNodeId            string        `json:"lastNodeId"`            // ID of last reached node
	LastNodeSequenceId    uint32        `json:"lastNodeSequenceId"`    // Sequence ID of last node
	Driving               bool          `json:"driving"`               // Indicates if AGV is driving
	Paused                bool          `json:"paused"`                // Indicates if AGV is paused
	NewBaseRequest        bool          `json:"newBaseRequest"`        // Request new base from master
	DistanceSinceLastNode float64       `json:"distanceSinceLastNode"` // Distance in meters
}
