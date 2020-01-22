package sdk_test

import (
	"github.com/grafana-tools/sdk"
	"os"
	"testing"
)

func Test_Alertnotification_CRUD(t *testing.T) {
	shouldSkip(t)
	client := newClient()

	alertnotifications, err := client.GetAllAlertNotifications()
	if err != nil {
		t.Error(err)
	}
	if len(alertnotifications) != 0 {
		t.Errorf("expected to get zero alertnotifications, got %#v", alertnotifications)
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

	id, err := client.CreateAlertNotification(an)
	if err != nil {
		t.Error(err)
	}

	anRetrieved, err := client.GetAlertNotificationID(uint(id))
	if err != nil {
		t.Error(err)
	}

	if anRetrieved.Name != an.Name {
		t.Errorf("got wrong name: expected %s, was %s", anRetrieved.Name, an.Name)
	}

	an.Name = "alertnotification2"
	err = client.UpdateAlertNotificationUID(an, "foobar")
	if err != nil {
		t.Error(err)
	}

	err = client.DeleteAlertNotificationUID("foobar")
	if err != nil {
		t.Error(err)
	}

	an, err = client.GetAlertNotificationUID("foobar")
	if err == nil {
		t.Errorf("expected the alertnotification to be deleted")
	}
}
