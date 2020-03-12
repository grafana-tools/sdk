package sdk_test

import (
	"github.com/grafana-tools/sdk"
	"testing"
)

func Test_Folder_CRUD(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)

	var f = sdk.Folder{
		Title: "test-folder",
	}
	var err error

	fRecived, err := client.CreateFolder(f)
	if err != nil {
		t.Fatal(err)
	}
	if fRecived.Title != f.Title {
		t.Fatalf("got wrong title: expected %s, was %s", f.Title, fRecived.Title)
	}

	fs, err := client.GetAllFolders()
	if err != nil {
		t.Fatal(err)
	}
	if len(fs) == 0 {
		t.Fatalf("expected to get zero folders, got %#v", fs)
	}

	fg, err := client.GetFolderByUID(fRecived.UID)
	if err != nil {
		t.Fatal(err)
	}
	if fRecived.Title != fg.Title {
		t.Fatalf("got wrong title: expected %s, was %s", fRecived.Title, fg.Title)
	}

	fg, err = client.GetFolderByID(fRecived.ID)
	if err != nil {
		t.Fatal(err)
	}
	if fRecived.Title != fg.Title {
		t.Fatalf("got wrong title: expected %s, was %s", fRecived.Title, fg.Title)
	}

	fRecived.Title = "test-update-folder"
	fg, err = client.UpdateFolderByUID(fRecived)
	if err != nil {
		t.Fatal(err)
	}
	if fRecived.Title != fg.Title {
		t.Fatalf("got wrong title: expected %s, was %s", fRecived.Title, fg.Title)
	}

	r, err := client.DeleteFolderByUID(fRecived.UID)
	if err != nil {
		t.Fatal(err)
	}
	if !r {
		t.Fatal("delete failed")
	}
}
