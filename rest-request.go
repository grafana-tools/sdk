package sdk

/*
   Copyright 2016-2017 Alexander I.Grafov <grafov@gmail.com>

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

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

// DefaultHTTPClient initialized Grafana with appropriate conditions.
// It allows you globally redefine HTTP client.
var DefaultHTTPClient = http.DefaultClient

// Client uses Grafana REST API for interacting with Grafana server.
type Client struct {
	baseURL string
	key     string
	client  *http.Client
}

// StatusMessage reflects status message as it returned by Grafana REST API.
type StatusMessage struct {
	ID      *uint   `json:"id"`
	Message *string `json:"message"`
	Slug    *string `json:"slug"`
	Version *int    `json:"version"`
	Status  *string `json:"resp"`
}

// NewClient initializes client for interacting with an instance of Grafana server.
func NewClient(apiURL, apiKey string, client *http.Client) *Client {
	return &Client{baseURL: apiURL, key: fmt.Sprintf("Bearer %s", apiKey), client: client}
}

func (r *Client) get(query string, params url.Values) ([]byte, int, error) {
	return r.doRequest("GET", query, params, nil)
}

func (r *Client) patch(query string, params url.Values, body []byte) ([]byte, int, error) {
	return r.doRequest("PATCH", query, params, bytes.NewBuffer(body))
}

func (r *Client) put(query string, params url.Values, body []byte) ([]byte, int, error) {
	return r.doRequest("PUT", query, params, bytes.NewBuffer(body))
}

func (r *Client) post(query string, params url.Values, body []byte) ([]byte, int, error) {
	return r.doRequest("POST", query, params, bytes.NewBuffer(body))
}

func (r *Client) delete(query string) ([]byte, int, error) {
	return r.doRequest("DELETE", query, nil, nil)
}

func (r *Client) doRequest(method, query string, params url.Values, buf io.Reader) ([]byte, int, error) {
	u, _ := url.Parse(r.baseURL)
	u.Path = path.Join(u.Path, query)
	if params != nil {
		u.RawQuery = params.Encode()
	}
	req, err := http.NewRequest(method, u.String(), buf)
	req.Header.Set("Authorization", r.key)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "autograf")
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return data, resp.StatusCode, err
}
