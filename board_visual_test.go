package sdk_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/chromedp/chromedp"
	"github.com/grafana-tools/sdk"
)

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>
   Copyright 2016-2021 The Grafana SDK authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

	   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.

   ॐ तारे तुत्तारे तुरे स्व
*/

// Smoke tests for Grafana's singlestat panel.
// Adds a new dashboard with example data via the API and checks if something is there.
func TestSinglestatPanel(t *testing.T) {
	/* These are just defaults values Grafana uses that I have tested */
	b := sdk.NewBoard("exampleboard")
	b.Time.From = "now-5m"
	b.Time.To = "now"
	r := b.AddRow("examplerow")
	panel := sdk.NewSinglestat("teststat")
	panel.SinglestatPanel.Colors = []string{"#299c46", "rgba(237, 129, 40, 0.89)", "#d44a3a"}
	panel.SinglestatPanel.NullPointMode = "connected"
	panel.CommonPanel.Renderer = nil
	fc := "rgba(31, 118, 189, 0.18)"
	lc := "rgb(31, 120, 193)"
	panel.SinglestatPanel.SparkLine = sdk.SparkLine{
		FillColor: &fc,
		Full:      false,
		LineColor: &lc,
		Show:      false,
		YMin:      nil,
		YMax:      nil,
	}
	panel.SinglestatPanel.ValueMaps = []sdk.ValueMap{
		{
			Op:       "=",
			TextType: "N/A",
			Value:    "null",
		},
	}
	panel.SinglestatPanel.ValueName = "avg"

	r.Add(panel)
	cl := getClient(t)

	db, err := cl.SetDashboard(context.TODO(), *b, sdk.SetDashboardParams{
		FolderID:  sdk.DefaultFolderId,
		Overwrite: false,
	})
	if err != nil {
		t.Fatalf("failed setting dashboard: %v", err)
	}

	durl := getDebugURL(t)

	t.Logf("Got Chrome's URL: %s", durl)
	actxt, cancelActxt := chromedp.NewRemoteAllocator(context.Background(), durl)
	defer cancelActxt()

	ctx, cancel := chromedp.NewContext(
		actxt,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var res string

	fullAddr := fmt.Sprintf("http://%s%s", "grafana:3000", *db.URL)
	t.Logf("Got Grafana's URL: %s", fullAddr)

	err = chromedp.Run(ctx,
		chromedp.Navigate(fullAddr),
		chromedp.WaitReady(`grafana-app`),
		chromedp.TextContent(`span.singlestat-panel-value`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		t.Fatalf("running chromedp has failed: %v", err)
	}

	if res == "" {
		t.Fatalf("expected single-stat panel to have some value")
	}
}

// Smoke tests for Grafana's singlestat panel.
func TestGaugePanel(t *testing.T) {
	// Creating a dashboard using json model
	dashboardModelStr := "{\n  \"annotations\": {\n    \"list\": [\n      {\n        \"builtIn\": 1,\n        \"datasource\": \"-- Grafana --\",\n        \"enable\": true,\n        \"hide\": true,\n        \"iconColor\": \"rgba(0, 211, 255, 1)\",\n        \"name\": \"Annotations & Alerts\",\n        \"type\": \"dashboard\"\n      }\n    ]\n  },\n  \"editable\": true,\n  \"gnetId\": null,\n  \"graphTooltip\": 0,\n  \"id\": 2,\n  \"links\": [],\n  \"panels\": [\n    {\n      \"datasource\": \"-- Grafana --\",\n      \"fieldConfig\": {\n        \"defaults\": {\n          \"custom\": {},\n          \"mappings\": [],\n          \"thresholds\": {\n            \"mode\": \"absolute\",\n            \"steps\": [\n              {\n                \"color\": \"green\",\n                \"value\": null\n              },\n              {\n                \"color\": \"red\",\n                \"value\": 80\n              }\n            ]\n          }\n        },\n        \"overrides\": []\n      },\n      \"gridPos\": {\n        \"h\": 9,\n        \"w\": 12,\n        \"x\": 0,\n        \"y\": 0\n      },\n      \"id\": 2,\n      \"options\": {\n        \"reduceOptions\": {\n          \"calcs\": [\n            \"mean\"\n          ],\n          \"fields\": \"\",\n          \"values\": false\n        },\n        \"showThresholdLabels\": false,\n        \"showThresholdMarkers\": true\n      },\n      \"pluginVersion\": \"7.3.6\",\n      \"targets\": [\n        {\n          \"queryType\": \"randomWalk\",\n          \"refId\": \"A\"\n        }\n      ],\n      \"timeFrom\": null,\n      \"timeShift\": null,\n      \"title\": \"Panel Title\",\n      \"type\": \"gauge\"\n    }\n  ],\n  \"schemaVersion\": 26,\n  \"style\": \"dark\",\n  \"tags\": [],\n  \"templating\": {\n    \"list\": []\n  },\n  \"time\": {\n    \"from\": \"now-6h\",\n    \"to\": \"now\"\n  },\n  \"timepicker\": {},\n  \"timezone\": \"\",\n  \"title\": \"gaugetest\",\n  \"uid\": \"gauge\",\n  \"version\": 1\n}"
	var board sdk.Board
	if err := json.Unmarshal([]byte(dashboardModelStr), &board); err != nil {
		t.Fatal("unable to unmarshal gauge panel")
	}

	cl := getClient(t)
	cl.DeleteDashboard(context.TODO(),"gaugetest")
	db, err := cl.SetDashboard(context.TODO(), board, sdk.SetDashboardParams{
		FolderID:  sdk.DefaultFolderId,
		Overwrite: false,
	})
	if err != nil {
		t.Fatalf("failed setting dashboard: %v", err)
	}

	durl := getDebugURL(t)

	t.Logf("Got Chrome's URL: %s", durl)
	actxt, cancelActxt := chromedp.NewRemoteAllocator(context.Background(), durl)
	defer cancelActxt()

	ctx, cancel := chromedp.NewContext(
		actxt,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var res string

	fullAddr := fmt.Sprintf("http://%s%s", "grafana:3000", *db.URL)
	t.Logf("Got Grafana's URL: %s", fullAddr)

	err = chromedp.Run(ctx,
		chromedp.Navigate(fullAddr),
		chromedp.WaitReady(`grafana-app`),
		chromedp.TextContent(`span.gauge-panel-value`, &res, chromedp.NodeVisible, chromedp.ByQuery),
	)
	if err != nil {
		t.Fatalf("running chromedp has failed: %v", err)
	}

	if res == "" {
		t.Fatalf("expected gauge panel to have some value")
	}
}