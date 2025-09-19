package visualization

import (
    "encoding/json"
    "strings"
    "testing"

    vda5050 "github.com/kalifun/vda5050-types-go"
    "github.com/kalifun/vda5050-types-go/state"
)

func TestVisualizationJSON(t *testing.T) {
    vx := 0.5
    vis := Visualization{
        ProtocolHeader: vda5050.ProtocolHeader{HeaderId: 1, Timestamp: "2025-01-01T00:00:00Z", Version: "1.3.2", Manufacturer: "ACME", SerialNumber: "SN-VIS"},
        AgvPosition: state.AgvPosition{PositionInitialized: true, X: 1, Y: 2, Theta: 0.1, MapId: "map-1"},
        Velocity:    state.Velocity{Vx: &vx},
    }
    b, err := json.Marshal(vis)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    if !strings.Contains(s, "\"agvPosition\"") || !strings.Contains(s, "\"velocity\"") {
        t.Fatalf("expected agvPosition and velocity keys, got: %s", s)
    }
    if strings.Contains(s, "\"vy\"") || strings.Contains(s, "\"omega\"") {
        t.Fatalf("expected vy and omega omitted when nil, got: %s", s)
    }
}

