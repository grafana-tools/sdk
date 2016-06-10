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
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/grafov/autograf/grafana"
)

// BoardProperties keeps metadata of a dashboard.
type BoardProperties struct {
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

// GetDashboard loads a dashboard from Grafana instance along with metadata for a dashboard.
func (r *Instance) GetDashboard(slug string) (grafana.Board, BoardProperties, error) {
	var (
		raw    []byte
		result struct {
			Meta  BoardProperties `json:"meta"`
			Board grafana.Board   `json:"dashboard"`
		}
		code int
		err  error
	)
	slug, _ = setPrefix(slug)
	if raw, code, err = r.get(fmt.Sprintf("api/dashboards/%s", slug), nil); err != nil {
		return grafana.Board{}, BoardProperties{}, err
	}
	if code != 200 {
		return grafana.Board{}, BoardProperties{}, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&result); err != nil {
		return grafana.Board{}, BoardProperties{}, fmt.Errorf("unmarshal board with meta: %s\n%s", err, raw)
	}
	return result.Board, result.Meta, err
}

// GetRawDashboard loads a dashboard JSON from Grafana instance along with metadata for a dashboard.
// Contrary to GetDashboard() it not unpack loaded JSON to grafana.Board structure. Instead it
// returns it as byte slice. It guarantee that data of dashboard returned untouched by conversion
// with grafana.Board so no matter how properly fields from a current version of Grafana mapped to
// our grafana.Board fields. It useful for backuping purposes when you want a dashboard exactly with
// same data as it exported by Grafana.
func (r *Instance) GetRawDashboard(slug string) ([]byte, BoardProperties, error) {
	var (
		raw    []byte
		result struct {
			Meta  BoardProperties `json:"meta"`
			Board json.RawMessage `json:"dashboard"`
		}
		code int
		err  error
	)
	slug, _ = setPrefix(slug)
	if raw, code, err = r.get(fmt.Sprintf("api/dashboards/%s", slug), nil); err != nil {
		return nil, BoardProperties{}, err
	}
	if code != 200 {
		return nil, BoardProperties{}, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&result); err != nil {
		return nil, BoardProperties{}, fmt.Errorf("unmarshal board with meta: %s\n%s", err, raw)
	}
	return []byte(result.Board), result.Meta, err
}

// FoundBoard keeps result of search with metadata of a dashboard.
type FoundBoard struct {
	ID        uint     `json:"id"`
	Title     string   `json:"title"`
	URI       string   `json:"uri"`
	Type      string   `json:"type"`
	Tags      []string `json:"tags"`
	IsStarred bool     `json:"isStarred"`
}

// SearchDashboards search dashboards by substring of their title. It allows restrict the result set with
// only starred dashboards and only for tags (logical OR applied to multiple tags).
func (r *Instance) SearchDashboards(query string, starred bool, tags ...string) ([]FoundBoard, error) {
	var (
		raw    []byte
		boards []FoundBoard
		code   int
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
	if raw, code, err = r.get("api/search", q); err != nil {
		return nil, err
	}
	if code != 200 {
		return nil, fmt.Errorf("HTTP error %d: returns %s", code, raw)
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
		isBoardFromDB bool
		newBoard      struct {
			Dashboard grafana.Board `json:"dashboard"`
			Overwrite bool          `json:"overwrite"`
		}
		raw  []byte
		resp StatusMessage
		code int
		err  error
	)
	if board.Slug, isBoardFromDB = cleanPrefix(board.Slug); !isBoardFromDB {
		return errors.New("only database dashboard (with 'db/' prefix in a slug) can be set")
	}
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
		isBoardFromDB bool
		raw           []byte
		reply         StatusMessage
		err           error
	)
	if slug, isBoardFromDB = setPrefix(slug); !isBoardFromDB {
		return StatusMessage{}, errors.New("only database dashboards (with 'db/' prefix in a slug) can be removed")
	}
	if raw, err = r.delete(fmt.Sprintf("api/dashboards/db/%s", slug)); err != nil {
		return StatusMessage{}, err
	}
	err = json.Unmarshal(raw, &reply)
	return reply, err
}

// implicitely use dashboards from Grafana DB not from a file system
func setPrefix(slug string) (string, bool) {
	if strings.HasPrefix(slug, "db") {
		return slug, true
	}
	if strings.HasPrefix(slug, "file") {
		return slug, false
	}
	return fmt.Sprintf("db/%s", slug), true
}

// assume we use database dashboard by default
func cleanPrefix(slug string) (string, bool) {
	if strings.HasPrefix(slug, "db") {
		return slug[3:], true
	}
	if strings.HasPrefix(slug, "file") {
		return slug[3:], false
	}
	return fmt.Sprintf("%s", slug), true
}
