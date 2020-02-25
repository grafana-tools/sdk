package sdk_test

import (
	"testing"
)

func TestClient_CheckHealth(t *testing.T) {
	shouldSkip(t)
	client := getClient()

	err := client.CheckHealth()
	if err != nil {
		t.Fatal(err)
	}
}
