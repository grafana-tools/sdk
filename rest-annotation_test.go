package sdk_test

import (
	"net/url"
	"testing"

	"github.com/grafana-tools/sdk"
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
		t.Errorf("expected limit to be %s, but was %s", "3", l)
	}

	tags := v["tags"]
	if len(tags) != 2 {
		t.Errorf("expected length of tags to be %d, but was %d", 2, len(tags))
	}
	if tags[1] != "bar" {
		t.Errorf("expected last tag to be %s, but was %s", "bar", tags[1])
	}

	tp := v.Get("type")
	if tp != "annotation" {
		t.Errorf("expected type to be %s, but was %s", "annotation", tp)
	}

	id := v.Get("dashboardId")
	if id != "1" {
		t.Errorf("expected dashboard to be %s, but was %s", "1", id)
	}
}
