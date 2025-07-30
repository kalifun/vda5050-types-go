package factsheet

import "github.com/kalifun/vda5050-types-go"

// AGVGeometry represents the physical dimensions of the AGV
type AGVGeometry struct {
	Length float64 `json:"length"` // Length in meters
	Width  float64 `json:"width"`  // Width in meters
	Height float64 `json:"height"` // Height in meters
}

// LoadSpecification represents load handling capabilities
type LoadSpecification struct {
	LoadType  string  `json:"loadType"`  // Type of load that can be handled
	MaxWeight float64 `json:"maxWeight"` // Maximum weight in kg
	MaxLength float64 `json:"maxLength"` // Maximum length in meters
	MaxWidth  float64 `json:"maxWidth"`  // Maximum width in meters
	MaxHeight float64 `json:"maxHeight"` // Maximum height in meters
}

// ActionParameter represents a parameter for an action
type ActionParameter struct {
	Key           string `json:"key"`           // Parameter key
	ValueDataType string `json:"valueDataType"` // Parameter type (string, number, boolean)
	Description   string `json:"description"`   // Parameter description
	IsOptional    bool   `json:"isOptional"`    // Whether parameter is optional
}

// ActionDefinition represents an action supported by the AGV
type ActionDefinition struct {
	ActionType        string            `json:"actionType"`        // Type of the action
	ActionDescription string            `json:"actionDescription"` // Description
	Parameters        []ActionParameter `json:"parameters"`        // List of parameters
	AllowedScopes     []string          `json:"allowedScopes"`     // Where action can be used
}

// Factsheet represents the capabilities and specifications of an AGV
type Factsheet struct {
	vda5050.ProtocolHeader
	AGVGeometry       AGVGeometry        `json:"agvGeometry"`       // Physical dimensions
	LoadSpecification LoadSpecification  `json:"loadSpecification"` // Load capabilities
	Actions           []ActionDefinition `json:"actions"`           // Supported actions
	MaxSpeed          float64            `json:"maxSpeed"`          // Maximum speed in m/s
	MaxRotationSpeed  float64            `json:"maxRotationSpeed"`  // Maximum rotation speed in rad/s
	MinRotationRadius float64            `json:"minRotationRadius"` // Minimum rotation radius in m
}
