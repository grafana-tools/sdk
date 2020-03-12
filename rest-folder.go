package sdk

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>
   Copyright 2016-2019 The Grafana SDK authors

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
	"encoding/json"
	"fmt"
)

// GetAllFolders gets all folders.
// Reflects GET /api/folders API call.
func (r *Client) GetAllFolders() ([]Folder, error) {
	var (
		raw  []byte
		fs   []Folder
		code int
		err  error
	)
	if raw, code, err = r.get("api/folders", nil); err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	err = json.Unmarshal(raw, &fs)
	return fs, err
}

// GetFolderByUID get folder by uid.
// Reflects GET /api/folders/:uid API call.
func (r *Client) GetFolderByUID(UID string) (Folder, error) {
	var (
		raw  []byte
		f    Folder
		code int
		err  error
	)
	if raw, code, err = r.get(fmt.Sprintf("api/folders/%s", UID), nil); err != nil {
		return f, err
	}
	if code != 200 {
		return f, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	err = json.Unmarshal(raw, &f)
	return f, err
}

// CreateFolder create folders.
// Reflects POST /api/folders API call.
func (r *Client) CreateFolder(f Folder) (Folder, error) {
	var (
		raw  []byte
		rf   Folder
		code int
		err  error
	)
	rf = Folder{}
	if raw, err = json.Marshal(f); err != nil {
		return rf, err
	}
	if raw, code, err = r.post("api/folders", nil, raw); err != nil {
		return rf, err
	}
	if code != 200 {
		return rf, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	err = json.Unmarshal(raw, &rf)
	return rf, err
}

// UpdateFolderByUID update folder by uid
// Reflects PUT /api/folders/:uid API call.
func (r *Client) UpdateFolderByUID(f Folder) (Folder, error) {
	var (
		raw  []byte
		rf   Folder
		code int
		err  error
	)
	rf = Folder{}
	if raw, err = json.Marshal(f); err != nil {
		return rf, err
	}
	if raw, code, err = r.put(fmt.Sprintf("api/folders/%s", f.UID), nil, raw); err != nil {
		return rf, err
	}
	if code != 200 {
		return f, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	err = json.Unmarshal(raw, &rf)
	return rf, err
}

// DeleteFolderByUID deletes an existing folder by uid.
// Reflects DELETE /api/folders/:uid API call.
func (r *Client) DeleteFolderByUID(UID string) (bool, error) {
	var (
		raw  []byte
		code int
		err  error
	)
	if raw, code, err = r.delete(fmt.Sprintf("api/folders/%s", UID)); err != nil {
		return false, err
	}
	if code != 200 {
		return false, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	return true, err
}

// GetFolderByID get folder by id.
// Reflects GET /api/folders/id/:id API call.
func (r *Client) GetFolderByID(ID int) (Folder, error) {
	var (
		raw  []byte
		f    Folder
		code int
		err  error
	)
	if raw, code, err = r.get(fmt.Sprintf("api/folders/id/%d", ID), nil); err != nil {
		return f, err
	}
	if code != 200 {
		return f, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	err = json.Unmarshal(raw, &f)
	return f, err
}
