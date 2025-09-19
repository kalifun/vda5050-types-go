package state

import (
    "encoding/json"
    "strings"
    "testing"
)

func TestMapJSONOmitDescription(t *testing.T) {
    m := Map{MapId: "map1", MapVersion: "v1", MapDescription: nil, MapStatus: MapStatusEnabled}
    b, _ := json.Marshal(m)
    s := string(b)
    if strings.Contains(s, "mapDescription") {
        t.Fatalf("expected mapDescription omitted, got: %s", s)
    }
    if !strings.Contains(s, "\"mapStatus\":\"ENABLED\"") {
        t.Fatalf("expected ENABLED status, got: %s", s)
    }
}

func TestBatteryStateOptionalFields(t *testing.T) {
    bs := BatteryState{BatteryCharge: 50, Charging: true}
    b, _ := json.Marshal(bs)
    s := string(b)
    for _, k := range []string{"batteryVoltage", "batteryHealth", "reach"} {
        if strings.Contains(s, k) {
            t.Fatalf("expected %s omitted when nil, got: %s", k, s)
        }
    }
}

func TestErrorJSON(t *testing.T) {
    e := Error{ErrorType: "E1", ErrorLevel: ErrorLevelFatal, ErrorReferences: []ErrorReference{{ReferenceKey: "nodeId", ReferenceValue: "n1"}}}
    b, _ := json.Marshal(e)
    s := string(b)
    if !strings.Contains(s, "\"errorLevel\":\"FATAL\"") {
        t.Fatalf("expected FATAL errorLevel, got: %s", s)
    }
    if !strings.Contains(s, "\"referenceKey\":\"nodeId\"") {
        t.Fatalf("expected referenceKey nodeId, got: %s", s)
    }
    if strings.Contains(s, "errorDescription") || strings.Contains(s, "errorHint") {
        t.Fatalf("expected optional errorDescription/errorHint omitted, got: %s", s)
    }
}

