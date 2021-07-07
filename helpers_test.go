package sdk_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/grafana-tools/sdk"
)

func getDebugURL(t *testing.T) string {
	t.Helper()

	resp, err := http.Get("http://localhost:9222/json/version")
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatal(err)
	}
	return result["webSocketDebuggerUrl"].(string)
}

func getFullUrl(t *testing.T) (addr, user, pass string) {
	t.Helper()

	addr = "http://localhost:3000"
	user = "admin"
	pass = "admin"

	if a := os.Getenv("GRAFANA_ADDR"); a != "" {
		addr = a
	}
	if u := os.Getenv("GRAFANA_USER"); u != "" {
		user = u
	}
	if p := os.Getenv("GRAFANA_PASS"); p != "" {
		pass = p
	}

	return
}

func getClient(t *testing.T) *sdk.Client {
	t.Helper()
	addr, user, pass := getFullUrl(t)

	cl, _ := sdk.NewClient(addr, fmt.Sprintf("%s:%s", user, pass), sdk.DefaultHTTPClient)
	return cl
}

func shouldSkip(t *testing.T) {
	t.Helper()

	if v := os.Getenv("GRAFANA_INTEGRATION"); v != "1" {
		t.Skipf("skipping because GRAFANA_INTEGRATION is %s, not 1", v)
	}
}
