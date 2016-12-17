package sdk

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   ॐ तारे तुत्तारे तुरे स्व
*/

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

func TestBoolInt_UnmarshalJSON_GotTrue(t *testing.T) {
	var sampleOut struct {
		Val BoolInt `json:"val"`
	}
	var sampleIn = []byte(`{"val":true}`)

	json.Unmarshal(sampleIn, &sampleOut)

	if !sampleOut.Val.Flag {
		t.Errorf("should be true but got false")
	}
	if sampleOut.Val.Value != nil {
		t.Error("int value should be empty")
	}
}

func TestBoolInt_UnmarshalJSON_GotFalse(t *testing.T) {
	var sampleOut struct {
		Val BoolInt `json:"val"`
	}
	var sampleIn = []byte(`{"val":false}`)

	json.Unmarshal(sampleIn, &sampleOut)

	if sampleOut.Val.Flag {
		t.Errorf("should be false but got true")
	}
	if sampleOut.Val.Value != nil {
		t.Error("int value should be empty")
	}
}

func TestBoolInt_UnmarshalJSON_GotInt(t *testing.T) {
	var sampleOut struct {
		Val BoolInt `json:"val"`
	}
	var sampleIn = []byte(`{"val":123456789}`)

	json.Unmarshal(sampleIn, &sampleOut)

	if sampleOut.Val.Flag {
		t.Error("should be false but got true")
	}
	if sampleOut.Val.Value == nil {
		t.Fatalf("should be 123456789 but got nil")
	}
	if *sampleOut.Val.Value != 123456789 {
		t.Errorf("should be 123456789 but got %v", sampleOut.Val.Value)
	}
}

func TestBoolInt_MarshalJSON_GotTrue(t *testing.T) {
	var sampleInp struct {
		Val BoolInt `json:"val"`
	}
	sampleInp.Val.Flag = true
	var sampleOut = []byte(`{"val":true}`)

	data, _ := json.Marshal(sampleInp)

	if bytes.Compare(data, sampleOut) != 0 {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestBoolInt_MarshalJSON_GotFalse(t *testing.T) {
	var sampleInp struct {
		Val BoolInt `json:"val"`
	}
	sampleInp.Val.Flag = false
	var sampleOut = []byte(`{"val":false}`)

	data, _ := json.Marshal(sampleInp)

	if bytes.Compare(data, sampleOut) != 0 {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestBoolInt_MarshalJSON_GotInt(t *testing.T) {
	var sampleInp struct {
		Val BoolInt `json:"val"`
	}
	var i int64 = 123456789
	sampleInp.Val.Value = &i
	var sampleOut = []byte(`{"val":123456789}`)

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

func TestGraph_SetTargetNew(t *testing.T) {
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

	graph.SetTarget(&target2)

	if len(graph.GraphPanel.Targets) != 2 {
		t.Errorf("should be 2 but %d", len(graph.GraphPanel.Targets))
	}
	if graph.GraphPanel.Targets[0].RefID != "A" {
		t.Errorf("should be equal A but %s", graph.GraphPanel.Targets[0].RefID)
	}
	if graph.GraphPanel.Targets[1].RefID != "B" {
		t.Errorf("should be equal B but %s", graph.GraphPanel.Targets[1].RefID)
	}
}

func TestGraph_SetTargetUpdate(t *testing.T) {
	var (
		target1 = Target{
			RefID:      "A",
			Datasource: "Sample Source 1",
			Expr:       "sample request 1"}
		target2 = Target{
			RefID:      "A",
			Datasource: "Sample Source 2",
			Expr:       "sample request 2"}
	)
	graph := NewGraph("")
	graph.AddTarget(&target1)

	graph.SetTarget(&target2)

	if len(graph.GraphPanel.Targets) != 1 {
		t.Errorf("should be 1 but %d", len(graph.GraphPanel.Targets))
	}
	if graph.GraphPanel.Targets[0].RefID != "A" {
		t.Errorf("should be equal A but %s", graph.GraphPanel.Targets[0].RefID)
	}
}

// Test on the panel sample with Elasticsearch datasource with Graylog query. Grafana 2.6.
func TestPanel_ElasticsearchSource_ParsedTargets(t *testing.T) {
	var rawPanel = []byte(`{
  "aliasColors": {},
  "bars": true,
  "datasource": "Example GrayLog",
  "editable": true,
  "error": false,
  "fill": 1,
  "grid": {
    "leftLogBase": 1,
    "leftMax": null,
    "leftMin": null,
    "rightLogBase": 1,
    "rightMax": null,
    "rightMin": null,
    "threshold1": null,
    "threshold1Color": "rgba(216, 200, 27, 0.27)",
    "threshold2": null,
    "threshold2Color": "rgba(234, 112, 112, 0.22)"
  },
  "id": 37,
  "isNew": true,
  "legend": {
    "avg": false,
    "current": false,
    "max": false,
    "min": false,
    "show": false,
    "total": false,
    "values": false
  },
  "lines": false,
  "linewidth": 2,
  "links": [
    {
      "params": "q=tag%3A%2Fid.*%2F+AND+level%3AERROR&rangetype=relative&relative=300#fields=message%2Csource",
      "title": "Example GrayLog Page",
      "type": "absolute",
      "url": "https://graylog/streams/xxx/messages"
    }
  ],
  "nullPointMode": "connected",
  "percentage": false,
  "pointradius": 5,
  "points": false,
  "renderer": "flot",
  "seriesOverrides": [],
  "span": 2,
  "stack": false,
  "steppedLine": false,
  "targets": [
    {
      "bucketAggs": [
        {
          "field": "timestamp",
          "id": "2",
          "settings": {
            "interval": "5m",
            "min_doc_count": 0
          },
          "type": "date_histogram"
        }
      ],
      "dsType": "elasticsearch",
      "metrics": [
        {
          "field": "select field",
          "id": "1",
          "type": "count"
        }
      ],
      "query": "tag:/.*.xxx.filtered/ AND tag:/id.*/ AND level:ERROR",
      "refId": "A",
      "target": "",
      "timeField": "timestamp"
    }
  ],
  "timeFrom": null,
  "timeShift": null,
  "title": "Example GrayLog Errors[5m]",
  "tooltip": {
    "shared": true,
    "value_type": "cumulative"
  },
  "transparent": true,
  "type": "graph",
  "x-axis": true,
  "y-axis": true,
  "y_formats": [
    "short",
    "short"
  ]
}`)

	var graph Panel
	err := json.Unmarshal(rawPanel, &graph)

	if err != nil {
		t.Fatalf("%s", err)
	}
	if len(graph.GraphPanel.Targets) != 1 {
		t.Errorf("should be 1 but %d", len(graph.GraphPanel.Targets))
	}
	if graph.GraphPanel.Targets[0].DsType == nil {
		t.Fatalf("should be \"elasticsearch\" but nil")
	}
	if *graph.GraphPanel.Targets[0].DsType != "elasticsearch" {
		t.Errorf("should be \"elasticsearch\" but %s", *graph.GraphPanel.Targets[0].DsType)
	}

}
