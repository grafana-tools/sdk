package sdk

import (
	"os"
	"testing"
)

var client = NewClient("http://localhost:3000", "admin:admin", DefaultHTTPClient)

func Test_Alertnotification_CRUD(t *testing.T) {
	if v := os.Getenv("GRAFANA_INTEGRATION"); v != "1" {
		t.Skipf("skipping because GRAFANA_INTEGRATION is %s, not 1", v)
	}

	alertnotifications, err := client.GetAllAlertNotifications()
	if err != nil {
		t.Error(err)
	}
	if len(alertnotifications) != 0 {
		t.Errorf("expected to get zero alertnotifications, got %#v", alertnotifications)
	}

	an := AlertNotification{
		Name:                  "alertnotification",
		Type:                  "prometheus",
		IsDefault:             true,
		DisableResolveMessage: true,
		SendReminder:          true,
		Frequency:             "123",
		UID:                   "foobar",
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
