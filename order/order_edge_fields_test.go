package order

import (
    "encoding/json"
    "strings"
    "testing"
)

func TestEdgeOrientationAndRotationJSON(t *testing.T) {
    typ := OrientationTangential
    allow := true
    rs := 1.57
    e := Edge{
        EdgeId:           "e3",
        SequenceId:       3,
        StartNodeId:      "n3",
        EndNodeId:        "n4",
        Orientation:      nil,
        OrientationType:  &typ,
        RotationAllowed:  &allow,
        MaxRotationSpeed: &rs,
        Released:         true,
    }
    b, err := json.Marshal(e)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    if !strings.Contains(s, "\"orientationType\":\"TANGENTIAL\"") {
        t.Fatalf("expected orientationType TANGENTIAL, got: %s", s)
    }
    if !strings.Contains(s, "\"rotationAllowed\":true") {
        t.Fatalf("expected rotationAllowed true, got: %s", s)
    }
    if !strings.Contains(s, "\"maxRotationSpeed\":1.57") {
        t.Fatalf("expected maxRotationSpeed 1.57, got: %s", s)
    }
    if strings.Contains(s, "\"orientation\"") {
        t.Fatalf("expected orientation omitted (nil), got: %s", s)
    }
}

