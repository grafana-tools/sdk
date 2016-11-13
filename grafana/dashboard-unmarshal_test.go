package grafana

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
	raw, _ := ioutil.ReadFile("../testdata/new-empty-dashboard-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_EmptyDashboardWithTemplating26(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("../testdata/empty-dashboard-with-templating-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_EmptyDashboardWithAnnotation26(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("../testdata/empty-dashboard-with-annotation-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_EmptyDashboardWithLinks26(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("../testdata/empty-dashboard-with-links-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_DashboardWithDefaultPanelsIn2Rows26(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("../testdata/default-panels-all-types-2-rows-dashboard-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_DashboardWithGraphWithTargets26(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("../testdata/default-panels-graph-with-targets-2.6.json")

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
		t.Errorf("panel has 2 targets but got %s", len(panel.GraphPanel.Targets))
	}
}

func TestUnmarshal_DashboardWithEmptyPanels30(t *testing.T) {
	var board Board
	raw, _ := ioutil.ReadFile("../testdata/dashboard-with-default-panels-grafana-3.0.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}
