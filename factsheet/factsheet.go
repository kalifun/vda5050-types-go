package factsheet

import (
	vda5050 "github.com/kalifun/vda5050-types-go"
	"github.com/kalifun/vda5050-types-go/state"
)

// Factsheet represents the capabilities and specifications of an AGV (see docs/factsheet.md)
type Factsheet struct {
	vda5050.ProtocolHeader
	TypeSpecification  TypeSpecification  `json:"typeSpecification"`
	PhysicalParameters PhysicalParameters `json:"physicalParameters"`
	ProtocolLimits     ProtocolLimits     `json:"protocolLimits"`
	ProtocolFeatures   ProtocolFeatures   `json:"protocolFeatures"`
	AgvGeometry        AgvGeometry        `json:"agvGeometry"`
	LoadSpecification  LoadSpecification  `json:"loadSpecification"`
	VehicleConfig      *VehicleConfig     `json:"vehicleConfig,omitempty"`
}

// ----- typeSpecification -----

type AgvKinematicType string

const (
	KinematicDiff       AgvKinematicType = "DIFF"
	KinematicOmni       AgvKinematicType = "OMNI"
	KinematicThreeWheel AgvKinematicType = "THREEWHEEL"
)

type AgvClass string

const (
	ClassForklift AgvClass = "FORKLIFT"
	ClassConveyor AgvClass = "CONVEYOR"
	ClassTugger   AgvClass = "TUGGER"
	ClassCarrier  AgvClass = "CARRIER"
)

type TypeSpecification struct {
	SeriesName        string           `json:"seriesName"`
	SeriesDescription *string          `json:"seriesDescription,omitempty"`
	AgvKinematic      AgvKinematicType `json:"agvKinematic"`
	AgvClass          AgvClass         `json:"agvClass"`
	MaxLoadMass       float64          `json:"maxLoadMass"`
	LocalizationTypes []string         `json:"localizationTypes"`
	NavigationTypes   []string         `json:"navigationTypes"`
}

// ----- physicalParameters -----

type PhysicalParameters struct {
	SpeedMin        float64  `json:"speedMin"`
	SpeedMax        float64  `json:"speedMax"`
	AngularSpeedMin *float64 `json:"angularSpeedMin,omitempty"`
	AngularSpeedMax *float64 `json:"angularSpeedMax,omitempty"`
	AccelerationMax float64  `json:"accelerationMax"`
	DecelerationMax float64  `json:"decelerationMax"`
	HeightMin       float64  `json:"heightMin"`
	HeightMax       float64  `json:"heightMax"`
	Width           float64  `json:"width"`
	Length          float64  `json:"length"`
}

// ----- protocolLimits -----

type MaxStringLens struct {
	MsgLen         *uint32 `json:"msgLen,omitempty"`
	TopicSerialLen *uint32 `json:"topicSerialLen,omitempty"`
	TopicElemLen   *uint32 `json:"topicElemLen,omitempty"`
}

type ProtocolLimits struct {
    MaxStringLens   *MaxStringLens `json:"maxStringLens,omitempty"`
    MaxArrayLens    *MaxArrayLens  `json:"maxArrayLens,omitempty"`
    IdLen           *uint32        `json:"idLen,omitempty"`
    IdNumericalOnly *bool          `json:"idNumericalOnly,omitempty"`
    EnumLen         *uint32        `json:"enumLen,omitempty"`
}

// MaxArrayLens describes maximum lengths for various arrays
type MaxArrayLens struct {
    OrderNodes               *uint32 `json:"order.nodes,omitempty"`
    OrderEdges               *uint32 `json:"order.edges,omitempty"`
    TrajectoryKnotVector     *uint32 `json:"trajectory.knotVector,omitempty"`
    TrajectoryControlPoints  *uint32 `json:"trajectory.controlPoints,omitempty"`
    StateNodeStates          *uint32 `json:"state.nodeStates,omitempty"`
    StateEdgeStates          *uint32 `json:"state.edgeStates,omitempty"`
}

// ----- protocolFeatures -----

type SupportType string

const (
	SupportSupported SupportType = "SUPPORTED"
	SupportRequired  SupportType = "REQUIRED"
)

type OptionalParameter struct {
	Parameter   string      `json:"parameter"`
	Support     SupportType `json:"support"`
	Description *string     `json:"description,omitempty"`
}

type ActionScope string

const (
	ScopeInstant ActionScope = "INSTANT"
	ScopeNode    ActionScope = "NODE"
	ScopeEdge    ActionScope = "EDGE"
)

// FactsheetActionParameter describes action parameters in the factsheet (differs from order/state parameters)
type FactsheetActionParameter struct {
	Key           string  `json:"key"`
	ValueDataType string  `json:"valueDataType"`
	Description   *string `json:"description,omitempty"`
	IsOptional    *bool   `json:"isOptional,omitempty"`
}

type AgvAction struct {
	ActionType        string                     `json:"actionType"`
	ActionDescription *string                    `json:"actionDescription,omitempty"`
	ActionScopes      []ActionScope              `json:"actionScopes"`
	ActionParameters  []FactsheetActionParameter `json:"actionParameters,omitempty"`
	ResultDescription *string                    `json:"resultDescription,omitempty"`
	BlockingTypes     []vda5050.BlockingType     `json:"blockingTypes,omitempty"`
}

type ProtocolFeatures struct {
	OptionalParameters []OptionalParameter `json:"optionalParameters,omitempty"`
	AgvActions         []AgvAction         `json:"agvActions,omitempty"`
}

// ----- agvGeometry -----

type WheelType string

const (
	WheelDrive   WheelType = "DRIVE"
	WheelCaster  WheelType = "CASTER"
	WheelFixed   WheelType = "FIXED"
	WheelMecanum WheelType = "MECANUM"
)

type WheelPosition struct {
	X     float64  `json:"x"`
	Y     float64  `json:"y"`
	Theta *float64 `json:"theta,omitempty"`
}

type WheelDefinition struct {
	Type               WheelType     `json:"type"`
	IsActiveDriven     bool          `json:"isActiveDriven"`
	IsActiveSteered    bool          `json:"isActiveSteered"`
	Position           WheelPosition `json:"position"`
	Diameter           float64       `json:"diameter"`
	Width              float64       `json:"width"`
	CenterDisplacement *float64      `json:"centerDisplacement,omitempty"`
	Constraints        *string       `json:"constraints,omitempty"`
}

type PolygonPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Envelope2D struct {
	Set           string         `json:"set"`
	PolygonPoints []PolygonPoint `json:"polygonPoints"`
	Description   *string        `json:"description,omitempty"`
}

type Envelope3D struct {
	Set         string                 `json:"set"`
	Format      string                 `json:"format"`
	Data        map[string]interface{} `json:"data,omitempty"`
	URL         *string                `json:"url,omitempty"`
	Description *string                `json:"description,omitempty"`
}

type AgvGeometry struct {
	WheelDefinitions []WheelDefinition `json:"wheelDefinitions,omitempty"`
	Envelopes2D      []Envelope2D      `json:"envelopes2d,omitempty"`
	Envelopes3D      []Envelope3D      `json:"envelopes3d,omitempty"`
}

// ----- loadSpecification -----

type LoadSet struct {
	SetName               string                      `json:"setName"`
	LoadType              string                      `json:"loadType"`
	LoadPositions         []string                    `json:"loadPositions,omitempty"`
	BoundingBoxReference  *state.BoundingBoxReference `json:"boundingBoxReference,omitempty"`
	LoadDimensions        *state.LoadDimensions       `json:"loadDimensions,omitempty"`
	MaxWeight             *float64                    `json:"maxWeight,omitempty"`
	MinLoadhandlingHeight *float64                    `json:"minLoadhandlingHeight,omitempty"`
	MaxLoadhandlingHeight *float64                    `json:"maxLoadhandlingHeight,omitempty"`
	MinLoadhandlingDepth  *float64                    `json:"minLoadhandlingDepth,omitempty"`
	MaxLoadhandlingDepth  *float64                    `json:"maxLoadhandlingDepth,omitempty"`
	MinLoadhandlingTilt   *float64                    `json:"minLoadhandlingTilt,omitempty"`
	MaxLoadhandlingTilt   *float64                    `json:"maxLoadhandlingTilt,omitempty"`
	AgvSpeedLimit         *float64                    `json:"agvSpeedLimit,omitempty"`
	AgvAccelerationLimit  *float64                    `json:"agvAccelerationLimit,omitempty"`
	AgvDecelerationLimit  *float64                    `json:"agvDecelerationLimit,omitempty"`
	PickTime              *float64                    `json:"pickTime,omitempty"`
	DropTime              *float64                    `json:"dropTime,omitempty"`
	Description           *string                     `json:"description,omitempty"`
}

type LoadSpecification struct {
	LoadPositions []string  `json:"loadPositions,omitempty"`
	LoadSets      []LoadSet `json:"loadSets"`
}

// ----- vehicleConfig -----

type VersionInfo struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type NetworkInfo struct {
	DnsServers     []string `json:"dnsServers,omitempty"`
	NtpServers     []string `json:"ntpServers,omitempty"`
	LocalIpAddress *string  `json:"localIpAddress,omitempty"`
	Netmask        *string  `json:"netmask,omitempty"`
	DefaultGateway *string  `json:"defaultGateway,omitempty"`
}

type VehicleConfig struct {
	Versions []VersionInfo `json:"versions,omitempty"`
	Network  *NetworkInfo  `json:"network,omitempty"`
}
