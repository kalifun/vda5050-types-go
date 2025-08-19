# VDA5050-types-go

<div align="center">

[![](https://img.shields.io/github/v/tag/kalifun/vda5050-types-go)](https://github.com/kalifun/vda5050-types-go/tags)
[![Go Report Card](https://goreportcard.com/badge/github.com/kalifun/vda5050-types-go)](https://goreportcard.com/report/github.com/kalifun/vda5050-types-go)
[![GoDoc](https://godoc.org/github.com/kalifun/vda5050-types-go?status.svg)](https://godoc.org/github.com/kalifun/vda5050-types-go)
[![](https://img.shields.io/github/license/kalifun/vda5050-types-go)](https://github.com/kalifun/vda5050-types-go/blob/main/LICENSE)
[![](https://img.shields.io/github/last-commit/kalifun/vda5050-types-go)](https://github.com/kalifun/vda5050-types-go/commits/main)

[English](README.md) | [中文](README.zh.md)

</div>

`vda5050-types-go` 是一个 Go 语言包，为 VDA5050 协议提供了一套完整的数据类型。VDA5050 是用于自动导引车（AGV）与主控系统之间通信的标准。

该包定义了 AGV 操作所需的核心消息结构，包括：

*   **订单 (Orders)**：创建和管理订单，包含用于 AGV 导航的详细节点和边定义。
*   **概况 (Factsheets)**：描述 AGV 的能力，包括物理尺寸、负载规格和支持的动作。
*   **连接 (Connections)**：管理连接状态并处理 MQTT 的遗嘱消息。
*   **状态 (States)**：监控 AGV 状态，包括电池电量、错误和动作状态。
*   **可视化 (Visualizations)**：定义用于跟踪 AGV 运动的可视化信息。

该包旨在为 Go 开发者提供一个可靠且易于使用的 VDA5050 标准实现。

## 安装

在您的项目中使用 `vda5050-types-go`，您可以使用 `go get` 命令：

```bash
go get github.com/kalifun/vda5050-types-go
```

## 使用示例

以下是一个如何创建 `Order` 对象并将其编组为 JSON 的简单示例：

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
	// 创建一个新订单
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

	// 编组为 JSON
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonData))
}
```

## 核心数据类型

该包为所有主要的 VDA5050 消息提供了 Go 结构体。有关每种类型的详细信息，请参阅 [GoDoc](https://pkg.go.dev/github.com/kalifun/vda5050-types-go)。

-   **`order.Order`**: 表示发送给 AGV 的订单，包含一系列节点和边。
-   **`factsheet.Factsheet`**: 描述 AGV 的能力。
-   **`state.State`**: 提供 AGV 的当前状态，包括位置、速度和错误。
-   **`connection.Connection`**: 管理 AGV 与主控系统之间的连接状态。
-   **`vda5050.ProtocolHeader`**: 所有 VDA5050 消息中包含的通用报头。

## 贡献

欢迎贡献！如果您发现错误或有功能请求，请随时提交拉取请求或开启一个 issue。

## 许可证

该项目根据 MIT 许可证授权 - 有关详细信息，请参阅 [LICENSE](LICENSE) 文件。