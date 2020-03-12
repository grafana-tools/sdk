package sdk_test

import (
	"testing"
)

func TestClient_GetHealth(t *testing.T) {
	shouldSkip(t)
	client := getClient(t)

	health, err := client.GetHealth()
	if err != nil {
		t.Fatal(err)
	}
	if health.Database != "ok" {
		t.Fatalf("expected `Database` to be %v, got %v.", "ok", health.Database)
	}
}
