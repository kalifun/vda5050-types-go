package instant_actions

import "github.com/kalifun/vda5050-types-go"

// InstantActionParameter represents a parameter for an instant action
type InstantActionParameter struct {
	Key   string      `json:"key"`   // Parameter key
	Value interface{} `json:"value"` // Parameter value (can be any JSON type)
}

// InstantActions represents a message containing instant actions
type InstantActions struct {
	vda5050.ProtocolHeader
	Actions []InstantActionParameter `json:"actions"` // List of actions to execute
}
