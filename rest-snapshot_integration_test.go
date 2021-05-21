package sdk_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/grafana-tools/sdk"
)

func Test_Snapshot_Create(t *testing.T) {
	shouldSkip(t)
	ctx := context.Background()
	client := getClient(t)

	var board sdk.Board

	raw, _ := ioutil.ReadFile("testdata/new-empty-dashboard-2.6.json")

	if err := json.Unmarshal(raw, &board); err != nil {
		t.Fatal(err)
	}

	snapshotRequest := sdk.CreateSnapshotRequest{
		Dashboard: board,
		Expires:   3600,
	}

	resp, err := client.CreateSnapshot(ctx, snapshotRequest)
	if err != nil {
		t.Fatal(err)
	}

	if !strings.HasPrefix(*resp.URL, "http") {
		t.Fatalf("bad url: %s", *resp.URL)
	}
}
