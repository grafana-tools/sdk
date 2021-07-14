package sdk_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/grafana-tools/sdk"
)

func Test_Alerts_Read(t *testing.T) {
	shouldSkip(t)
	client := getClient(t)
	ctx := context.Background()

	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/dashboard-with-alerts.json")
	if err := json.Unmarshal(raw, &board); err != nil {
		t.Fatal(err)
	}

	if _, err := client.DeleteDashboard(ctx, board.UpdateSlug()); err != nil {
		t.Fatal(err)
	}
	if _, err := client.SetDashboard(ctx, board, sdk.SetDashboardParams{
		FolderID:  sdk.DefaultFolderId,
		Overwrite: false,
	}); err != nil {
		t.Fatal(err)
	}

	alerts, err := client.SearchAlerts(ctx, sdk.SearchDashboardID(int(board.ID)))
	if err != nil {
		t.Fatal(err)
	}

	if len(alerts) == 0 {
		t.Fatal(err)
	}
	const sampleAlertName = "Sample alert"
	var sampleAlertFound bool
	for _, a := range alerts {
		if a.Name == sampleAlertName {
			sampleAlertFound = true
			break
		}
	}
	if !sampleAlertFound {
		t.Fatal(err)
	}
}
