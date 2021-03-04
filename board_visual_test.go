package sdk_test

import (
	"context"
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
	shouldSkip(t)

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
