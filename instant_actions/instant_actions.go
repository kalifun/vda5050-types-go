package instant_actions

import "github.com/kalifun/vda5050-types-go"

// InstantAction represents an action to be executed immediately
type InstantAction struct {
	ActionType   string                    `json:"actionType"`                  // Type of the action
	ActionId     string                    `json:"actionId"`                    // Unique ID of the action
	BlockingType vda5050.BlockingType      `json:"blockingType"`                // How action blocks others
	Description  *string                   `json:"actionDescription,omitempty"` // Optional description
	Parameters   []vda5050.ActionParameter `json:"actionParameters,omitempty"`  // Optional parameters
}

// InstantActions represents a message containing instant actions
type InstantActions struct {
	vda5050.ProtocolHeader
	Actions []InstantAction `json:"actions"` // List of actions to execute
}
