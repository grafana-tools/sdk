package sdk_test

import (
	"context"
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
	ctx := context.Background()
	client := getClient(t)

	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/new-empty-dashboard-2.6.json")

	if err = json.Unmarshal(raw, &board); err != nil {
		t.Fatal(err)
	}

	client.DeleteDashboard(ctx, board.UpdateSlug())
	params := sdk.SetDashboardParams{
		FolderID:  sdk.DefaultFolderId,
		Overwrite: false,
	}
	if _, err = client.SetDashboard(ctx, board, params); err != nil {
		t.Fatal(err)
	}

	if boardLinks, err = client.SearchDashboards(ctx, "", false); err != nil {
		t.Fatal(err)
	}

	if boardLinks, err = client.SearchDashboardsInFolders(ctx, []uint{1, 2}, "", false); err != nil {
		t.Fatal(err)
	}

	for _, link := range boardLinks {
		_, _, err = client.GetDashboardByUID(ctx, link.UID)
		if err != nil {
			t.Fatal(err)
		}

		_, _, err = client.GetDashboardBySlug(ctx, link.URI)
		if err != nil {
			t.Fatal(err)
		}
	}
}
