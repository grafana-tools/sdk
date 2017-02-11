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
	"testing"
)

func TestIntString_Unmarshal(t *testing.T) {
	var i IntString
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
	i := NewIntString(100)

	body, err := json.Marshal(i)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(body, []byte(`100`)) {
		t.Error("Marshalled IntString is not valid: expected '100', got:", string(body))
	}
}
