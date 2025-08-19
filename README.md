# VDA5050-types-go

<div align="center">

[![](https://img.shields.io/github/v/tag/kalifun/vda5050-types-go)](https://github.com/kalifun/vda5050-types-go/tags)
[![Go Report Card](https://goreportcard.com/badge/github.com/kalifun/vda5050-types-go)](https://goreportcard.com/report/github.com/kalifun/vda5050-types-go)
[![GoDoc](https://godoc.org/github.com/kalifun/vda5050-types-go?status.svg)](https://godoc.org/github.com/kalifun/vda5050-types-go)
[![](https://img.shields.io/github/license/kalifun/vda5050-types-go)](https://github.com/kalifun/vda5050-types-go/blob/main/LICENSE)
[![](https://img.shields.io/github/last-commit/kalifun/vda5050-types-go)](https://github.com/kalifun/vda5050-types-go/commits/main)

[English](README.md) | [中文](README.zh.md)

</div>

`vda5050-types-go` is a Go package that provides a comprehensive set of data types for the VDA5050 protocol, a standard for communication between Automated Guided Vehicles (AGVs) and a master control system.

This package defines the core message structures required for AGV operations, including:

*   **Orders**: Create and manage orders with detailed node and edge definitions for AGV navigation.
*   **Factsheets**: Describe AGV capabilities, including physical dimensions, load specifications, and supported actions.
*   **Connections**: Manage connection states and handle MQTT last will messages.
*   **States**: Monitor AGV status, including battery levels, errors, and action states.
*   **Visualizations**: Define visualization information for tracking AGV movements.

The goal of this package is to provide a reliable and easy-to-use implementation of the VDA5050 standard for Go developers.



## Installation

To use `vda5050-types-go` in your project, you can use `go get`:

```bash
go get github.com/kalifun/vda5050-types-go
```


## Usage Example

Here's a simple example of how to create an `Order` and marshal it to JSON:

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kalifun/vda5050-types-go"
	"github.com/kalifun/vda5050-types-go/order"
)

func main() {
	// Create a new order
	o := order.Order{
		ProtocolHeader: vda5050.ProtocolHeader{
			HeaderId:     1,
			Timestamp:    time.Now().UTC().Format(time.RFC3339),
			Version:      "2.0.0",
			Manufacturer: "MyAGV",
			SerialNumber: "12345",
		},
		OrderId:       "order-1",
		OrderUpdateId: 1,
		Nodes: []order.Node{
			{
				NodeId:       "node-1",
				SequenceId:   1,
				Released:     true,
				Position:     &order.Position{X: 1.0, Y: 2.0, MapId: "map-1"},
				Actions:      []order.Action{},
			},
		},
		Edges: []order.Edge{},
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
```


## Core Data Types

This package provides Go structs for all major VDA5050 messages. For detailed information on each type, please refer to the [GoDoc](https://pkg.go.dev/github.com/kalifun/vda5050-types-go).

-   **`order.Order`**: Represents an order sent to an AGV, containing a sequence of nodes and edges.
-   **`factsheet.Factsheet`**: Describes the capabilities of an AGV.
-   **`state.State`**: Provides the current state of the AGV, including position, velocity, and errors.
-   **`connection.Connection`**: Manages the connection state between the AGV and the master control.
-   **`vda5050.ProtocolHeader`**: A common header included in all VDA5050 messages.


## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue if you find a bug or have a feature request.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
