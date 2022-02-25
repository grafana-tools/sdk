package sdk_test

import (
	"context"
	"testing"

	"github.com/grafana-tools/sdk"
)

func Test_FolderPermissions(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)
	ctx := context.Background()

	var f = sdk.Folder{
		Title: "test-permissions-folder",
	}
	var err error

	folder, err := client.CreateFolder(ctx, f)
	if err != nil {
		t.Fatal(err)
	}
	defer func(client *sdk.Client, ctx context.Context, UID string) {
		_, _ = client.DeleteFolderByUID(ctx, UID)
	}(client, ctx, folder.UID)

	permissions, err := client.GetFolderPermissions(ctx, folder.UID)
	if err != nil {
		t.Fatal(err)
	}
	defaultPermissionCount := len(permissions)
	if defaultPermissionCount == 0 {
		t.Fatalf("Expected default permission count to be zero but was %d", defaultPermissionCount)
	}

	actualUser, err := client.GetActualUser(ctx)
	if err != nil {
		t.Fatalf("failed to get actual user: %s", err.Error())
	}

	var updatePermissions []sdk.FolderPermission
	updatePermissions = append(updatePermissions, permissions...)
	updatePermissions = append(updatePermissions, sdk.FolderPermission{
		Permission: sdk.PermissionAdmin,
		UserId:     actualUser.ID,
	})

	_, err = client.UpdateFolderPermissions(ctx, folder.UID,
		updatePermissions...,
	)
	if err != nil {
		t.Fatal(err)
	}

	updatedPermissions, err := client.GetFolderPermissions(ctx, folder.UID)
	if err != nil {
		t.Fatal(err)
	}
	updatedPermissionCount := len(updatedPermissions)
	if (defaultPermissionCount + 1) != updatedPermissionCount {
		t.Fatal("Expected permissions count to increase by 1")
	}
}
