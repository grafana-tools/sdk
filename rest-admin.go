package sdk

import (
	"encoding/json"
	"fmt"
)

// CreateUser creates a new global user
// Only work with Basic Authentication
// It reflects POST /api/admin/users
func (r *Client) CreateUser(user User) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(user); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post("api/admin/users", nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// UpdateUserPermissions updates the permissions of a global user
// Only work with Basic Authentication
// It reflects PUT /api/admin/users/:userId/permissions
func (r *Client) UpdateUserPermissions(permissions UserPermissions, uid uint) (StatusMessage, error) {
	var (
		raw   []byte
		reply StatusMessage
		err   error
	)
	if raw, err = json.Marshal(permissions); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.put(fmt.Sprintf("api/admin/users/%d/permissions", uid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	err = json.Unmarshal(raw, &reply)
	return reply, err
}

func (r *Client) SwitchUserContext(uid uint, oid uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)

	if raw, _, err = r.post(fmt.Sprintf("/api/users/%d/using/%d", uid, oid), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}
