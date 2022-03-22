package sdk_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/grafana-tools/sdk"
	"github.com/stretchr/testify/assert"
)

func Test_Datasource_CRUD(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)
	ctx := context.Background()

	datasources, err := client.GetAllDatasources(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(datasources) != 0 {
		t.Fatalf("expected to get zero datasources, got %#v", datasources)
	}

	db := "grafanasdk"
	ds := sdk.Datasource{
		Name:      "elastic_datasource",
		Type:      "elasticsearch",
		IsDefault: false,
		Database:  &db,
		URL:       "http://1.2.3.4",
		Access:    "direct",
		JSONData: map[string]string{
			"esVersion":                  "5",
			"timeField":                  "@timestamp",
			"maxConcurrentShardRequests": "256",
		},
	}

	status, err := client.CreateDatasource(ctx, ds)
	if err != nil {
		t.Fatal(err)
	}

	dsRetrieved, err := client.GetDatasource(ctx, uint(*status.ID))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, dsRetrieved.Name, ds.Name, fmt.Sprintf("got wrong name: expected %s, was %s", dsRetrieved.Name, ds.Name))
	//assert.NotNilf(t)

	ds.Name = "elasticsdksource"
	ds.ID = dsRetrieved.ID
	ds.UID = "ZrY0GDS7z"
	readOnly := false
	ds.ReadOnly = &readOnly
	ds.TypeLogoURL = "www.duckduckgo.com"
	status, err = client.UpdateDatasource(ctx, ds)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.DeleteDatasourceByName(ctx, "elasticsdksource")
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.GetDatasource(ctx, uint(*status.ID))
	if err == nil {
		t.Fatalf("expected the datasource to be deleted")
	}
}
