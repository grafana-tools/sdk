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
	if !health.Alive {
		t.Fatalf("expected `health.Alive` to be %v, got %v.", true, health.Alive)
	}
}
