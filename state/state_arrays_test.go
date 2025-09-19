package state

import (
    "encoding/json"
    "strings"
    "testing"

    vda5050 "github.com/kalifun/vda5050-types-go"
)

// Note: errors/information fields do not have omitempty; nil slices serialize as null and
// empty slices serialize as []. This test documents that behavior for consumers.
func TestErrorsAndInformationNullVsEmpty(t *testing.T) {
    // Case 1: nil slices -> null
    st1 := State{
        ProtocolHeader: vda5050.ProtocolHeader{HeaderId: 1},
        Maps:           []Map{},
        OrderId:        "",
        Velocity:       Velocity{},
        Driving:        false,
        BatteryState:   BatteryState{BatteryCharge: 10, Charging: false},
        OperatingMode:  OperatingModeManual,
        SafetyState:    SafetyState{EStop: EStopNone},
    }
    b1, _ := json.Marshal(st1)
    s1 := string(b1)
    if !strings.Contains(s1, "\"errors\":null") || !strings.Contains(s1, "\"information\":null") {
        t.Fatalf("expected null for nil slices, got: %s", s1)
    }

    // Case 2: empty slices -> []
    st2 := st1
    st2.Errors = []Error{}
    st2.Information = []Info{}
    b2, _ := json.Marshal(st2)
    s2 := string(b2)
    if !strings.Contains(s2, "\"errors\":[]") || !strings.Contains(s2, "\"information\":[]") {
        t.Fatalf("expected [] for empty slices, got: %s", s2)
    }
}

