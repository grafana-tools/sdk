package grafana

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
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

type BoolString struct {
	Flag  bool
	Value string
}

func (s *BoolString) UnmarshalJSON(raw []byte) error {
	if raw == nil || bytes.Compare(raw, []byte(`"null"`)) == 0 {
		return nil
	}
	var (
		tmp string
		err error
	)
	if raw[0] != '"' {
		if bytes.Compare(raw, []byte("true")) == 0 {
			s.Flag = true
			return nil
		}
		if bytes.Compare(raw, []byte("false")) == 0 {
			return nil
		}
		return errors.New("bad boolean value provided")
	}
	if err = json.Unmarshal(raw, &tmp); err != nil {
		return err
	}
	s.Value = tmp
	return nil
}

func (s BoolString) MarshalJSON() ([]byte, error) {
	if s.Value != "" {
		var buf bytes.Buffer
		buf.WriteRune('"')
		buf.WriteString(s.Value)
		buf.WriteRune('"')
		return buf.Bytes(), nil
	}
	return strconv.AppendBool([]byte{}, s.Flag), nil
}

type BoolInt struct {
	Flag  bool
	Value *int64
}

func (s *BoolInt) UnmarshalJSON(raw []byte) error {
	if raw == nil || bytes.Compare(raw, []byte(`"null"`)) == 0 {
		return nil
	}
	var (
		tmp int64
		err error
	)
	if tmp, err = strconv.ParseInt(string(raw), 10, 64); err != nil {
		if bytes.Compare(raw, []byte("true")) == 0 {
			s.Flag = true
			return nil
		}
		if bytes.Compare(raw, []byte("false")) == 0 {
			return nil
		}
		return errors.New("bad value provided")
	}
	s.Value = &tmp
	return nil
}

func (s BoolInt) MarshalJSON() ([]byte, error) {
	if s.Value != nil {
		return strconv.AppendInt([]byte{}, *s.Value, 10), nil
	}
	return strconv.AppendBool([]byte{}, s.Flag), nil
}
