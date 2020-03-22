package sdk_test

import (
	"context"
	"testing"

	"github.com/grafana-tools/sdk"
)

func Test_Folder_CRUD(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)
	ctx := context.Background()

	var f1 = sdk.Folder{
		Title: "test-folder-1",
	}
	var err error

	fReceived1, err := client.CreateFolder(ctx, f1)
	if err != nil {
		t.Fatal(err)
	}
	if fReceived1.Title != f1.Title {
		t.Fatalf("got wrong title: expected %s, was %s", f1.Title, fReceived1.Title)
	}

	var f2 = sdk.Folder{
		Title: "test-folder-2",
	}

	fReceived2, err := client.CreateFolder(ctx, f2)
	if err != nil {
		t.Fatal(err)
	}
	if fReceived2.Title != f2.Title {
		t.Fatalf("got wrong title: expected %s, was %s", f2.Title, fReceived2.Title)
	}

	fs, err := client.GetAllFolders(ctx, sdk.Limit(1))
	if err != nil {
		t.Fatal(err)
	}
	if len(fs) != 1 {
		t.Fatalf("expected to get one folders, got %d", len(fs))
	}

	fg, err := client.GetFolderByUID(ctx, fReceived1.UID)
	if err != nil {
		t.Fatal(err)
	}
	if fReceived1.Title != fg.Title {
		t.Fatalf("got wrong title: expected %s, was %s", fReceived1.Title, fg.Title)
	}

	fg, err = client.GetFolderByID(ctx, fReceived1.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fReceived1.Title != fg.Title {
		t.Fatalf("got wrong title: expected %s, was %s", fReceived1.Title, fg.Title)
	}

	fReceived1.Title = "test-update-folder"
	fg, err = client.UpdateFolderByUID(ctx, fReceived1)
	if err != nil {
		t.Fatal(err)
	}
	if fReceived1.Title != fg.Title {
		t.Fatalf("got wrong title: expected %s, was %s", fReceived1.Title, fg.Title)
	}

	r, err := client.DeleteFolderByUID(ctx, fReceived1.UID)
	if err != nil {
		t.Fatal(err)
	}
	if !r {
		t.Fatal("delete failed")
	}
	r, err = client.DeleteFolderByUID(ctx, fReceived2.UID)
	if err != nil {
		t.Fatal(err)
	}
	if !r {
		t.Fatal("delete failed")
	}
}
