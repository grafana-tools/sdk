package api

/*
 Copyleft 2016 Alexander I.Grafov <grafov@gmail.com>

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.

 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

 ॐ तारे तुत्तारे तुरे स्व
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/grafov/autograf/grafana"
)

// Instance of Grafana.
type Instance struct {
	url string
	key string
}

// NewInstance keeps request data.
func NewInstance(apiURL, apiKey string) *Instance {
	return &Instance{url: apiURL, key: fmt.Sprintf("Bearer %s", apiKey)}
}

func (r *Instance) CreateDatasource(ds grafana.Datasource) {
}

func (r *Instance) UpdateDatasource(ds grafana.Datasource) {
}

func (r *Instance) GetAllDatasources() ([]grafana.Datasource, error) {
	var (
		raw []byte
		dss []grafana.Datasource
		err error
	)
	if raw, err = r.get("api/datasources", nil); err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &dss)
	return dss, err
}

func (r *Instance) SetDashboard(b *grafana.Board) {

}

func (r *Instance) GetDashboard(slug string) (grafana.Board, error) {
	var (
		raw   []byte
		board grafana.Board
		err   error
	)
	if raw, err = r.get(fmt.Sprintf("api/dashboards/db/%s", slug), nil); err != nil {
		return grafana.Board{}, err
	}
	err = json.Unmarshal(raw, &board)
	return board, err
}

// SearchDashboards search dashboards by query substring. Il allows restrict the result set with
// only starred dashboards and only for tags (logical OR applied to multiple tags).
func (r *Instance) SearchDashboards(query string, starred bool, tags ...string) ([]grafana.Board, error) {
	var (
		raw    []byte
		boards []grafana.Board
		err    error
	)
	u := url.URL{}
	q := u.Query()
	if query != "" {
		q.Set("query", query)
	}
	if starred {
		q.Set("starred", "true")
	}
	for _, tag := range tags {
		q.Add("tag", tag)
	}
	if raw, err = r.get("api/search", q); err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &boards)
	return boards, err
}

func (r *Instance) get(query string, params url.Values) ([]byte, error) {
	u, _ := url.Parse(r.url)
	u.Path = query
	if params != nil {
		u.RawQuery = params.Encode()
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Authorization", r.key)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "autograf")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (r *Instance) post(query string, params url.Values, body []byte) ([]byte, error) {
	u, _ := url.Parse(r.url)
	u.Path = query
	if params != nil {
		u.RawQuery = params.Encode()
	}
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
	req.Header.Set("Authorization", r.key)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "autograf")
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
