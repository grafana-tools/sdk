package sdk

import (
	"net/http"
	"testing"
)

var client = NewClient("http://localhost:3000", "admin:admin", DefaultClient)

func Test_Alertnotification_CRUD(t *testing.T) {
	alertnotifications, err := client.GetAllAlertNotifications()
	if err != nil {
		t.Error(err)
	}
	if len(alertnotifications) != 0 {
		t.Errorf("expected to get zero alertnotifications, got %#v", alertnotifications)
	}
}
