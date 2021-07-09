package sdk_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/grafana-tools/sdk"
)

func TestCreateDelete(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)
	ctx := context.Background()

	oName := "coolorg"
	o := sdk.Org{Name: oName}
	statusMessage, err := client.CreateOrg(ctx, o)
	if err != nil {
		t.Fatalf("failed to create an org: %v (%s)", statusMessage, err.Error())
	}
	t.Logf("got status message: %v\n", statusMessage)

	oID := *statusMessage.OrgID

	retrievedOrg, err := client.GetOrgById(ctx, oID)
	if err != nil {
		t.Fatalf("failed to retrieved org: %s", err.Error())
	}

	if retrievedOrg.Name != o.Name {
		t.Fatalf("got wrong org: got %s, expected %s", retrievedOrg.Name, o.Name)
	}

	_, err = client.DeleteOrg(ctx, oID)
	if err != nil {
		t.Fatalf("failed to delete org: %s", err.Error())
	}

	_, err = client.GetOrgById(ctx, oID)
	if err == nil {
		t.Fatalf("org %s is still there even though it should be deleted", o.Name)
	}
}

// TestUpdateOrgAddress checks if updating Org address works correctly
func TestUpdateOrgAddress(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)
	ctx := context.Background()

	// Create a new organization
	oName := "coolorg"
	o := sdk.Org{Name: oName}
	statusMessage, err := client.CreateOrg(ctx, o)
	if err != nil {
		t.Fatalf("failed to create an org: %v (%s)", statusMessage, err.Error())
	}
	oID := *statusMessage.OrgID

	// Test if updating organization by ID works as expected
	// Create a dummy address object
	address := sdk.Address{
		Address1: "CoolAddress1",
		Address2: "CoolAddress2",
		City:     "CoolCity",
		ZipCode:  "CoolZipCode",
		State:    "CoolState",
		Country:  "CoolCountry",
	}

	// Try updating organization address by Org ID
	statusMessage, err = client.UpdateOrgAddress(ctx, address, oID)
	if err != nil {
		t.Fatalf("failed to update the address: %v (%s)", statusMessage, err.Error())
	}

	retrievedOrg, err := client.GetOrgById(ctx, oID)
	if err != nil {
		t.Fatalf("failed to retrieved org: %s", err.Error())
	}

	// Check if retrieved address values are equal to expected ones
	if !reflect.DeepEqual(retrievedOrg.Address, address) {
		t.Fatalf("got wrong address: got %+v, expected %+v", retrievedOrg.Address, address)
	}

	// Test if updating current organization works as expected
	address = sdk.Address{
		Address1: "NiceAddress1",
		Address2: "NiceAddress2",
		City:     "NiceCity",
		ZipCode:  "NiceZipCode",
		State:    "NiceState",
		Country:  "NiceCountry",
	}

	statusMessage, err = client.SwitchActualUserContext(ctx, oID)
	if err != nil {
		t.Fatalf("failed to switch user context: %s", err.Error())
	}

	// Try updating current organization address
	statusMessage, err = client.UpdateActualOrgAddress(ctx, address)
	if err != nil {
		t.Fatalf("failed to update the address: %v (%s)", statusMessage, err.Error())
	}

	retrievedOrg, err = client.GetActualOrg(ctx)
	if err != nil {
		t.Fatalf("failed to retrieved org: %s", err.Error())
	}

	// Check if retrieved address values are equal to expected ones
	if !reflect.DeepEqual(retrievedOrg.Address, address) {
		t.Fatalf("got wrong address: got %v, expected %v", retrievedOrg.Address, address)
	}
}
