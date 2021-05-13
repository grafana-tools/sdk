package sdk_test

import (
	"context"
	"errors"
	"testing"

	"github.com/grafana-tools/sdk"
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
		UID:       "datasource_uid",
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

	if dsRetrieved.Name != ds.Name {
		t.Fatalf("got wrong name: expected %s, was %s", dsRetrieved.Name, ds.Name)
	}

	dsRetrieved, err = client.GetDatasourceByUID(ctx, "datasource_uid")
	if err != nil {
		t.Fatal(err)
	}

	if dsRetrieved.Name != ds.Name {
		t.Fatalf("got wrong name: expected %s, was %s", dsRetrieved.Name, ds.Name)
	}

	dsRetrieved, err = client.GetDatasourceByName(ctx, "elastic_datasource")
	if err != nil {
		t.Fatal(err)
	}

	if dsRetrieved.Name != ds.Name {
		t.Fatalf("got wrong name: expected %s, was %s", dsRetrieved.Name, ds.Name)
	}

	ds.Name = "elasticsdksource"
	ds.ID = dsRetrieved.ID
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

	_, err = client.GetDatasourceByUID(ctx, "datasource_uid")
	if !errors.As(err, &sdk.ErrNotFound{}) {
		t.Fatalf("expected error not found")
	}

}
