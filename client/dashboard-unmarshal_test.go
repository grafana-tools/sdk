package client

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/grafov/autograf/grafana"
)

func TestUnmarshal_NewEmptyDashboard26(t *testing.T) {
	var board grafana.Board
	raw, _ := ioutil.ReadFile("../testdata/new-empty-dashboard-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}

func TestUnmarshal_DashboardWithDefaultPanelsIn2Rows26(t *testing.T) {
	var board grafana.Board
	raw, _ := ioutil.ReadFile("../testdata/default-panels-all-types-2-rows-dashboard-2.6.json")

	err := json.Unmarshal(raw, &board)

	if err != nil {
		t.Error(err)
	}
}
