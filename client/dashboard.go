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
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/grafov/autograf/grafana"
)

type BoardMeta struct {
	IsStarred  bool      `json:"isStarred,omitempty"`
	IsHome     bool      `json:"isHome,omitempty"`
	IsSnapshot bool      `json:"isSnapshot,omitempty"`
	Type       string    `json:"type,omitempty"`
	CanSave    bool      `json:"canSave"`
	CanEdit    bool      `json:"canEdit"`
	CanStar    bool      `json:"canStar"`
	Slug       string    `json:"slug"`
	Expires    time.Time `json:"expires"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
	UpdatedBy  string    `json:"updatedBy"`
	CreatedBy  string    `json:"createdBy"`
	Version    int       `json:"version"`
}

type BoardWithMeta struct {
	Meta  BoardMeta     `json:"meta"`
	Board grafana.Board `json:"dashboard"`
}

func (r *Instance) GetDashboard(slug string) (BoardWithMeta, error) {
	var (
		raw   []byte
		board BoardWithMeta
		err   error
	)
	if raw, err = r.get(fmt.Sprintf("api/dashboards/%s", slug), nil); err != nil {
		return BoardWithMeta{}, err
	}
	err = json.Unmarshal(raw, &board)
	return board, err
}

type FoundBoard struct {
	ID        uint     `json:"id"`
	Title     string   `json:"title"`
	URI       string   `json:"uri"`
	Type      string   `json:"type"`
	Tags      []string `json:"tags"`
	IsStarred bool     `json:"isStarred"`
}

// SearchDashboards search dashboards by query substring. Il allows restrict the result set with
// only starred dashboards and only for tags (logical OR applied to multiple tags).
func (r *Instance) SearchDashboards(query string, starred bool, tags ...string) ([]FoundBoard, error) {
	var (
		raw    []byte
		boards []FoundBoard
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

// SetDashboard updates existing dashboard or creates a new one.
// Set dasboard ID to nil to create a new dashboard.
// Set overwrite to true if you want to overwrite existing dashboard with
// newer version or with same dashboard title.
func (r *Instance) SetDashboard(board grafana.Board, overwrite bool) error {
	var (
		newBoard struct {
			Dashboard grafana.Board `json:"dashboard"`
			Overwrite bool          `json:"overwrite"`
		}
		raw  []byte
		resp StatusMessage
		code int
		err  error
	)
	newBoard.Dashboard = board
	newBoard.Overwrite = overwrite
	if !overwrite {
		newBoard.Dashboard.ID = 0
	}
	if raw, err = json.Marshal(newBoard); err != nil {
		return err
	}
	if raw, code, err = r.post("api/dashboards/db", nil, raw); err != nil {
		return err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return err
	}
	switch code {
	case 401:
		return fmt.Errorf("%d %s", code, *resp.Message)
	case 412:
		return fmt.Errorf("%d %s", code, *resp.Message)
	}
	return nil
}

// DeleteDashboard deletes dashboard that selected by slug string.
func (r *Instance) DeleteDashboard(slug string) (StatusMessage, error) {
	var (
		raw   []byte
		reply StatusMessage
		err   error
	)
	if raw, err = r.delete(fmt.Sprintf("api/dashboards/db/%s", slug)); err != nil {
		return StatusMessage{}, err
	}
	err = json.Unmarshal(raw, &reply)
	return reply, err
}
