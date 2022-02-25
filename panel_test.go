package sdk_test

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>
   Copyright 2016-2019 The Grafana SDK authors

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

	"github.com/grafana-tools/sdk"
)

func TestStackVal_UnmarshalJSON_GotTrue(t *testing.T) {
	var sampleOut struct {
		Val sdk.BoolString `json:"val"`
	}
	var sampleIn = []byte(`{"val":true}`)

	if err := json.Unmarshal(sampleIn, &sampleOut); err != nil {
		t.Fatal(err)
	}

	if !sampleOut.Val.Flag {
		t.Errorf("should be true but got false")
	}
	if sampleOut.Val.Value != "" {
		t.Error("string value should be empty")
	}
}

func TestStackVal_UnmarshalJSON_GotFalse(t *testing.T) {
	var sampleOut struct {
		Val sdk.BoolString `json:"val"`
	}
	var sampleIn = []byte(`{"val":false}`)

	if err := json.Unmarshal(sampleIn, &sampleOut); err != nil {
		t.Fatal(err)
	}

	if sampleOut.Val.Flag {
		t.Errorf("should be false but got true")
	}
	if sampleOut.Val.Value != "" {
		t.Error("string value should be empty")
	}
}

func TestStackVal_UnmarshalJSON_GotString(t *testing.T) {
	var sampleOut struct {
		Val sdk.BoolString `json:"val"`
	}
	var sampleIn = []byte(`{"val":"A"}`)

	if err := json.Unmarshal(sampleIn, &sampleOut); err != nil {
		t.Fatal(err)
	}

	if sampleOut.Val.Flag {
		t.Error("should be false but got true")
	}
	if sampleOut.Val.Value != "A" {
		t.Errorf("should be 'A' but got '%s'", sampleOut.Val.Value)
	}
}

func TestStackVal_MarshalJSON_GotTrue(t *testing.T) {
	var sampleInp struct {
		Val sdk.BoolString `json:"val"`
	}
	sampleInp.Val.Flag = true
	var sampleOut = []byte(`{"val":true}`)

	data, err := json.Marshal(sampleInp)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, sampleOut) {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestStackVal_MarshalJSON_GotFalse(t *testing.T) {
	var sampleInp struct {
		Val sdk.BoolString `json:"val"`
	}
	sampleInp.Val.Flag = false
	var sampleOut = []byte(`{"val":false}`)

	data, err := json.Marshal(sampleInp)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, sampleOut) {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestStackVal_MarshalJSON_GotString(t *testing.T) {
	var sampleInp struct {
		Val sdk.BoolString `json:"val"`
	}
	sampleInp.Val.Value = "A"
	var sampleOut = []byte(`{"val":"A"}`)

	data, err := json.Marshal(sampleInp)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, sampleOut) {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestBoolInt_UnmarshalJSON_GotTrue(t *testing.T) {
	var sampleOut struct {
		Val sdk.BoolInt `json:"val"`
	}
	var sampleIn = []byte(`{"val":true}`)

	if err := json.Unmarshal(sampleIn, &sampleOut); err != nil {
		t.Fatal(err)
	}

	if !sampleOut.Val.Flag {
		t.Errorf("should be true but got false")
	}
	if sampleOut.Val.Value != nil {
		t.Error("int value should be empty")
	}
}

func TestBoolInt_UnmarshalJSON_GotFalse(t *testing.T) {
	var sampleOut struct {
		Val sdk.BoolInt `json:"val"`
	}
	var sampleIn = []byte(`{"val":false}`)

	if err := json.Unmarshal(sampleIn, &sampleOut); err != nil {
		t.Fatal(err)
	}

	if sampleOut.Val.Flag {
		t.Errorf("should be false but got true")
	}
	if sampleOut.Val.Value != nil {
		t.Error("int value should be empty")
	}
}

func TestBoolInt_UnmarshalJSON_GotInt(t *testing.T) {
	var sampleOut struct {
		Val sdk.BoolInt `json:"val"`
	}
	var sampleIn = []byte(`{"val":123456789}`)

	if err := json.Unmarshal(sampleIn, &sampleOut); err != nil {
		t.Fatal(err)
	}

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
		Val sdk.BoolInt `json:"val"`
	}
	sampleInp.Val.Flag = true
	var sampleOut = []byte(`{"val":true}`)

	data, err := json.Marshal(sampleInp)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, sampleOut) {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestBoolInt_MarshalJSON_GotFalse(t *testing.T) {
	var sampleInp struct {
		Val sdk.BoolInt `json:"val"`
	}
	sampleInp.Val.Flag = false
	var sampleOut = []byte(`{"val":false}`)

	data, err := json.Marshal(sampleInp)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, sampleOut) {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestBoolInt_MarshalJSON_GotInt(t *testing.T) {
	var sampleInp struct {
		Val sdk.BoolInt `json:"val"`
	}
	var i int64 = 123456789
	sampleInp.Val.Value = &i
	var sampleOut = []byte(`{"val":123456789}`)

	data, err := json.Marshal(sampleInp)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, sampleOut) {
		t.Errorf("should be %s but got %s", sampleOut, data)
	}
}

func TestNewGraph(t *testing.T) {
	var title = "Sample Title"

	graph := sdk.NewGraph(title)

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

func TestNewTimeseries(t *testing.T) {
	var title = "Sample Title"

	timeseries := sdk.NewTimeseries(title)

	if timeseries.TimeseriesPanel == nil {
		t.Error("should be not nil")
	}
	if timeseries.GraphPanel != nil {
		t.Error("should be nil")
	}
	if timeseries.TextPanel != nil {
		t.Error("should be nil")
	}
	if timeseries.DashlistPanel != nil {
		t.Error("should be nil")
	}
	if timeseries.SinglestatPanel != nil {
		t.Error("should be nil")
	}
	if timeseries.Title != title {
		t.Errorf("title should be %s but %s", title, timeseries.Title)
	}
}

func TestGraph_AddTarget(t *testing.T) {
	var target = sdk.Target{
		RefID:      "A",
		Datasource: "Sample Source",
		Expr:       "sample request"}
	graph := sdk.NewGraph("")

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
		target1 = sdk.Target{
			RefID:      "A",
			Datasource: "Sample Source 1",
			Expr:       "sample request 1"}
		target2 = sdk.Target{
			RefID:      "B",
			Datasource: "Sample Source 2",
			Expr:       "sample request 2"}
	)
	graph := sdk.NewGraph("")
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
		target1 = sdk.Target{
			RefID:      "A",
			Datasource: "Sample Source 1",
			Expr:       "sample request 1"}
		target2 = sdk.Target{
			RefID:      "A",
			Datasource: "Sample Source 2",
			Expr:       "sample request 2"}
	)
	graph := sdk.NewGraph("")
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

	var graph sdk.Panel
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

// Test on the panel sample with stackdriver datasource.
func TestPanel_Stackdriver_ParsedTargets(t *testing.T) {
	var rawPanel = []byte(`{
  "aliasColors": {},
  "bars": false,
  "dashLength": 10,
  "dashes": false,
  "datasource": "awesome-stackdriver",
  "fill": 1,
  "gridPos": {
	"h": 8,
	"w": 12,
	"x": 0,
	"y": 0
  },
  "id": 2,
  "legend": {
	"avg": false,
	"current": false,
	"max": false,
	"min": false,
	"show": true,
	"total": false,
	"values": false
  },
  "lines": true,
  "linewidth": 1,
  "links": [],
  "nullPointMode": "null",
  "options": {},
  "percentage": false,
  "pointradius": 2,
  "points": false,
  "renderer": "flot",
  "seriesOverrides": [],
  "spaceLength": 10,
  "stack": false,
  "steppedLine": false,
  "targets": [
	{
	  "aliasBy": "",
	  "alignOptions": [
		{
		  "expanded": true,
		  "label": "Alignment options",
		  "options": [
			{
			  "label": "delta",
			  "metricKinds": [
				"CUMULATIVE",
				"DELTA"
			  ],
			  "text": "delta",
			  "value": "ALIGN_DELTA",
			  "valueTypes": [
				"INT64",
				"DOUBLE",
				"MONEY",
				"DISTRIBUTION"
			  ]
			},
			{
			  "label": "rate",
			  "metricKinds": [
				"CUMULATIVE",
				"DELTA"
			  ],
			  "text": "rate",
			  "value": "ALIGN_RATE",
			  "valueTypes": [
				"INT64",
				"DOUBLE",
				"MONEY"
			  ]
			},
			{
			  "label": "min",
			  "metricKinds": [
				"GAUGE",
				"DELTA"
			  ],
			  "text": "min",
			  "value": "ALIGN_MIN",
			  "valueTypes": [
				"INT64",
				"DOUBLE",
				"MONEY"
			  ]
			},
			{
			  "label": "max",
			  "metricKinds": [
				"GAUGE",
				"DELTA"
			  ],
			  "text": "max",
			  "value": "ALIGN_MAX",
			  "valueTypes": [
				"INT64",
				"DOUBLE",
				"MONEY"
			  ]
			},
			{
			  "label": "mean",
			  "metricKinds": [
				"GAUGE",
				"DELTA"
			  ],
			  "text": "mean",
			  "value": "ALIGN_MEAN",
			  "valueTypes": [
				"INT64",
				"DOUBLE",
				"MONEY"
			  ]
			},
			{
			  "label": "count",
			  "metricKinds": [
				"GAUGE",
				"DELTA"
			  ],
			  "text": "count",
			  "value": "ALIGN_COUNT",
			  "valueTypes": [
				"INT64",
				"DOUBLE",
				"MONEY",
				"BOOL"
			  ]
			},
			{
			  "label": "sum",
			  "metricKinds": [
				"GAUGE",
				"DELTA"
			  ],
			  "text": "sum",
			  "value": "ALIGN_SUM",
			  "valueTypes": [
				"INT64",
				"DOUBLE",
				"MONEY",
				"DISTRIBUTION"
			  ]
			},
			{
			  "label": "stddev",
			  "metricKinds": [
				"GAUGE",
				"DELTA"
			  ],
			  "text": "stddev",
			  "value": "ALIGN_STDDEV",
			  "valueTypes": [
				"INT64",
				"DOUBLE",
				"MONEY"
			  ]
			},
			{
			  "label": "percent change",
			  "metricKinds": [
				"GAUGE",
				"DELTA"
			  ],
			  "text": "percent change",
			  "value": "ALIGN_PERCENT_CHANGE",
			  "valueTypes": [
				"INT64",
				"DOUBLE",
				"MONEY"
			  ]
			}
		  ]
    }
	  ],
	  "alignmentPeriod": "stackdriver-auto",
	  "crossSeriesReducer": "REDUCE_MEAN",
	  "defaultProject": "loading project...",
	  "filters": [
		"resource.label.subscription_id",
		"=",
		"some_subscription_id"
	  ],
	  "groupBy": [],
	  "metricKind": "DELTA",
	  "metricType": "pubsub.googleapis.com/subscription/ack_message_count",
	  "perSeriesAligner": "ALIGN_DELTA",
	  "refId": "A",
	  "unit": "1",
	  "usedAlignmentPeriod": 60,
	  "valueType": "INT64"
	}
  ],
  "thresholds": [],
  "timeFrom": null,
  "timeRegions": [],
  "timeShift": null,
  "title": "Pubsub Ack msg count",
  "tooltip": {
	"shared": true,
	"sort": 0,
	"value_type": "individual"
  },
  "type": "graph",
  "xaxis": {
	"buckets": null,
	"mode": "time",
	"name": null,
	"show": true,
	"values": []
  },
  "yaxes": [
	{
	  "format": "short",
	  "label": null,
	  "logBase": 1,
	  "max": null,
	  "min": null,
	  "show": true
	},
	{
	  "format": "short",
	  "label": null,
	  "logBase": 1,
	  "max": null,
	  "min": null,
	  "show": true
	}
  ],
  "yaxis": {
	"align": false,
	"alignLevel": null
  }
}`)

	var graph sdk.Panel
	err := json.Unmarshal(rawPanel, &graph)

	if err != nil {
		t.Fatalf("%s", err)
	}
	if len(graph.GraphPanel.Targets) != 1 {
		t.Fatalf("should be 1 but %d", len(graph.GraphPanel.Targets))
	}
	if graph.GraphPanel.Targets[0].MetricType != "pubsub.googleapis.com/subscription/ack_message_count" {
		t.Fatalf("should be \"pubsub.googleapis.com/subscription/ack_message_count\" but is not")
	}
}

func TestPanel_Timeseries(t *testing.T) {
	var rawPanel = []byte(`{
		"id": 2,
		"gridPos": {
		  "x": 0,
		  "y": 0,
		  "w": 12,
		  "h": 9
		},
		"type": "timeseries",
		"title": "Panel Title",
		"options": {
		"tooltip": {
		  "mode": "single"
		},
		"legend": {
		  "displayMode": "list",
		  "placement": "bottom",
		  "calcs": []
		}
		},
		"fieldConfig": {
		"defaults": {
		  "custom": {
		  "drawStyle": "line",
		  "lineInterpolation": "linear",
		  "barAlignment": 0,
		  "lineWidth": 1,
		  "fillOpacity": 0,
		  "gradientMode": "none",
		  "spanNulls": false,
		  "showPoints": "auto",
		  "pointSize": 5,
		  "stacking": {
			"mode": "none",
			"group": "A"
		  },
		  "axisPlacement": "auto",
		  "axisLabel": "",
		  "scaleDistribution": {
			"type": "linear"
		  },
		  "hideFrom": {
			"tooltip": false,
			"viz": false,
			"legend": false
		  },
		  "thresholdsStyle": {
			"mode": "off"
		  }
		  },
		  "color": {
		  "mode": "palette-classic"
		  },
		  "thresholds": {
		  "mode": "absolute",
		  "steps": [
			{
			"value": 0.1,
			"color": "green"
			},
			{
			"value": 80,
			"color": "red"
			}
		  ]
		  },
		  "mappings": []
		},
		"overrides": []
		},
		"targets": [
		  {
			"expr": "test_expr",
			"legendFormat": "",
			"interval": "",
			"exemplar": true,
			"refId": "A",
			"datasource": "Sample datasource"
		  }
		],
		"datasource": null
	}`)
	var timeseries sdk.Panel
	err := json.Unmarshal(rawPanel, &timeseries)

	if err != nil {
		t.Fatalf("%s", err)
	}

	if len(timeseries.TimeseriesPanel.Targets) != 1 {
		t.Fatalf("should be 1 but %d", len(timeseries.TimeseriesPanel.Targets))
	}
}

// TestCustomPanelOutput_MarshalJSON marshals new custom panel to JSON,
// then marshals that json to map[string]interface{},\
// and then checks both custom and non-custom keys are present and correct.
func TestCustomPanelOutput_MarshalJSON(t *testing.T) {
	var (
		titleKey    = "title"
		titleValue  = "test title"
		customKey   = "test_key"
		customValue = "bar_value"
	)
	p := sdk.NewCustom(titleValue)
	custom := map[string]interface{}(*p.CustomPanel)
	custom[customKey] = customValue
	b, err := json.Marshal(p)
	if err != nil {
		t.Fatalf("failed to marshal custom panel: %v", err)
	}
	var j = make(map[string]interface{})
	err = json.Unmarshal(b, &j)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	val, ok := j[customKey]
	if !ok {
		t.Fatalf("failed to find key %s in map %v", customKey, j)
	}
	if val != customValue {
		t.Fatalf("wrong value of %s: got %s, expected %s", customKey, val, customValue)
	}
	val, ok = j[titleKey]
	if !ok {
		t.Fatalf("failed to find key %s in map %v", titleKey, j)
	}
	if val != titleValue {
		t.Fatalf("wrong value of %s: got %s, expected %s", titleKey, val, titleValue)
	}

}
