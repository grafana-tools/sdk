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

	board.ID = 1234
	board.Title = "barfoo"

	if _, err = client.DeleteDashboard(ctx, board.UpdateSlug()); err != nil {
		t.Fatal(err)
	}

	params := sdk.SetDashboardParams{
		FolderID:  sdk.DefaultFolderId,
		Overwrite: false,
	}
	if _, err = client.SetDashboard(ctx, board, params); err != nil {
		t.Fatal(err)
	}

	if boardLinks, err = client.Search(ctx, sdk.SearchType(sdk.SearchTypeDashboard)); err != nil {
		t.Fatal(err)
	}

	if len(boardLinks) == 0 {
		t.Fatal("search query returned empty dashboard list")
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

func Test_Dashboard_CRUD_By_UID(t *testing.T) {
	var (
		err         error
		boardResult sdk.Board
	)

	shouldSkip(t)
	ctx := context.Background()
	client := getClient(t)

	var board sdk.Board

	raw, _ := ioutil.ReadFile("testdata/empty-dashboard-with-uid-2.6.json")

	if err = json.Unmarshal(raw, &board); err != nil {
		t.Fatal(err)
	}

	board.ID = 4321
	board.Title = "foobar"

	//Cleanup if Already exists
	if _, err = client.DeleteDashboardByUID(ctx, board.UID); err != nil {
		t.Fatal(err)
	}

	params := sdk.SetDashboardParams{
		FolderID:  sdk.DefaultFolderId,
		Overwrite: false,
	}
	if _, err = client.SetDashboard(ctx, board, params); err != nil {
		t.Fatal(err)
	}

	if boardResult, _, err = client.GetDashboardByUID(ctx, board.UID); err != nil {
		t.Fatal(err)
	}

	if boardResult.UID != board.UID {
		t.Fatal("Created board could not be found")
	}

	//Remove the dashboard that was created.
	if _, err = client.DeleteDashboardByUID(ctx, board.UID); err != nil {
		t.Fatal(err)
	}

	//Verify that it has been deleted
	if boardResult, _, err = client.GetDashboardByUID(ctx, board.UID); err == nil {
		t.Fatal("Failed to delete dashboard, it can still be retrieved")
	}

}

func Test_GetDashboardVersionsByDashboardID(t *testing.T) {
	var (
		board sdk.Board
		err   error
		start = sdk.QueryParamStart(0)
		limit = sdk.QueryParamLimit(10)
	)
	ctx := context.Background()
	client := getClient(t)
	raw, _ := ioutil.ReadFile("testdata/new-empty-dashboard-2.6.json")

	if err = json.Unmarshal(raw, &board); err != nil {
		t.Fatal(err)
	}
	board.UID = "1234"
	if _, err = client.DeleteDashboardByUID(ctx, board.UID); err != nil {
		t.Fatal(err)
	}

	params := sdk.SetDashboardParams{
		FolderID:  sdk.DefaultFolderId,
		Overwrite: false,
	}
	// create initial dashboard
	if _, err = client.SetDashboard(ctx, board, params); err != nil {
		t.Fatal(err)
	}

	board, _, err = client.GetDashboardByUID(ctx, board.UID)
	if err != nil {
		t.Fatal(err)
	}

	// update dashboard to create a new version
	params.Overwrite = true
	if _, err = client.SetDashboard(ctx, board, params); err != nil {
		t.Fatal(err)
	}
	board, _, err = client.GetDashboardByUID(ctx, board.UID)
	if err != nil {
		t.Fatal(err)
	}

	// fetch versions
	versions, err := client.GetDashboardVersionsByDashboardID(ctx, board.ID, start, limit)
	if err != nil {
		t.Fatal(err)
	}

	// we should have 2 versions
	if len(versions) != 2 {
		t.Fatal("dashboard should have 2 versions")
	}
	// the latest version in Grafana should match the current version of the board
	// API returns versions sorted DESC
	if versions[0].Version != board.Version {
		t.Fatal("dashboard version from dashboards/uid/:uid API does not match latest dashboard version from dashboards/id/:dashboardId/versions API")
	}

	// fetch non-existing dashboard to validate error case
	_, err = client.GetDashboardVersionsByDashboardID(ctx, 42, start, limit)
	if err == nil {
		t.Fatal("when fetching dashboard version with erroneous inputs, it should return an error, got nil error")
	}
}
