package state

// ActionStatus represents the current status of an action
type ActionStatus string

const (
	Waiting  ActionStatus = "waiting"  // Action has not started
	Running  ActionStatus = "running"  // Action is being executed
	Paused   ActionStatus = "paused"   // Action is paused
	Finished ActionStatus = "finished" // Action completed successfully
	Failed   ActionStatus = "failed"   // Action failed to complete
)

// ActionState represents the current state of an action
type ActionState struct {
	ActionId     string       `json:"actionId"`             // Unique ID of the action
	ActionType   string       `json:"actionType"`           // Type of the action
	ActionStatus ActionStatus `json:"actionStatus"`         // Current status
	ResultDesc   *string      `json:"resultDesc,omitempty"` // Description of the result
	ErrorDesc    *string      `json:"errorDesc,omitempty"`  // Description of error if failed
}
