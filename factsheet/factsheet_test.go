package factsheet

import (
	"encoding/json"
	"strings"
	"testing"

	vda5050 "github.com/kalifun/vda5050-types-go"
)

func TestFactsheetMarshalUnmarshal(t *testing.T) {
	// Prepare minimal, valid factsheet
	fs := Factsheet{
		ProtocolHeader: vda5050.ProtocolHeader{
			HeaderId:     1,
			Timestamp:    "2025-01-01T00:00:00Z",
			Version:      "1.3.2",
			Manufacturer: "ACME",
			SerialNumber: "SN-001",
		},
		TypeSpecification: TypeSpecification{
			SeriesName:        "Series-X",
			AgvKinematic:      KinematicDiff,
			AgvClass:          ClassCarrier,
			MaxLoadMass:       1000,
			LocalizationTypes: []string{"NATURAL"},
			NavigationTypes:   []string{"AUTONOMOUS"},
		},
		PhysicalParameters: PhysicalParameters{
			SpeedMin:        0.1,
			SpeedMax:        1.5,
			AccelerationMax: 0.5,
			DecelerationMax: 0.7,
			HeightMin:       0.8,
			HeightMax:       1.5,
			Width:           0.9,
			Length:          1.2,
		},
		ProtocolLimits: ProtocolLimits{},
		ProtocolFeatures: ProtocolFeatures{
			OptionalParameters: []OptionalParameter{{
				Parameter: "order.nodes.nodePosition.allowedDeviationTheta",
				Support:   SupportSupported,
			}},
		},
		AgvGeometry: AgvGeometry{},
		LoadSpecification: LoadSpecification{
			LoadSets: []LoadSet{},
		},
	}

	b, err := json.Marshal(fs)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}

	// Simple presence checks for top-level sections
	s := string(b)
	for _, key := range []string{
		"typeSpecification", "physicalParameters", "protocolLimits", "protocolFeatures", "agvGeometry", "loadSpecification",
	} {
		if !strings.Contains(s, key) {
			t.Fatalf("expected JSON to contain key %q, got: %s", key, s)
		}
	}

	// Round-trip
	var fs2 Factsheet
	if err := json.Unmarshal(b, &fs2); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	if fs2.TypeSpecification.SeriesName != fs.TypeSpecification.SeriesName {
		t.Fatalf("round-trip mismatch: got %q", fs2.TypeSpecification.SeriesName)
	}
}

func TestFactsheetMaxArrayLens(t *testing.T) {
    msgLen := uint32(1024)
    nodes := uint32(50)
    edges := uint32(60)
    knots := uint32(32)
    cps := uint32(16)
    ns := uint32(40)
    es := uint32(40)
    fs := Factsheet{
        ProtocolHeader: vda5050.ProtocolHeader{},
        TypeSpecification: TypeSpecification{
            SeriesName:        "S",
            AgvKinematic:      KinematicDiff,
            AgvClass:          ClassCarrier,
            MaxLoadMass:       1,
            LocalizationTypes: []string{"NATURAL"},
            NavigationTypes:   []string{"AUTONOMOUS"},
        },
        PhysicalParameters: PhysicalParameters{SpeedMin: 0.1, SpeedMax: 1, AccelerationMax: 1, DecelerationMax: 1, HeightMin: 1, HeightMax: 2, Width: 1, Length: 1},
        ProtocolLimits: ProtocolLimits{
            MaxStringLens: &MaxStringLens{MsgLen: &msgLen},
            MaxArrayLens:  &MaxArrayLens{OrderNodes: &nodes, OrderEdges: &edges, TrajectoryKnotVector: &knots, TrajectoryControlPoints: &cps, StateNodeStates: &ns, StateEdgeStates: &es},
        },
        ProtocolFeatures: ProtocolFeatures{},
        AgvGeometry:      AgvGeometry{},
        LoadSpecification: LoadSpecification{
            LoadSets: []LoadSet{},
        },
    }
    b, err := json.Marshal(fs)
    if err != nil {
        t.Fatalf("marshal failed: %v", err)
    }
    s := string(b)
    // Check dotted JSON keys are emitted under maxArrayLens
    for _, key := range []string{"order.nodes", "order.edges", "trajectory.knotVector", "trajectory.controlPoints", "state.nodeStates", "state.edgeStates"} {
        if !strings.Contains(s, "\""+key+"\"") {
            t.Fatalf("expected maxArrayLens to contain key %q, got: %s", key, s)
        }
    }
}
