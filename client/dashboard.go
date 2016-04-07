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

	"github.com/grafov/autograf/grafana"
)

func (r *Instance) SetDashboard(b *grafana.Board) {

}

func (r *Instance) GetDashboard(slug string) (grafana.Board, error) {
	var (
		raw   []byte
		board grafana.Board
		err   error
	)
	if raw, err = r.get(fmt.Sprintf("api/dashboards/db/%s", slug), nil); err != nil {
		return grafana.Board{}, err
	}
	err = json.Unmarshal(raw, &board)
	return board, err
}

// SearchDashboards search dashboards by query substring. Il allows restrict the result set with
// only starred dashboards and only for tags (logical OR applied to multiple tags).
func (r *Instance) SearchDashboards(query string, starred bool, tags ...string) ([]grafana.Board, error) {
	var (
		raw    []byte
		boards []grafana.Board
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
