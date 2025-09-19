package connection

import (
    "encoding/json"
    "strings"
    "testing"

    vda5050 "github.com/kalifun/vda5050-types-go"
)

func TestConnectionJSONShape(t *testing.T) {
    c := Connection{
        ProtocolHeader: vda5050.ProtocolHeader{
            HeaderId:    10,
            Timestamp:   "2025-01-01T00:00:00Z",
            Version:     "1.3.2",
            Manufacturer: "ACME",
            SerialNumber: "SN-CONN",
        },
        ConnectionState: Online,
    }
    b, err := json.Marshal(c)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    if !strings.Contains(s, "\"connectionState\":\"ONLINE\"") {
        t.Fatalf("expected ONLINE connectionState, got: %s", s)
    }
    if strings.Contains(s, "supportedVersions") {
        t.Fatalf("did not expect supportedVersions in JSON, got: %s", s)
    }
}

