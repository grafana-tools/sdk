package sdk_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/grafana-tools/sdk"
)

func Test_Dashboard_CRUD(t *testing.T) {
	shouldSkip(t)
	ctx := context.Background()
	client := getClient(t)

	var board sdk.Board
	raw, _ := ioutil.ReadFile("testdata/new-empty-dashboard-2.6.json")

	err := json.Unmarshal(raw, &board)
	if err != nil {
		t.Fatal(err)
	}

	// Check dashboard deletion.
	client.DeleteDashboard(ctx, board.UpdateSlug())

	// Check setting dashboard.
	params := sdk.SetDashboardParams{
		FolderID:  sdk.DefaultFolderId,
		Overwrite: false,
	}
	_, err = client.SetDashboard(ctx, board, params)
	if err != nil {
		t.Fatal(err)
	}

	// Check regular search of dashboards.
	searchBoardLinks, err := client.SearchDashboards(ctx, "", false)
	if err != nil {
		t.Fatal(err)
	}
	for _, link := range searchBoardLinks {
		// Check dashboard retrieval by UID.
		_, _, err = client.GetDashboardByUID(ctx, link.UID)
		if err != nil {
			t.Fatal(err)
		}

		// Check dashboard retrieval by Slug.
		_, _, err = client.GetDashboardBySlug(ctx, link.URI)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Check searching dashboards inside folders.
	searchInFolderBoardLinks, err := client.SearchDashboardsInFolders(ctx, []uint{0}, "", false)
	if err != nil {
		t.Fatal(err)
	}
	for _, link := range searchInFolderBoardLinks {
		// Check dashboard retrieval by UID.
		_, _, err = client.GetDashboardByUID(ctx, link.UID)
		if err != nil {
			t.Fatal(err)
		}

		// Check dashboard retrieval by Slug.
		_, _, err = client.GetDashboardBySlug(ctx, link.URI)
		if err != nil {
			t.Fatal(err)
		}
	}
}
