package sdk_test

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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/grafana-tools/sdk"
)

func TestIntString_Unmarshal(t *testing.T) {
	var i sdk.IntString
	raw := []byte(`100`)

	err := json.Unmarshal(raw, &i)
	if err != nil {
		t.Error(err)
	}

	if i.Valid != true {
		t.Error("Unmarshalled IntString is not valid")
	}

	if i.Value != 100 {
		t.Errorf("Unmarshalled IntString should be 100, got: %d", i.Value)
	}
}

func TestIntString_Marshal(t *testing.T) {
	i := sdk.NewIntString(100)

	body, err := json.Marshal(i)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(body, []byte(`100`)) {
		t.Error("Marshalled IntString is not valid: expected '100', got:", string(body))
	}
}

func TestStringSliceString_Unmarshal(t *testing.T) {
	tests := map[string]struct {
		raw    string
		exp    sdk.StringSliceString
		expErr bool
	}{
		"Having an nil should unmarshall correctly as invalid.": {
			raw: `"null"`,
			exp: sdk.StringSliceString{
				Valid: false,
			},
		},

		"Having a string should unmarshall correctly.": {
			raw: `"this is a test"`,
			exp: sdk.StringSliceString{
				Value: []string{"this is a test"},
				Valid: true,
			},
		},

		"Having a empty array should unmarshall correctly.": {
			raw: `[]`,
			exp: sdk.StringSliceString{
				Value: []string{},
				Valid: true,
			},
		},

		"Having a single value array should unmarshall correctly.": {
			raw: `["this is a test"]`,
			exp: sdk.StringSliceString{
				Value: []string{"this is a test"},
				Valid: true,
			},
		},

		"Having a multiple value array should unmarshall correctly.": {
			raw: `["this", "is", "a", "test"]`,
			exp: sdk.StringSliceString{
				Value: []string{"this", "is", "a", "test"},
				Valid: true,
			},
		},

		"Having a wrong value should fail.": {
			raw:    `[`,
			expErr: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			var got sdk.StringSliceString

			err := json.Unmarshal([]byte(test.raw), &got)

			if test.expErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Got unexpected error: %s", err)
				}

				if test.exp.Valid != got.Valid {
					t.Errorf("Valid field is not valid, expected %t; got: %t", test.exp.Valid, got.Valid)
				}

				if fmt.Sprintf("%#v", test.exp.Value) != fmt.Sprintf("%#v", got.Value) {
					t.Errorf("Value field is not valid, expected %#v; got: %#v", test.exp.Value, got.Value)
				}

			}
		})
	}
}

func TestStringSliceString_Marshall(t *testing.T) {
	tests := map[string]struct {
		value  *sdk.StringSliceString
		exp    string
		expErr bool
	}{
		"Having a single value should unmarshall correctly.": {
			value: &sdk.StringSliceString{
				Value: []string{"this is a test"},
				Valid: true,
			},
			exp: `["this is a test"]`,
		},

		"Having an invalid value should return null.": {
			value: &sdk.StringSliceString{},
			exp:   `"null"`,
		},

		"Having a empty array should unmarshall correctly.": {
			value: &sdk.StringSliceString{
				Value: []string{},
				Valid: true,
			},
			exp: `[]`,
		},

		"Having a multiple value array should unmarshall correctly.": {
			value: &sdk.StringSliceString{
				Value: []string{"this", "is", "a", "test"},
				Valid: true,
			},
			exp: `["this","is","a","test"]`,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := json.Marshal(test.value)

			if test.expErr {
				if err == nil {
					t.Error("Expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Got unexpected error: %s", err)
				}

				if test.exp != string(got) {
					t.Errorf("Marshaled value is invalid, expected %s; got: %s", test.exp, string(got))
				}
			}
		})
	}
}
