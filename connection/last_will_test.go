package connection

import (
    "encoding/json"
    "strings"
    "testing"

    vda5050 "github.com/kalifun/vda5050-types-go"
)

func TestLastWillJSON(t *testing.T) {
    lw := LastWill{ProtocolHeader: vda5050.ProtocolHeader{HeaderId: 2, Timestamp: "2025-01-01T00:00:00Z", Version: "1.3.2", Manufacturer: "ACME", SerialNumber: "SN-LW"}, ConnectionState: Offline}
    b, err := json.Marshal(lw)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    if !strings.Contains(string(b), "\"connectionState\":\"OFFLINE\"") {
        t.Fatalf("expected OFFLINE last will, got: %s", string(b))
    }
}

