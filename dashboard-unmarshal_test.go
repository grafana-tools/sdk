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
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/grafana-tools/sdk"
)

func TestUnmarshal_NewEmptyDashboard26(t *testing.T) {
	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/new-empty-dashboard-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_EmptyDashboardWithTemplating26(t *testing.T) {
	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/empty-dashboard-with-templating-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_EmptyDashboardWithAnnotation26(t *testing.T) {
	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/empty-dashboard-with-annotation-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_EmptyDashboardWithLinks26(t *testing.T) {
	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/empty-dashboard-with-links-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_DashboardWithDefaultPanelsIn2Rows26(t *testing.T) {
	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/default-panels-all-types-2-rows-dashboard-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_DashboardWithGraphWithTargets26(t *testing.T) {
	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/default-panels-graph-with-targets-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
	if len(board.Rows) != 1 {
		t.Errorf("there are 1 row defined but got %d", len(board.Rows))
	}
	if len(board.Rows[0].Panels) != 1 {
		t.Errorf("there are 1 panel defined but got %d", len(board.Rows[0].Panels))
	}
	panel := board.Rows[0].Panels[0]
	if panel.OfType != sdk.GraphType {
		t.Errorf("panel type should be %d (\"graph\") type but got %d", sdk.GraphType, panel.OfType)
	}
	if panel.Datasource != sdk.MixedSource {
		t.Errorf("panel Datasource should be \"%s\" but got \"%s\"", sdk.MixedSource, panel.Datasource)
	}
	if len(panel.GraphPanel.Targets) != 2 {
		t.Errorf("panel has 2 targets but got %d", len(panel.GraphPanel.Targets))
	}
	if len(panel.GraphPanel.Targets[1].Tags) != 1 {
		t.Fatalf("should be 1 but got %d", len(panel.GraphPanel.Targets[0].Tags))
	}
	var tag = panel.GraphPanel.Targets[1].Tags[0]

	if tag.Key != "key" && tag.Operator != "=" && tag.Value != "testvalue" {
		t.Fatalf("Unexpected Target Tags: got %s", tag)
	}

}

func TestUnmarshal_DashboardWithEmptyPanels30(t *testing.T) {
	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/dashboard-with-default-panels-grafana-3.0.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_DashboardWithHiddenTemplates(t *testing.T) {
	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/empty-dashboard-with-templating-4.0.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}

	if board.Templating.List[1].Hide != sdk.TemplatingHideVariable {
		t.Errorf("templating has hidden variable '%d', got %d", sdk.TemplatingHideVariable, board.Templating.List[1].Hide)
	}
}

func TestUnmarshal_DashboardWithMixedYaxes(t *testing.T) {
	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/dashboard-with-panels-with-mixed-yaxes.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}

	p1, p2 := board.Rows[0].Panels[0], board.Rows[0].Panels[1]
	min1, max1 := p1.Yaxes[0].Min, p1.Yaxes[0].Max
	min2, max2 := p1.Yaxes[1].Min, p1.Yaxes[1].Max
	min3, max3 := p2.Yaxes[0].Min, p2.Yaxes[0].Max
	min4, max4 := p2.Yaxes[1].Min, p2.Yaxes[1].Max

	if min1.Value != 0 || min1.Valid != true {
		t.Errorf("panel #1 has wrong min value: %f, expected: %f", min1.Value, 0.0)
	}
	if max1.Value != 100 || max1.Valid != true {
		t.Errorf("panel #1 has wrong max value: %f, expected: %f", max1.Value, 100.0)
	}

	if min2 != nil {
		t.Errorf("panel #1 has wrong min value: %v, expected: %v", min2, nil)
	}
	if max2 != nil {
		t.Errorf("panel #1 has wrong max value: %v, expected: %v", max2, nil)
	}

	if min3.Value != 0 || min3.Valid != true {
		t.Errorf("panel #2 has wrong min value: %f, expected: %f", min3.Value, 0.0)
	}
	if max3 != nil {
		t.Errorf("panel #2 has wrong max value: %v, expected: %v", max3, nil)
	}

	if min4 != nil {
		t.Errorf("panel #2 has wrong min value: %v, expected: %v", min4, nil)
	}
	if max4.Value != 50 || max4.Valid != true {
		t.Errorf("panel #1 has wrong max value: %f, expected: %f", max4.Value, 100.0)
	}
}

func TestUnmarshal_DashboardWithGraphWithTargets83(t *testing.T) {
	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/default-panels-graph-with-target-8.3.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Fatal(err)
	}
	if len(board.Panels) != 2 {
		t.Fatalf("board should have 2 panels but got %d", len(board.Panels))
	}
	rowPanel := board.Panels[0]
	if rowPanel.OfType != sdk.RowType {
		t.Errorf("panel type should be %d (\"row\") type but got %d", sdk.GraphType, rowPanel.OfType)
	}

	panel := board.Panels[1]
	if panel.OfType != sdk.GraphType {
		t.Errorf("panel type should be %d (\"graph\") type but got %d", sdk.GraphType, panel.OfType)
	}

	if len(panel.GraphPanel.Targets) != 1 {
		t.Errorf("panel has 1 targets but got %d", len(panel.GraphPanel.Targets))
	}

	target := panel.GraphPanel.Targets[0]
	datasource, ok := target.Datasource.(map[string]interface{})
	if !ok {
		t.Fatalf("target Datasource should be a map but got %T", panel.Datasource)
	}
	if datasource["type"] != "prometheus" {
		t.Errorf("target datasource should be of type \"prometheus\" but got %s", datasource["type"])
	}

}
