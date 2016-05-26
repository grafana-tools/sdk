package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

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

type User struct {
	Login          string `json:"login"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Theme          string `json:"theme"`
	OrgID          int    `json:"orgId"`
	IsGrafanaAdmin bool   `json:"isGrafanaAdmin"`
}

// GetUser loads a dashboard from Grafana instance along with metadata for a dashboard.
func (r *Instance) GetUser() (User, error) {
	var (
		raw  []byte
		user User
		code int
		err  error
	)
	if raw, code, err = r.get("api/user", nil); err != nil {
		return User{}, err
	}
	if code != 200 {
		return User{}, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&user); err != nil {
		return User{}, fmt.Errorf("unmarshal user: %s\n%s", err, raw)
	}
	return user, err
}
