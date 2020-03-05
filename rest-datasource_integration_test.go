package sdk_test

import (
	"testing"

	"github.com/grafana-tools/sdk"
)

func Test_Datasource_CRUD(t *testing.T) {
	shouldSkip(t)

	client := getClient(t)

	datasources, err := client.GetAllDatasources()
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

	status, err := client.CreateDatasource(ds)
	if err != nil {
		t.Fatal(err)
	}

	dsRetrieved, err := client.GetDatasource(uint(*status.ID))
	if err != nil {
		t.Fatal(err)
	}

	if dsRetrieved.Name != ds.Name {
		t.Fatalf("got wrong name: expected %s, was %s", dsRetrieved.Name, ds.Name)
	}

	ds.Name = "elasticsdksource"
	ds.ID = dsRetrieved.ID
	status, err = client.UpdateDatasource(ds)
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.DeleteDatasourceByName("elasticsdksource")
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.GetDatasource(uint(*status.ID))
	if err == nil {
		t.Fatalf("expected the datasource to be deleted")
	}
}
