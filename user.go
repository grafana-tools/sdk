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

const (
	RoleAdmin  = "Admin"
	RoleEditor = "Editor"
	RoleViewer = "Viewer"
)

// User represents fields for both global and group users.
type User struct {
	Id             *uint64 `json:"id,omitempty"`
	Login          *string `json:"login,omitempty"`
	Email          *string `json:"email,omitempty"`
	orgID          uint    `json:"orgId,omitempty"`
	Orgs           []uint  `json:"-"`
	Name           string  `json:"name"`
	Theme          string  `json:"theme"`
	IsAdmin        bool    `json:"isAdmin"`
	isGrafanaAdmin *bool   `json:"isGrafanaAdmin,omitempty"` // GET by ID uses that field but others use IsAdmin
}

// NewUser creates a structure for a new user.
// It should be used in API calls for both global and group users.
func NewUser() *User {
	var user User
	// XXX
	return &user
}
