package instant_actions

import "github.com/kalifun/vda5050"

// InstantAction represents an action to be executed immediately
type InstantAction struct {
	ActionType   string               `json:"actionType"`            // Type of the action
	ActionId     string               `json:"actionId"`              // Unique ID of the action
	BlockingType vda5050.BlockingType `json:"blockingType"`          // How action blocks others
	Description  *string              `json:"description,omitempty"` // Optional description
	Parameters   map[string]string    `json:"parameters,omitempty"`  // Optional parameters
}

// InstantActions represents a message containing instant actions
type InstantActions struct {
	vda5050.ProtocolHeader
	Actions []InstantAction `json:"actions"` // List of actions to execute
}
