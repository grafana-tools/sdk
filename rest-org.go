// +build draft

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
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateOrg creates a new organization
// It reflects POST /api/orgs
func (r *Client) CreateOrg(org Org) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(org); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post("api/orgs", nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// GetActualOrg gets current organization.
// It reflects GET /api/org API call.
func (r *Client) GetActualOrg() (Org, error) {
	var (
		raw  []byte
		org  Org
		code int
		err  error
	)
	if raw, code, err = r.get("api/org", nil); err != nil {
		return org, err
	}
	if code != http.StatusOK {
		return org, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&org); err != nil {
		return org, fmt.Errorf("unmarshal org: %s\n%s", err, raw)
	}
	return org, err
}

// GetOrgById gets organization by organization Id
// It reflects GET /api/orgs/:orgId
func (r *Client) GetOrgById(oid uint) (Org, error) {
	var (
		raw  []byte
		org  Org
		code int
		err  error
	)
	if raw, code, err = r.get(fmt.Sprintf("api/orgs/%d", oid), nil); err != nil {
		return org, err
	}

	if code != http.StatusOK {
		return org, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&org); err != nil {
		return org, fmt.Errorf("unmarshal org: %s\n%s", err, raw)
	}
	return org, err
}

// GetOrgByOrgName gets organization by organization Id
// It reflects GET /api/orgs/:orgId
func (r *Client) GetOrgByOrgName(name string) (Org, error) {
	var (
		raw  []byte
		org  Org
		code int
		err  error
	)
	if raw, code, err = r.get(fmt.Sprintf("api/orgs/name/%s", name), nil); err != nil {
		return org, err
	}

	if code != http.StatusOK {
		return org, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&org); err != nil {
		return org, fmt.Errorf("unmarshal org: %s\n%s", err, raw)
	}
	return org, err
}

// UpdateActualOrg updates current organization.
// It reflects PUT /api/org API call.
func (r *Client) UpdateActualOrg(org Org) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(org); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.put("api/org", nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// GetActualOrgUsers get all users within the actual organisation.
func (r *Client) GetActualOrgUsers() ([]User, error) {
	var (
		raw   []byte
		users []User
		code  int
		err   error
	)
	if raw, code, err = r.get("api/org/users", nil); err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&users); err != nil {
		return nil, fmt.Errorf("unmarshal org: %s\n%s", err, raw)
	}
	return users, err
}

// AddUserToActualOrg creates a new organization
// It reflects POST /api/org/users
func (r *Client) AddUserToActualOrg(userRole UserRole) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(userRole); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post("api/org/users", nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// UpdateUser updates the existing user
// It reflects POST /api/org/users/:userId
func (r *Client) UpdateUser(user User, uid uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(user); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post(fmt.Sprintf("api/org/users/%s", uid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// DeleteActualOrgUser delete user in actual organisation.
// It reflects DELETE /api/org/users/:userId API call.
func (r *Client) DeleteActualOrgUser(uid uint) (StatusMessage, error) {
	var (
		raw   []byte
		reply StatusMessage
		err   error
	)
	if raw, _, err = r.delete(fmt.Sprintf("api/org/%d", uid)); err != nil {
		return StatusMessage{}, err
	}
	err = json.Unmarshal(raw, &reply)
	return reply, err
}

// AddUserToOrg add user to organization with id
// It reflects POST /api/orgs/:orgId/users API call.
func (r *Client) AddUserToOrg(user UserRole, oid uint) (StatusMessage, error) {
	var (
		raw   []byte
		reply StatusMessage
		err   error
	)
	if raw, err = json.Marshal(user); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post(fmt.Sprintf("api/orgs/%d/users", oid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	err = json.Unmarshal(raw, &reply)
	return reply, err
}
