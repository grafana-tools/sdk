package grafana

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestStackVal_UnmarshalJSON_GotTrue(t *testing.T) {
	var sampleOut struct {
		Val BoolString `json:"val"`
	}
	var sampleIn = []byte(`{"val":true}`)

	json.Unmarshal(sampleIn, &sampleOut)

	if !sampleOut.Val.Flag {
		t.Errorf("should be true but got false")
	}
	if sampleOut.Val.Value != "" {
		t.Error("string value should be empty")
	}
}

func TestStackVal_UnmarshalJSON_GotFalse(t *testing.T) {
	var sampleOut struct {
		Val BoolString `json:"val"`
	}
	var sampleIn = []byte(`{"val":false}`)

	json.Unmarshal(sampleIn, &sampleOut)

	if sampleOut.Val.Flag {
		t.Errorf("should be false but got true")
	}
	if sampleOut.Val.Value != "" {
		t.Error("string value should be empty")
	}
}

func TestStackVal_UnmarshalJSON_GotString(t *testing.T) {
	var sampleOut struct {
		Val BoolString `json:"val"`
	}
	var sampleIn = []byte(`{"val":"A"}`)

	json.Unmarshal(sampleIn, &sampleOut)

	if sampleOut.Val.Flag {
		t.Error("should be false but got true")
	}
	if sampleOut.Val.Value != "A" {
		t.Errorf("should be 'A' but got '%s'", sampleOut.Val.Value)
	}
}

func TestStackVal_MarshalJSON_GotTrue(t *testing.T) {
	var sampleInp struct {
		Val BoolString `json:"val"`
	}
	sampleInp.Val.Flag = true
	var sampleOut = []byte(`{"val":true}`)

	data, _ := json.Marshal(sampleInp)

	if bytes.Compare(data, sampleOut) != 0 {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestStackVal_MarshalJSON_GotFalse(t *testing.T) {
	var sampleInp struct {
		Val BoolString `json:"val"`
	}
	sampleInp.Val.Flag = false
	var sampleOut = []byte(`{"val":false}`)

	data, _ := json.Marshal(sampleInp)

	if bytes.Compare(data, sampleOut) != 0 {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestStackVal_MarshalJSON_GotString(t *testing.T) {
	var sampleInp struct {
		Val BoolString `json:"val"`
	}
	sampleInp.Val.Value = "A"
	var sampleOut = []byte(`{"val":"A"}`)

	data, _ := json.Marshal(sampleInp)

	if bytes.Compare(data, sampleOut) != 0 {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestNewGraph(t *testing.T) {
	var title = "Sample Title"

	graph := NewGraph(title)

	if graph.GraphPanel == nil {
		t.Error("should be not nil")
	}
	if graph.TextPanel != nil {
		t.Error("should be nil")
	}
	if graph.DashlistPanel != nil {
		t.Error("should be nil")
	}
	if graph.SinglestatPanel != nil {
		t.Error("should be nil")
	}
	if graph.Title != title {
		t.Errorf("title should be %s but %s", title, graph.Title)
	}
}

func TestGraph_AddTarget(t *testing.T) {
	var target = Target{
		RefID:      "A",
		Datasource: "Sample Source",
		Expr:       "sample request"}
	graph := NewGraph("")

	graph.AddTarget(&target)

	if len(graph.GraphPanel.Targets) != 1 {
		t.Errorf("should be 1 but %d", len(graph.GraphPanel.Targets))
	}
	if graph.GraphPanel.Targets[0].RefID != "A" {
		t.Errorf("should be equal A but %s", graph.GraphPanel.Targets[0].RefID)
	}
}

func TestGraph_SetTarget(t *testing.T) {
	var (
		target1 = Target{
			RefID:      "A",
			Datasource: "Sample Source 1",
			Expr:       "sample request 1"}
		target2 = Target{
			RefID:      "B",
			Datasource: "Sample Source 2",
			Expr:       "sample request 2"}
	)
	graph := NewGraph("")
	graph.AddTarget(&target1)

	graph.AddTarget(&target2)

	if len(graph.GraphPanel.Targets) != 1 {
		t.Errorf("should be 1 but %d", len(graph.GraphPanel.Targets))
	}
	if graph.GraphPanel.Targets[0].RefID != "B" {
		t.Errorf("should be equal B but %s", graph.GraphPanel.Targets[0].RefID)
	}
}
