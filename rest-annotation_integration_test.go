package sdk_test

import (
	"github.com/grafana-tools/sdk"
	"log"
	"net/url"
	"testing"
	"time"
)

func TestAnnotationOptions(t *testing.T) {
	v := make(url.Values)
	params := []sdk.GetAnnotationsParams{
		sdk.WithTag("foo"),
		sdk.WithTag("bar"),
		sdk.WithLimit(3),
		sdk.WithAnnotationType(),
		sdk.WithDashboard(1),
	}

	for _, p := range params {
		p(v)
	}

	l := v.Get("limit")
	if l != "3" {
		t.Errorf("exptected limit to be %s, but was %s", "3", l)
	}

	tags := v["tags"]
	if len(tags) != 2 {
		t.Errorf("expedted length of tags to be %d, but was %d", 2, len(tags))
	}
	if tags[1] != "bar" {
		t.Errorf("expedted last tag to be %s, but was %s", "bar", tags[1])
	}

	tp := v.Get("type")
	if tp != "annotation" {
		t.Errorf("exptected type to be %s, but was %s", "annotation", tp)
	}

	id := v.Get("dashboardId")
	if id != "1" {
		t.Errorf("exptected dashboard to be %s, but was %s", "1", id)
	}
}

func TestAnnotations(t *testing.T) {
	shouldSkip(t)
	client := getClient()

	ar := sdk.CreateAnnotationRequest{
		Text: "test",
		Time: time.Now().UnixNano() / 1_000_000,
	}
	resp, err := client.CreateAnnotation(ar)
	if err != nil {
		t.Fatal(err)
	}

	checkResponse(t, resp, "Annotation added")
	id := *resp.ID
	log.Printf("addotation %d added", id)

	resp, err = client.PatchAnnotation(id, sdk.PatchAnnotationRequest{Text: "new text"})
	if err != nil {
		t.Fatal(err)
	}
	checkResponse(t, resp, "Annotation patched")

	anns, err := client.GetAnnotations(sdk.WithLimit(100))
	if err != nil {
		t.Fatal(err)
	}
	var found bool
	for _, a := range anns {
		log.Printf("%+v", a)
		if a.ID == id {
			found = true
		}
	}
	if !found {
		t.Errorf("annotation not found")
	}

	resp, err = client.DeleteAnnotation(id)
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
