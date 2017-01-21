package sdk

/*
   Copyright 2016 Alexander I.Grafov <grafov@gmail.com>

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
)

// GetActualUser gets an actual user.
func (r *Client) GetActualUser() (User, error) {
	var (
		raw  []byte
		user User
		code int
		err  error
	)
	if raw, code, err = r.get("api/user", nil); err != nil {
		return user, err
	}
	if code != 200 {
		return user, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&user); err != nil {
		return user, fmt.Errorf("unmarshal user: %s\n%s", err, raw)
	}
	return user, err
}

// GetUser gets an user by ID.
func (r *Client) GetUser(id uint) (User, error) {
	var (
		raw  []byte
		user User
		code int
		err  error
	)
	if raw, code, err = r.get(fmt.Sprintf("api/users/%d", id), nil); err != nil {
		return user, err
	}
	if code != 200 {
		return user, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&user); err != nil {
		return user, fmt.Errorf("unmarshal user: %s\n%s", err, raw)
	}
	return user, err
}

// GetAllUsers gets all users.
func (r *Client) GetAllUsers() ([]User, error) {
	var (
		raw   []byte
		users []User
		code  int
		err   error
	)
	if raw, code, err = r.get("api/users", nil); err != nil {
		return users, err
	}
	if code != 200 {
		return users, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&users); err != nil {
		return users, fmt.Errorf("unmarshal users: %s\n%s", err, raw)
	}
	return users, err
}
