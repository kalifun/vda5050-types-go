package state

import (
    "encoding/json"
    "strings"
    "testing"

    vda5050 "github.com/kalifun/vda5050-types-go"
)

func TestStateOptionalOmit(t *testing.T) {
    s := State{
        ProtocolHeader: vda5050.ProtocolHeader{
            HeaderId:    1,
            Timestamp:   "2025-01-01T00:00:00Z",
            Version:     "1.3.2",
            Manufacturer: "ACME",
            SerialNumber: "SN-STATE",
        },
        Maps:               []Map{},
        OrderId:            "",
        OrderUpdateId:      0,
        ZoneSetId:          nil,
        LastNodeId:         "",
        LastNodeSequenceId: 0,
        NodeStates:         []NodeState{},
        EdgeStates:         []EdgeState{},
        AgvPosition:        nil,
        Velocity:           Velocity{},
        Loads:              nil,
        Driving:            false,
        Paused:             nil,
        NewBaseRequest:     nil,
        DistanceSinceLastNode: nil,
        ActionStates:       []ActionState{},
        BatteryState:       BatteryState{BatteryCharge: 80, Charging: false},
        OperatingMode:      OperatingModeAutomatic,
        Errors:             nil,
        Information:        nil,
        SafetyState:        SafetyState{EStop: EStopNone, FieldViolation: false},
    }

    b, err := json.Marshal(s)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    out := string(b)

    // Optional fields should be omitted when nil
    for _, key := range []string{"zoneSetId", "agvPosition", "paused", "newBaseRequest", "distanceSinceLastNode"} {
        if strings.Contains(out, "\""+key+"\"") {
            t.Fatalf("expected %s to be omitted, got: %s", key, out)
        }
    }
    if !strings.Contains(out, "\"operatingMode\":\"AUTOMATIC\"") {
        t.Fatalf("expected operatingMode AUTOMATIC, got: %s", out)
    }
}

func TestActionStatusJSON(t *testing.T) {
    as := ActionState{
        ActionId:     "a1",
        ActionStatus: ActionInitializing,
    }
    b, err := json.Marshal(as)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    if !strings.Contains(s, "\"actionStatus\":\"INITIALIZING\"") {
        t.Fatalf("expected INITIALIZING status, got: %s", s)
    }
    // Ensure optional fields are omitted
    for _, key := range []string{"actionType", "actionDescription", "resultDescription"} {
        if strings.Contains(s, "\""+key+"\"") {
            t.Fatalf("expected %s omitted when nil, got: %s", key, s)
        }
    }
}

