package sdk

/*
   Copyright 2016-2020 The Grafana SDK authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

	   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type (
	AlertEvalMatch struct {
		Metric string `json:"metric"`
		Tags   struct {
			ServiceName string `json:"service_name"`
		} `json:"tags"`
		Value int `json:"value"`
	}
	AlertEvalData struct {
		EvalMatches []AlertEvalMatch `json:"evalMatches"`
	}
	FoundAlert struct {
		ID             int           `json:"id"`
		DashboardID    int           `json:"dashboardId"`
		DashboardUID   string        `json:"dashboardUid"`
		DashboardSlug  string        `json:"dashboardSlug"`
		PanelID        int           `json:"panelId"`
		Name           string        `json:"name"`
		State          string        `json:"state"`
		NewStateDate   time.Time     `json:"newStateDate"`
		EvalDate       time.Time     `json:"evalDate"`
		EvalData       AlertEvalData `json:"evalData"`
		ExecutionError string        `json:"executionError"`
		URL            string        `json:"url"`
	}
)

// SearchAlerts searches for alerts with query params specified.SearchAlerts
//
// Refleects GET /api/alerts API call.
func (r *Client) SearchAlerts(ctx context.Context, params ...SearchParam) ([]FoundAlert, error) {
	var (
		raw    []byte
		alerts []FoundAlert
		code   int
		err    error
	)
	u := url.URL{}
	q := u.Query()
	for _, p := range params {
		p(&q)
	}
	if raw, code, err = r.get(ctx, "api/alerts", q); err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	err = json.Unmarshal(raw, &alerts)
	return alerts, err
}

type AlertStatusMessage struct {
	AlertID *int    `json:"alertId"`
	State   *string `json:"state"`
	Message *string `json:"message"`
}

// SetAlertPauseByID updates whether or not an alert is paused.SetAlertPauseByID
//
// Reflects POST /api/alerts/:id/pause
func (r *Client) SetAlertPauseByID(ctx context.Context, id int, pause bool) (AlertStatusMessage, error) {
	var (
		setPauseStatus struct {
			Paused bool `json:"paused"`
		}
		resp AlertStatusMessage
		raw  []byte
		code int
		err  error
	)
	setPauseStatus.Paused = pause
	if raw, err = json.Marshal(setPauseStatus); err != nil {
		return AlertStatusMessage{}, err
	}
	if raw, code, err = r.post(ctx, fmt.Sprintf("api/alerts/%d/pause", id), nil, raw); err != nil {
		return AlertStatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return AlertStatusMessage{}, err
	}
	if code != 200 {
		return resp, fmt.Errorf("HTTP error %d: returns %s", code, *resp.Message)
	}
	return resp, nil
}
