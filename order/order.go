package order

import "github.com/kalifun/vda5050-types-go"

// Order represents a VDA5050 order message
type Order struct {
	vda5050.ProtocolHeader
	OrderId       string  `json:"orderId"`             // Unique identifier of the order
	OrderUpdateId uint32  `json:"orderUpdateId"`       // Order update identification, unique per orderId
	ZoneSetId     *string `json:"zoneSetId,omitempty"` // Optional: Unique identifier of the zone set
	Nodes         []Node  `json:"nodes"`               // List of nodes in the order
	Edges         []Edge  `json:"edges"`               // List of edges connecting the nodes
}
