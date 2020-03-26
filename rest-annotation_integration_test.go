package sdk_test

import (
	"context"
	"testing"
	"time"

	"github.com/grafana-tools/sdk"
)

func TestAnnotations(t *testing.T) {
	shouldSkip(t)
	client := getClient(t)
	ctx := context.Background()

	ar := sdk.CreateAnnotationRequest{
		Text: "test",
		Time: time.Now().UnixNano() / 1000000,
	}
	resp, err := client.CreateAnnotation(ctx, ar)
	if err != nil {
		t.Fatal(err)
	}

	checkResponse(t, resp, "Annotation added")
	id := *resp.ID

	resp, err = client.PatchAnnotation(ctx, id, sdk.PatchAnnotationRequest{Text: "new text"})
	if err != nil {
		t.Fatal(err)
	}
	checkResponse(t, resp, "Annotation patched")

	anns, err := client.GetAnnotations(ctx, sdk.WithLimit(100))
	if err != nil {
		t.Fatal(err)
	}
	var found bool
	for _, a := range anns {
		if a.ID == id {
			found = true
		}
	}
	if !found {
		t.Errorf("annotation not found")
	}

	resp, err = client.DeleteAnnotation(ctx, id)
	if err != nil {
		t.Fatal(err)
	}
	checkResponse(t, resp, "Annotation deleted")
}

func checkResponse(t *testing.T, r sdk.StatusMessage, msg string) {
	if r.Message == nil {
		t.Errorf("expected message, but was nil")
	} else if *r.Message != msg {
		t.Errorf("expected message '%s', but got '%s'", msg, *r.Message)
	}
}
