package order

import "github.com/kalifun/vda5050-types-go"

// NodePosition defines the position of a node in world coordinates
type NodePosition struct {
	X                     float64  `json:"x"`                               // X coordinate in meters
	Y                     float64  `json:"y"`                               // Y coordinate in meters
	Theta                 *float64 `json:"theta,omitempty"`                 // Orientation in radians (-π to π) (optional)
	AllowedDeviationXY    *float64 `json:"allowedDeviationXY,omitempty"`    // Maximum deviation in position in meters (optional)
	AllowedDeviationTheta *float64 `json:"allowedDeviationTheta,omitempty"` // Maximum deviation in orientation in radians (optional)
	MapId                 string   `json:"mapId"`                           // ID of the map the node is located in
	MapDescription        *string  `json:"mapDescription,omitempty"`        // Optional description of the map
}

// Action represents an action to be executed at a node or edge
type Action struct {
	ActionType        string                    `json:"actionType"`                  // Type of the action
	ActionId          string                    `json:"actionId"`                    // Unique ID of the action
	ActionDescription *string                   `json:"actionDescription,omitempty"` // Optional description of the action
	BlockingType      vda5050.BlockingType      `json:"blockingType"`                // NONE, SOFT or HARD
	Parameters        []vda5050.ActionParameter `json:"actionParameters,omitempty"`  // Optional parameters (per spec)
}

// Node represents a node in the order's path
type Node struct {
	NodeId          string        `json:"nodeId"`                    // Unique identifier of the node
	SequenceId      uint32        `json:"sequenceId"`                // Sequence number in the order
	NodePosition    *NodePosition `json:"nodePosition,omitempty"`    // Optional for some vehicle types
	Actions         []Action      `json:"actions,omitempty"`         // Actions to be executed at this node
	Released        bool          `json:"released"`                  // Indicates if node is part of base
	NodeDescription *string       `json:"nodeDescription,omitempty"` // Optional description of the node
}
