package sdk_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/grafana-tools/sdk"
)

func Test_Dashboard_CRUD(t *testing.T) {
	var (
		boardLinks []sdk.FoundBoard
		err        error
	)

	shouldSkip(t)

	client := getClient(t)

	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/new-empty-dashboard-2.6.json")

	if err = json.Unmarshal(raw, &board); err != nil {
		t.Fatal(err)
	}

	client.DeleteDashboard(board.UpdateSlug())
	params := sdk.SetDashboardParams{
		FolderID:  sdk.DefaultFolderId,
		Overwrite: false,
	}
	if _, err = client.SetDashboard(board, params); err != nil {
		t.Fatal(err)
	}

	if boardLinks, err = client.SearchDashboards("", false); err != nil {
		t.Fatal(err)
	}

	for _, link := range boardLinks {
		_, _, err = client.GetDashboardByUID(link.UID)
		if err != nil {
			t.Fatal(err)
		}

		_, _, err = client.GetDashboardBySlug(link.URI)
		if err != nil {
			t.Fatal(err)
		}
	}
}
