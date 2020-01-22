package sdk_test

import (
	"github.com/grafana-tools/sdk"
	"testing"
)

func Test_Datasource_CRUD(t *testing.T) {
	shouldSkip(t)

	client := getClient()

	datasources, err := client.GetAllDatasources()
	if err != nil {
		t.Error(err)
	}
	if len(datasources) != 0 {
		t.Errorf("expected to get zero datasources, got %#v", datasources)
	}

	db := "grafanasdk"
	ds := sdk.Datasource{
		Name:      "elastic_datasource",
		Type:      "elasticsearch",
		IsDefault: false,
		Database:  &db,
		URL:       "http://1.2.3.4",
		JSONData: map[string]string{
			"esVersion":                  "5",
			"timeField":                  "@timestamp",
			"maxConcurrentShardRequests": "256",
		},
	}

	status, err := client.CreateDatasource(ds)
	if err != nil {
		t.Error(err)
	}

	dsRetrieved, err := client.GetDatasource(uint(*status.ID))
	if err != nil {
		t.Error(err)
	}

	if dsRetrieved.Name != ds.Name {
		t.Errorf("got wrong name: expected %s, was %s", dsRetrieved.Name, ds.Name)
	}

	ds.Name = "elasticsdksource"
	ds.ID = dsRetrieved.ID
	status, err = client.UpdateDatasource(ds)
	if err != nil {
		t.Error(err)
	}

	status, err = client.DeleteDatasourceByName("elasticsdksource")
	if err != nil {
		t.Error(err)
	}

	_, err = client.GetDatasource(uint(*status.ID))
	if err == nil {
		t.Errorf("expected the datasource to be deleted")
	}
}
