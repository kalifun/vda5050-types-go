package instant_actions

import "github.com/kalifun/vda5050-types-go"

// InstantAction represents an action to be executed immediately
type InstantAction struct {
	ActionType   string               `json:"actionType"`                  // Type of the action
	ActionId     string               `json:"actionId"`                    // Unique ID of the action
	BlockingType vda5050.BlockingType `json:"blockingType"`                // How action blocks others
	Description  *string              `json:"actionDescription,omitempty"` // Optional description
	Parameters   []ActionParameter    `json:"actionParameters,omitempty"`  // Optional parameters
}

// InstantActionParameter represents a parameter for an instant action
type ActionParameter struct {
	Key   string      `json:"key"`   // Parameter key
	Value interface{} `json:"value"` // Parameter value (can be any JSON type)
}

// InstantActions represents a message containing instant actions
type InstantActions struct {
	vda5050.ProtocolHeader
	Actions []InstantAction `json:"actions"` // List of actions to execute
}
