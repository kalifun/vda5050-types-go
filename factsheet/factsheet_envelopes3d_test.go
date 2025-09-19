package factsheet

import (
    "encoding/json"
    "strings"
    "testing"
)

func TestAgvGeometryEnvelopes3DJSON(t *testing.T) {
    url := "https://example.com/model.dxf"
    desc := "3D envelope"
    geo := AgvGeometry{
        Envelopes3D: []Envelope3D{{
            Set:         "3DSET",
            Format:      "DXF",
            Data:        map[string]interface{}{"foo": "bar"},
            URL:         &url,
            Description: &desc,
        }},
    }
    b, err := json.Marshal(geo)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    for _, key := range []string{"envelopes3d", "format", "data", "url", "description"} {
        if !strings.Contains(s, "\""+key+"\"") {
            t.Fatalf("expected key %s in JSON, got: %s", key, s)
        }
    }
}

