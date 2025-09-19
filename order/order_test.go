package order

import (
    "encoding/json"
    "strings"
    "testing"

    vda5050 "github.com/kalifun/vda5050-types-go"
)

func TestTrajectoryJSONKeys(t *testing.T) {
    w := 2.5
    e := Edge{
        EdgeId:      "e1",
        SequenceId:  1,
        StartNodeId: "n1",
        EndNodeId:   "n2",
        Trajectory: &Trajectory{
            Degree:     3,
            KnotVector: []float64{0, 0, 0, 1, 1, 1},
            ControlPoints: []ControlPoint{
                {X: 0, Y: 0},
                {X: 1, Y: 1, Weight: &w},
                {X: 2, Y: 0},
            },
        },
    }

    b, err := json.Marshal(e)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    // Ensure spec keys are present and legacy names are absent
    if !strings.Contains(s, "\"knotVector\"") {
        t.Fatalf("expected JSON to contain knotVector, got: %s", s)
    }
    if !strings.Contains(s, "\"controlPoints\"") {
        t.Fatalf("expected JSON to contain controlPoints, got: %s", s)
    }
    if strings.Contains(s, "\"knots\"") {
        t.Fatalf("did not expect JSON to contain legacy 'knots', got: %s", s)
    }

    // Check weight is present for second control point, and omitted where unset
    if !strings.Contains(s, "\"weight\":2.5") {
        t.Fatalf("expected JSON to contain weight for a control point, got: %s", s)
    }
}

func TestCorridorRefPointJSON(t *testing.T) {
    ref := CorridorRefKinematicCenter
    e := Edge{
        EdgeId:      "e2",
        SequenceId:  2,
        StartNodeId: "n2",
        EndNodeId:   "n3",
        Corridor: &Corridor{
            LeftWidth:       0.5,
            RightWidth:      0.7,
            CorridorRefPoint: &ref,
        },
    }

    b, err := json.Marshal(e)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    if !strings.Contains(s, "\"corridorRefPoint\":\"KINEMATICCENTER\"") {
        t.Fatalf("expected JSON to contain corridorRefPoint, got: %s", s)
    }
}

func TestNodeActionParametersAndNodePositionOmit(t *testing.T) {
    // Build Node with optional theta omitted
    node := Node{
        NodeId:     "n1",
        SequenceId: 1,
        NodePosition: &NodePosition{
            X: 1, Y: 2, // Theta nil
        },
        Actions: []Action{{
            ActionType:   "beep",
            ActionId:     "a1",
            BlockingType: vda5050.None,
            Parameters:   []vda5050.ActionParameter{{Key: "duration", Value: 1}},
        }},
        Released: true,
    }
    b, err := json.Marshal(node)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    if !strings.Contains(s, "\"actionParameters\"") {
        t.Fatalf("expected actionParameters key in JSON, got: %s", s)
    }
    if strings.Contains(s, "\"theta\"") {
        t.Fatalf("expected theta omitted when nil, got: %s", s)
    }
}
