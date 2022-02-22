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
	teams, err := client.SearchTeams(ctx,
		sdk.WithQuery(teamName),
		sdk.WithPagesize(20),
		sdk.WithPage(1),
	)
	if err != nil {
		t.Fatal(err)
	}
	if len(teams.Teams) != 0 {
		t.Fatalf("expected to get zero teams, got %#v", teams)
	}

	team := sdk.Team{
		Name:  teamName,
		Email: fmt.Sprintf("%s@test.com", teamName),
		OrgID: 1,
	}

	_, err = client.CreateTeam(ctx, team)
	if err != nil {
		t.Fatal(err)
	}

	teamByName, err := client.GetTeamByName(ctx, teamName)
	if err != nil {
		t.Fatal(err)
	}
	teamId := teamByName.ID

	_, err = client.GetTeam(ctx, teamId)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.UpdateTeam(ctx, teamId, sdk.Team{
		Name:  "newTestTeam",
		Email: "",
		OrgID: 1,
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.DeleteTeam(ctx, teamId)
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

	_, err := client.CreateTeam(ctx, team)
	if err != nil {
		t.Fatal(err)
	}

	teamByName, err := client.GetTeamByName(ctx, teamName)
	if err != nil {
		t.Fatal(err)
	}
	teamId := teamByName.ID

	teamMembers, err := client.GetTeamMembers(ctx, teamId)
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

	_, err = client.AddTeamMember(ctx, teamId, actualUser.ID)
	if err != nil {
		t.Fatal(err)
	}

	teamMembers, err = client.GetTeamMembers(ctx, teamId)
	if err != nil {
		t.Fatal(err)
	}
	if len(teamMembers) != 1 {
		t.Fatalf("expected to get one teams members, got %#v", teamMembers)
	}

	_, err = client.DeleteTeamMember(ctx, teamId, actualUser.ID)
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

	_, err := client.CreateTeam(ctx, team)
	if err != nil {
		t.Fatal(err)
	}

	teamByName, err := client.GetTeamByName(ctx, teamName)
	if err != nil {
		t.Fatal(err)
	}
	teamId := teamByName.ID

	_, err = client.GetTeamPreferences(ctx, teamId)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.UpdateTeamPreferences(ctx, teamId, sdk.TeamPreferences{
		Theme:           "dark",
		HomeDashboardId: 0,
		Timezone:        "UTC",
	})
	if err != nil {
		t.Fatal(err)
	}
}
