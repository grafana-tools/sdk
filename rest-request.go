package sdk

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>

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
	"io/ioutil"
	"net/http"
	"net/url"
)

// DefaultHTTPClient initialized Grafana with appropriate conditions.
// It allows you globally redefine HTTP client.
var DefaultHTTPClient = http.DefaultClient

// Instance of Grafana.
type Instance struct {
	baseURL string
	key     string
	client  *http.Client
}

type StatusMessage struct {
	ID      *uint   `json:"id"`
	Message *string `json:"message"`
	Slug    *string `json:"slug"`
	Version *int    `json:"version"`
	Status  *string `json:"resp"`
}

// New keeps request data.
func New(apiURL, apiKey string, client *http.Client) *Instance {
	return &Instance{baseURL: apiURL, key: fmt.Sprintf("Bearer %s", apiKey), client: client}
}

func (r *Instance) get(query string, params url.Values) ([]byte, int, error) {
	u, _ := url.Parse(r.baseURL)
	u.Path = query
	if params != nil {
		u.RawQuery = params.Encode()
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Set("Authorization", r.key)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "autograf")
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return data, resp.StatusCode, err
}

func (r *Instance) put(query string, params url.Values, body []byte) ([]byte, int, error) {
	u, _ := url.Parse(r.baseURL)
	u.Path = query
	if params != nil {
		u.RawQuery = params.Encode()
	}
	req, err := http.NewRequest("PUT", u.String(), bytes.NewBuffer(body))
	req.Header.Set("Authorization", r.key)
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("User-Agent", "autograf")
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return data, resp.StatusCode, err
}

func (r *Instance) post(query string, params url.Values, body []byte) ([]byte, int, error) {
	u, _ := url.Parse(r.baseURL)
	u.Path = query
	if params != nil {
		u.RawQuery = params.Encode()
	}
	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(body))
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

func (r *Instance) delete(query string) ([]byte, error) {
	u, _ := url.Parse(r.baseURL)
	u.Path = query
	req, err := http.NewRequest("DELETE", u.String(), nil)
	req.Header.Set("Authorization", r.key)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "autograf")
	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
