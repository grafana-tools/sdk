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
	"fmt"
)

// Checks health of the grafana base url
// Reflects GET BaseURL API call.
func (r *Client) CheckHealth() error {
	var (
		raw  []byte
		code int
		err  error
	)
	if raw, code, err = r.get("", nil); err != nil {
		return err
	}
	if code != 200 {
		return fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	return nil
}
