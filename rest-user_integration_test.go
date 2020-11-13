package sdk_test

import (
	"context"
	"testing"

	"github.com/grafana-tools/sdk"
)

func Test_User_SmokeTests(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)
	ctx := context.Background()

	actualUser, err := client.GetActualUser(ctx)
	if err != nil {
		t.Fatalf("failed to get actual user: %s", err.Error())
	}

	retrievedUser, err := client.GetUser(ctx, actualUser.ID)
	if err != nil {
		t.Fatalf("failed to get user with ID %d: %s", actualUser.ID, err.Error())
	}

	if actualUser.Name != retrievedUser.Name {
		t.Fatalf("retrieved a different user: %s vs. %s", actualUser.Name, retrievedUser.Name)
	}

	allUsers, err := client.GetAllUsers(ctx)
	if err != nil {
		t.Fatalf("failed to get all users: %s", err.Error())
	}

	var found bool
	for _, u := range allUsers {
		if u.ID == retrievedUser.ID && u.Name == retrievedUser.Name {
			found = true
			break
		}
	}
	if found == false {
		t.Fatalf("failed to find an user with ID %d", actualUser.ID)
	}
}

// Test_User_SearchUsers searches for the actual user
// and plays around with pagination.
func Test_User_SearchUsers(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)
	ctx := context.Background()

	actualUser, err := client.GetActualUser(ctx)
	if err != nil {
		t.Fatalf("failed to get actual user: %s", err.Error())
	}

	q := actualUser.Login
	var currInd int = -1

	if pgUsers, err := client.SearchUsersWithPaging(ctx, nil, nil, nil); err == nil {
		for i, u := range pgUsers.Users {
			if u.Login == q {
				currInd = i
			}
		}
	} else {
		t.Fatalf("failed to search for users with paging: %s", err.Error())
	}

	if currInd == -1 {
		t.Fatal("failed to find the actual user")
	}

	// TODO(GiedriusS): add test case for index < currInd but we need to add more users first.
	// TODO(GiedriusS): add test case for querying with `q'.

	// Test that we cannot in general find that user.
	perPage := currInd + 1000
	numPage := 1
	nonExistentUser := "foobar"
	afterUsers, err := client.SearchUsersWithPaging(ctx, &nonExistentUser, &perPage, &numPage)
	if err != nil {
		t.Fatal(err)
	}

	var afterInd int = -1
	for i, u := range afterUsers.Users {
		if u.Login == q {
			afterInd = i
		}
	}
	if afterInd != -1 {
		t.Fatal("actually found the user when we were not supposed to")
	}
}

func Test_User_SwitchActualUserContext(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)
	ctx := context.Background()

	const orgName = "Test Organization"
	org := sdk.Org{
		Name: orgName,
	}

	status, err := client.CreateOrg(ctx, org)
	if err != nil {
		t.Fatalf("failed to create organization '%s': %s", orgName, err.Error())
	}

	status, err = client.SwitchActualUserContext(ctx, *status.OrgID)
	if err != nil {
		t.Fatalf("failed to switch user context to new organization: %s", err.Error())
	}

	actualOrg, err := client.GetActualOrg(ctx)
	if err != nil {
		t.Fatalf("failed to get current organization: %s", err.Error())
	}

	if actualOrg.Name != orgName {
		t.Fatalf("current organization is not the expected one. got: %s, want: %s",
			actualOrg.Name, orgName)
	}
}
