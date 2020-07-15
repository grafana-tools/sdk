package sdk_test

import (
	"context"
	"testing"

	"github.com/galamiram/sdk"
)

func Test_Alertnotification_CRUD(t *testing.T) {
	shouldSkip(t)
	client := getClient(t)
	ctx := context.Background()

	alertnotifications, err := client.GetAllAlertNotifications(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(alertnotifications) != 0 {
		t.Fatalf("expected to get zero alertnotifications, got %#v", alertnotifications)
	}

	an := sdk.AlertNotification{
		Name:                  "team-a-email-notifier",
		Type:                  "email",
		IsDefault:             false,
		DisableResolveMessage: false,
		SendReminder:          false,
		Frequency:             "15m",
		UID:                   "foobar",
		Settings: map[string]string{
			"addresses": "dev@null.com",
		},
	}

	id, err := client.CreateAlertNotification(ctx, an)
	if err != nil {
		t.Fatal(err)
	}

	anRetrieved, err := client.GetAlertNotificationID(ctx, uint(id))
	if err != nil {
		t.Fatal(err)
	}

	if anRetrieved.Name != an.Name {
		t.Fatalf("got wrong name: expected %s, was %s", anRetrieved.Name, an.Name)
	}

	an.Name = "alertnotification2"
	err = client.UpdateAlertNotificationUID(ctx, an, "foobar")
	if err != nil {
		t.Fatal(err)
	}

	err = client.DeleteAlertNotificationUID(ctx, "foobar")
	if err != nil {
		t.Fatal(err)
	}

	an, err = client.GetAlertNotificationUID(ctx, "foobar")
	if err == nil {
		t.Fatalf("expected the alertnotification to be deleted")
	}
}
