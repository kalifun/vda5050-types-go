package factsheet

import (
    "encoding/json"
    "strings"
    "testing"

    vda5050 "github.com/kalifun/vda5050-types-go"
)

func TestProtocolFeaturesAgvActionsJSON(t *testing.T) {
    pf := ProtocolFeatures{
        OptionalParameters: []OptionalParameter{{Parameter: "order.nodes.nodeMarker", Support: SupportRequired}},
        AgvActions: []AgvAction{{
            ActionType:        "beep",
            ActionDescription: nil,
            ActionScopes:      []ActionScope{ScopeInstant, ScopeNode},
            ActionParameters: []FactsheetActionParameter{{
                Key:           "duration",
                ValueDataType: "NUMBER",
                IsOptional:    boolPtr(true),
            }},
            ResultDescription: nil,
            BlockingTypes:     []vda5050.BlockingType{vda5050.None, vda5050.Hard},
        }},
    }
    b, err := json.Marshal(pf)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    for _, key := range []string{"agvActions", "actionParameters", "actionScopes", "blockingTypes"} {
        if !strings.Contains(s, "\""+key+"\"") {
            t.Fatalf("expected key %s in JSON, got: %s", key, s)
        }
    }
    if !strings.Contains(s, "\"NUMBER\"") || !strings.Contains(s, "\"HARD\"") {
        t.Fatalf("expected valueDataType NUMBER and blockingTypes HARD, got: %s", s)
    }
}

func boolPtr(b bool) *bool { return &b }

