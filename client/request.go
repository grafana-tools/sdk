package client

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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

var HTTPClient = http.Client{}

// Instance of Grafana.
type Instance struct {
	baseURL string
	key     string
}

type StatusMessage struct {
	ID      *uint   `json:"id"`
	Message *string `json:"message"`
	Slug    *string `json:"slug"`
	Version *int    `json:"version"`
	Status  *string `json:"resp"`
}

// New keeps request data.
func New(apiURL, apiKey string) *Instance {
	return &Instance{baseURL: apiURL, key: fmt.Sprintf("Bearer %s", apiKey)}
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
	client := &HTTPClient
	resp, err := client.Do(req)
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
	client := &http.Client{Timeout: requestTimeout}
	resp, err := client.Do(req)
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
	client := &http.Client{Timeout: requestTimeout}
	resp, err := client.Do(req)
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
	client := &http.Client{Timeout: requestTimeout}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
