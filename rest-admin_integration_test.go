package sdk_test

import (
	"testing"

	"github.com/grafana-tools/sdk"
)

func TestAdminOperations(t *testing.T) {
	shouldSkip(t)
	client := getClient(t)

	u := sdk.User{
		Login:          "test",
		Name:           "name",
		Email:          "grafana@sdk.com",
		Theme:          "",
		Password:       "vvvvvvvvvvvv",
		IsGrafanaAdmin: false,
	}

	st, err := client.CreateUser(u)
	if err != nil {
		t.Fatalf("failed to create an user: %s", err.Error())
	}
	if st.Message != nil && *st.Message == "failed to create user" {
		t.Fatal("failed to create an user for some unknown reason")
	}

	uid := *st.ID

	retrievedUser, err := client.GetUser(uid)
	if err != nil {
		t.Fatalf("failed to get user: %s", err.Error())
	}

	if retrievedUser.Login != u.Login || retrievedUser.Email != u.Email {
		t.Fatal("retrieved data does not match what was created")
	}

	_, err = client.UpdateUserPermissions(sdk.UserPermissions{IsGrafanaAdmin: true}, uid)
	if err != nil {
		t.Fatalf("failed to convert the user into an admin: %s", err)
	}

	retrievedUser, err = client.GetUser(uid)
	if err != nil {
		t.Fatalf("failed to get user: %s", err.Error())
	}
	if retrievedUser.IsGrafanaAdmin != true {
		t.Fatal("user should be an admin but is not")
	}
}
