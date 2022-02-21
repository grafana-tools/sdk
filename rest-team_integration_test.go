package sdk_test

import (
	"context"
	"fmt"
	"github.com/grafana-tools/sdk"
	"testing"
)

func Test_Team_CRUD(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)
	ctx := context.Background()

	teamName := "mytestteam"
	teams, err := client.SearchTeamsWithPaging(ctx, &teamName, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(teams.Teams) != 0 {
		t.Fatalf("expected to get zero teams, got %#v", teams)
	}
	_, err = client.GetTeamByName(ctx, teamName)
	if err == nil {
		t.Fatal("expected request to fail for team that doesnt exist")
	}

	team := sdk.Team{
		Name:  teamName,
		Email: fmt.Sprintf("%s@test.com", teamName),
		OrgID: 1,
	}

	status, err := client.CreateTeam(ctx, team)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.GetTeam(ctx, *status.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.UpdateTeam(ctx, *status.ID, sdk.Team{
		Name:  "newTestTeam",
		Email: "",
		OrgID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.DeleteTeam(ctx, *status.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_TeamMember_CRUD(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)
	ctx := context.Background()

	teamName := "teamWithMembers"
	team := sdk.Team{
		Name:  teamName,
		Email: fmt.Sprintf("%s@test.com", teamName),
		OrgID: 1,
	}

	status, err := client.CreateTeam(ctx, team)
	if err != nil {
		t.Fatal(err)
	}

	teamMembers, err := client.GetTeamMembers(ctx, *status.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(teamMembers) != 0 {
		t.Fatalf("expected to get zero teams members, got %#v", teamMembers)
	}

	actualUser, err := client.GetActualUser(ctx)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.AddTeamMember(ctx, *status.ID, actualUser.ID)
	if err != nil {
		t.Fatal(err)
	}

	teamMembers, err = client.GetTeamMembers(ctx, *status.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(teamMembers) == 1 {
		t.Fatalf("expected to get one teams members, got %#v", teamMembers)
	}

	_, err = client.DeleteTeamMember(ctx, *status.ID, actualUser.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_TeamPreferences(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)
	ctx := context.Background()

	teamName := "teamWithPreferences"
	team := sdk.Team{
		Name:  teamName,
		Email: fmt.Sprintf("%s@test.com", teamName),
		OrgID: 1,
	}

	status, err := client.CreateTeam(ctx, team)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.GetTeamPreferences(ctx, *status.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.UpdateTeamPreferences(ctx, *status.ID, sdk.TeamPreferences{
		Theme:           "dark",
		HomeDashboardId: 0,
		Timezone:        "UTC",
	})
	if err != nil {
		t.Fatal(err)
	}
}
