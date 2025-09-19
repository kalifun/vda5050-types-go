package state

// ActionStatus represents the current status of an action (per docs/state.md)
type ActionStatus string

const (
	ActionWaiting      ActionStatus = "WAITING"
	ActionInitializing ActionStatus = "INITIALIZING"
	ActionRunning      ActionStatus = "RUNNING"
	ActionPaused       ActionStatus = "PAUSED"
	ActionFinished     ActionStatus = "FINISHED"
	ActionFailed       ActionStatus = "FAILED"
)

// ActionState represents the current state of an action
type ActionState struct {
	ActionId          string       `json:"actionId"`                    // Unique ID of the action
	ActionType        *string      `json:"actionType,omitempty"`        // Type of the action (optional)
	ActionDescription *string      `json:"actionDescription,omitempty"` // Additional information on the action (optional)
	ActionStatus      ActionStatus `json:"actionStatus"`                // Current status
	ResultDescription *string      `json:"resultDescription,omitempty"` // Description of the result (optional)
}
