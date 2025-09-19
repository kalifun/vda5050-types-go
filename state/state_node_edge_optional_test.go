package state

import (
    "encoding/json"
    "strings"
    "testing"

    "github.com/kalifun/vda5050-types-go/order"
)

func TestNodeStateOmitNodePositionWhenNil(t *testing.T) {
    ns := NodeState{NodeId: "n1", SequenceId: 1, Released: true, NodePosition: nil}
    b, _ := json.Marshal(ns)
    if strings.Contains(string(b), "nodePosition") {
        t.Fatalf("expected nodePosition omitted when nil, got: %s", string(b))
    }
}

func TestEdgeStateOmitTrajectoryWhenNil(t *testing.T) {
    es := EdgeState{EdgeId: "e1", SequenceId: 1, Released: true, Trajectory: nil}
    b, _ := json.Marshal(es)
    if strings.Contains(string(b), "trajectory") {
        t.Fatalf("expected trajectory omitted when nil, got: %s", string(b))
    }
}

func TestEdgeStateWithTrajectoryJSON(t *testing.T) {
    tr := &order.Trajectory{Degree: 1, KnotVector: []float64{0, 0, 1, 1}, ControlPoints: []order.ControlPoint{{X: 0, Y: 0}, {X: 1, Y: 0}}}
    es := EdgeState{EdgeId: "e2", SequenceId: 2, Released: true, Trajectory: tr}
    b, _ := json.Marshal(es)
    s := string(b)
    if !strings.Contains(s, "knotVector") || !strings.Contains(s, "controlPoints") {
        t.Fatalf("expected trajectory keys present, got: %s", s)
    }
}

