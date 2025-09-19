package instant_actions

import (
    "encoding/json"
    "strings"
    "testing"

    vda5050 "github.com/kalifun/vda5050-types-go"
)

func TestInstantActionParametersKey(t *testing.T) {
    msg := InstantActions{
        ProtocolHeader: vda5050.ProtocolHeader{HeaderId: 1, Timestamp: "2025-01-01T00:00:00Z", Version: "1.3.2", Manufacturer: "ACME", SerialNumber: "SN-IA"},
        Actions: []InstantAction{{
            ActionType:   "beep",
            ActionId:     "a1",
            BlockingType: vda5050.None,
            Description:  nil,
            Parameters:   []vda5050.ActionParameter{{Key: "duration", Value: 2}},
        }},
    }
    b, err := json.Marshal(msg)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    if !strings.Contains(s, "\"actionParameters\"") {
        t.Fatalf("expected actionParameters key, got: %s", s)
    }
    if strings.Contains(s, "\"parameters\"") {
        t.Fatalf("did not expect legacy parameters key, got: %s", s)
    }
}

