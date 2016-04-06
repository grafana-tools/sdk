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

func (r *Instance) GetAllDatasources() []grafana.Datasource {
	var dss []grafana.Datasource
	data, err := r.get("api/datasources")
	err = json.Unmarshal(data, &dss)
	fmt.Printf("%+v\n", err) // output for debug

	return dss
}

func (r *Instance) SetBoard(b *grafana.Board) {

}

func (r *Instance) GetBoard(slug string) grafana.Board {
	return grafana.Board{}
}

func (r *Instance) get(query string) ([]byte, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", r.url, query), nil)
	fmt.Printf("%s\n", fmt.Sprintf("%s%s", r.url, query)) // output for debug

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

func (r *Instance) post(query string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", r.url, query), bytes.NewBuffer(body))
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
