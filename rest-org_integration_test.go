package sdk_test

import (
	"github.com/grafana-tools/sdk"
	"testing"
)

func TestCreateDelete(t *testing.T) {
	shouldSkip(t)

	client := getClient()

	oName := "coolorg"
	o := sdk.Org{Name: oName}
	statusMessage, err := client.CreateOrg(o)
	if err != nil {
		t.Fatalf("failed to create an org: %v (%s)", statusMessage, err.Error())
	}
	oID := *statusMessage.OrgID

	retrievedOrg, err := client.GetOrgById(oID)
	if err != nil {
		t.Fatalf("failed to retrieved org: %s", err.Error())
	}

	if retrievedOrg.Name != o.Name {
		t.Fatalf("got wrong org: got %s, expected %s", retrievedOrg.Name, o.Name)
	}

	_, err = client.DeleteOrg(oID)
	if err != nil {
		t.Fatalf("failed to delete org: %s", err.Error())
	}

	_, err = client.GetOrgById(oID)
	if err == nil {
		t.Fatalf("org %s is still there even though it should be deleted", o.Name)
	}
}
