package factsheet

import (
    "encoding/json"
    "strings"
    "testing"
)

func TestAgvGeometryWheelAndEnvelopes2D(t *testing.T) {
    theta := 0.5
    geo := AgvGeometry{
        WheelDefinitions: []WheelDefinition{{
            Type:            WheelDrive,
            IsActiveDriven:  true,
            IsActiveSteered: false,
            Position:        WheelPosition{X: 0.1, Y: -0.2, Theta: &theta},
            Diameter:        0.3,
            Width:           0.1,
        }},
        Envelopes2D: []Envelope2D{{
            Set:           "MECH",
            PolygonPoints: []PolygonPoint{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}},
        }},
    }
    b, err := json.Marshal(geo)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    for _, key := range []string{"wheelDefinitions", "position", "diameter", "envelopes2d", "polygonPoints"} {
        if !strings.Contains(s, "\""+key+"\"") {
            t.Fatalf("expected key %s in JSON, got: %s", key, s)
        }
    }
}

