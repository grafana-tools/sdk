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
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestUnmarshal_NewEmptyDashboard26(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("testdata/new-empty-dashboard-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_EmptyDashboardWithTemplating26(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("testdata/empty-dashboard-with-templating-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_EmptyDashboardWithAnnotation26(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("testdata/empty-dashboard-with-annotation-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_EmptyDashboardWithLinks26(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("testdata/empty-dashboard-with-links-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_DashboardWithDefaultPanelsIn2Rows26(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("testdata/default-panels-all-types-2-rows-dashboard-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_DashboardWithGraphWithTargets26(t *testing.T) {
	var board Board
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
	if panel.OfType != GraphType {
		t.Errorf("panel type should be %d (\"graph\") type but got %d", GraphType, panel.OfType)
	}
	if *panel.Datasource != MixedSource {
		t.Errorf("panel Datasource should be \"%s\" but got \"%s\"", MixedSource, *panel.Datasource)
	}
	if len(panel.GraphPanel.Targets) != 2 {
		t.Errorf("panel has 2 targets but got %d", len(panel.GraphPanel.Targets))
	}
}

func TestUnmarshal_DashboardWithEmptyPanels30(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("testdata/dashboard-with-default-panels-grafana-3.0.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_DashboardWithHiddenTemplates(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("testdata/empty-dashboard-with-templating-4.0.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}

	if board.Templating.List[1].Hide != TemplatingHideVariable {
		t.Errorf("templating has hidden variable '%d', got %d", TemplatingHideVariable, board.Templating.List[1].Hide)
	}
}

func TestUnmarshal_DashboardWithMixedYaxes(t *testing.T) {
	var board Board
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
		t.Errorf("panel #1 has wrong min value: %d, expected: %d", min1.Value, 0)
	}
	if max1.Value != 100 || max1.Valid != true {
		t.Errorf("panel #1 has wrong max value: %d, expected: %d", max1.Value, 100)
	}

	if min2 != nil {
		t.Errorf("panel #1 has wrong min value: %v, expected: %v", min2, nil)
	}
	if max2 != nil {
		t.Errorf("panel #1 has wrong max value: %v, expected: %v", max2, nil)
	}

	if min3.Value != 0 || min3.Valid != true {
		t.Errorf("panel #2 has wrong min value: %d, expected: %d", min3.Value, 0)
	}
	if max3 != nil {
		t.Errorf("panel #2 has wrong max value: %v, expected: %v", max3, nil)
	}

	if min4 != nil {
		t.Errorf("panel #2 has wrong min value: %v, expected: %v", min4, nil)
	}
	if max4.Value != 50 || max4.Valid != true {
		t.Errorf("panel #1 has wrong max value: %d, expected: %d", max4.Value, 100)
	}
}
