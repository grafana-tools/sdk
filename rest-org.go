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

// GetOrgByOrgName gets organization by organization name
// It reflects GET /api/orgs/name/:orgName
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

//UpdateOrg updates the organization identified by oid.
// It reflects PUT /api/orgs/:orgId
func (r *Client) UpdateOrg(org Org, oid uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(org); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.put(fmt.Sprintf("api/orgs/%d", oid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// DeleteOrg deletes the organization identified by the organization id
// Reflects DELETE /api/orgs/:orgId
func (r *Client) DeleteOrg(oid uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, _, err = r.delete(fmt.Sprintf("api/orgs/%d", oid)); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// GetActualOrgUsers get all users within the actual organisation.
func (r *Client) GetActualOrgUsers() ([]OrgUser, error) {
	var (
		raw   []byte
		users []OrgUser
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

// GetOrgUsers gets the users for the organization specified by organization id
// Reflects GET /api/orgs/:orgId/users
func (r *Client) GetOrgUsers(oid uint) ([]OrgUser, error) {
	var (
		raw   []byte
		users []OrgUser
		code  int
		err   error
	)
	if raw, code, err = r.get(fmt.Sprintf("api/orgs/%d/users", oid), nil); err != nil {
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

// AddActualOrgUser creates a new organization
// It reflects POST /api/org/users
func (r *Client) AddActualOrgUser(userRole UserRole) (StatusMessage, error) {
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
func (r *Client) UpdateActualOrgUser(user UserRole, uid uint) (StatusMessage, error) {
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
	if raw, _, err = r.delete(fmt.Sprintf("api/org/users/%d", uid)); err != nil {
		return StatusMessage{}, err
	}
	err = json.Unmarshal(raw, &reply)
	return reply, err
}

// AddUserToOrg add user to organization with id
// It reflects POST /api/orgs/:orgId/users API call.
func (r *Client) AddOrgUser(user UserRole, oid uint) (StatusMessage, error) {
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

// UpdateOrgUser updates the user specified by uid within the organization specified by oid
// It reflects PATCH /api/orgs/:orgId/users/:userId API call.
func (r *Client) UpdateOrgUser(user UserRole, oid, uid uint) (StatusMessage, error) {
	var (
		raw   []byte
		reply StatusMessage
		err   error
	)
	if raw, err = json.Marshal(user); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.patch(fmt.Sprintf("api/orgs/%d/users/%d", oid, uid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	err = json.Unmarshal(raw, &reply)
	return reply, err
}

// DeleteOrgUser deletes the user specified by uid within the organization specified by oid
// It reflects DELETE /api/orgs/:orgId/users/:userId API call.
func (r *Client) DeleteOrgUser(oid, uid uint) (StatusMessage, error) {
	var (
		raw   []byte
		reply StatusMessage
		err   error
	)
	if raw, _, err = r.delete(fmt.Sprintf("api/orgs/%d/users/%d", oid, uid)); err != nil {
		return StatusMessage{}, err
	}
	err = json.Unmarshal(raw, &reply)
	return reply, err
}
