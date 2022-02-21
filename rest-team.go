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
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// SearchTeamsWithPaging search teams with paging.
// query optional.  query value is contained in one of the name, login or email fields. Query values with spaces need to be url encoded e.g. query=Jane%20Doe
// perpage optional. default 1000
// page optional. default 1
// http://docs.grafana.org/http_api/team/#search-teams
// http://docs.grafana.org/http_api/team/#search-teams-with-paging
//
// Reflects GET /api/teams/search API call.
func (r *Client) SearchTeamsWithPaging(ctx context.Context, query *string, perpage *int, page *int) (PageTeams, error) {
	var (
		raw       []byte
		pageTeams PageTeams
		code      int
		err       error
	)

	var params url.Values = nil
	if perpage != nil && page != nil {
		if params == nil {
			params = url.Values{}
		}
		params["perpage"] = []string{fmt.Sprint(*perpage)}
		params["page"] = []string{fmt.Sprint(*page)}
	}

	if query != nil {
		if params == nil {
			params = url.Values{}
		}
		params["query"] = []string{*query}
	}

	if raw, code, err = r.get(ctx, "api/teams/search", params); err != nil {
		return pageTeams, err
	}
	if code != 200 {
		return pageTeams, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&pageTeams); err != nil {
		return pageTeams, fmt.Errorf("unmarshal teams: %s\n%s", err, raw)
	}
	return pageTeams, err
}

func (r *Client) GetTeamByName(ctx context.Context, name string) (PageTeams, error) {
	var (
		raw       []byte
		pageTeams PageTeams
		code      int
		err       error
	)

	var params = url.Values{}
	params["name"] = []string{name}

	if raw, code, err = r.get(ctx, "api/teams/search", params); err != nil {
		return pageTeams, err
	}
	if code != 200 {
		return pageTeams, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&pageTeams); err != nil {
		return pageTeams, fmt.Errorf("unmarshal teams: %s\n%s", err, raw)
	}
	return pageTeams, err
}

// GetTeam gets an team by ID.
// Reflects GET /api/teams/:id API call.
func (r *Client) GetTeam(ctx context.Context, id uint) (Team, error) {
	var (
		raw  []byte
		team Team
		code int
		err  error
	)
	if raw, code, err = r.get(ctx, fmt.Sprintf("api/teams/%d", id), nil); err != nil {
		return team, err
	}
	if code != 200 {
		return team, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&team); err != nil {
		return team, fmt.Errorf("unmarshal team: %s\n%s", err, raw)
	}
	return team, err
}

// CreateTeam creates a new team.
// Reflects POST /api/teams API call.
func (r *Client) CreateTeam(ctx context.Context, t Team) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(t); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post(ctx, "api/teams", nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// UpdateTeam updates a team.
// Reflects PUT /api/teams/:id API call.
func (r *Client) UpdateTeam(ctx context.Context, id uint, t Team) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(t); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.put(ctx, fmt.Sprintf("api/teams/%d", id), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// DeleteTeam deletes a team.
// Reflects DELETE /api/teams/:id API call.
func (r *Client) DeleteTeam(ctx context.Context, id uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, _, err = r.delete(ctx, fmt.Sprintf("api/teams/%d", id)); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// GetTeamMembers gets the members of a team by id.
// Reflects GET /api/teams/:teamId/members API call.
func (r *Client) GetTeamMembers(ctx context.Context, teamId uint) ([]TeamMember, error) {
	var (
		raw         []byte
		teamMembers []TeamMember
		code        int
		err         error
	)
	if raw, code, err = r.get(ctx, fmt.Sprintf("api/teams/%d/members", teamId), nil); err != nil {
		return teamMembers, err
	}
	if code != 200 {
		return teamMembers, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&teamMembers); err != nil {
		return teamMembers, fmt.Errorf("unmarshal team: %s\n%s", err, raw)
	}
	return teamMembers, err
}

// AddTeamMember adds a member to a team.
// Reflects POST /api/teams/:teamId/members API call.
func (r *Client) AddTeamMember(ctx context.Context, teamId uint, userId uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(struct {
		UserId uint `json:"userId"`
	}{
		UserId: userId,
	}); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post(ctx, fmt.Sprintf("api/teams/%d/members", teamId), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// DeleteTeamMember removes a ream member from a team by id.
// Reflects DELETE /api/teams/:teamId/:userId API call.
func (r *Client) DeleteTeamMember(ctx context.Context, teamId uint, userId uint) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, _, err = r.delete(ctx, fmt.Sprintf("api/teams/%d/members/%d", teamId, userId)); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// GetTeamPreferences gets the preferences for a team by id.
// Reflects GET /api/teams/:teamId/preferences API call.
func (r *Client) GetTeamPreferences(ctx context.Context, teamId uint) (TeamPreferences, error) {
	var (
		raw             []byte
		teamPreferences TeamPreferences
		code            int
		err             error
	)
	if raw, code, err = r.get(ctx, fmt.Sprintf("api/teams/%d/preferences", teamId), nil); err != nil {
		return teamPreferences, err
	}
	if code != 200 {
		return teamPreferences, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&teamPreferences); err != nil {
		return teamPreferences, fmt.Errorf("unmarshal team: %s\n%s", err, raw)
	}
	return teamPreferences, err
}

// UpdateTeamPreferences updates the preferences for a team by id.
// Reflects PUT /api/teams/:teamId/preferences API call.
func (r *Client) UpdateTeamPreferences(ctx context.Context, teamId uint, tp TeamPreferences) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(tp); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.put(ctx, fmt.Sprintf("api/teams/%d/preferences", teamId), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}
