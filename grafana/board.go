package grafana

import (
	"bytes"
	"encoding/json"
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

var (
	boardID uint
)

type (
	// Board represents Grafana dashboard.
	Board struct {
		ID              uint     `json:"id"`
		Title           string   `json:"title"`
		OriginalTitle   string   `json:"originalTitle"`
		Tags            []string `json:"tags"`
		Style           string   `json:"style"`
		Timezone        string   `json:"timezone"`
		Editable        bool     `json:"editable"`
		HideControls    bool     `json:"hideControls" graf:"hide-controls"`
		SharedCrosshair bool     `json:"sharedCrosshair" graf:"shared-crosshair"`
		Rows            []*Row   `json:"rows"`
		Templating      struct {
			List []templateVar `json:"list"`
		} `json:"templating"`
		Annotations struct {
			List []annotation `json:"list"`
		} `json:"annotations"`
		SchemaVersion uint   `json:"schemiaVersion"`
		Version       uint   `json:"version"`
		Links         []link `json:"links"`
		Time          struct {
			From string `json:"from"`
			To   string `json:"to"`
		} `json:"time"`
		Timepicker struct {
			RefreshIntervals []string `json:"refresh_intervals"`
			TimeOptions      []string `json:"time_options"`
		} `json:"timepicker"`
		panelID uint
	}
	templateVar struct {
		Name        string   `json:"name"`
		Type        string   `json:"type"`
		Datasource  *string  `json:"datasource"`
		Refresh     bool     `json:"refresh"`
		Options     []option `json:"options"`
		IncludeAll  bool     `json:"includeAll"`
		AllFormat   string   `json:"allFormat"`
		Multi       bool     `json:"multi"`
		MultiFormat string   `json:"multiFormat"`
		Query       string   `json:"query"`
		Current     current  `json:"current"`
		Label       string   `json:"label"`
	}
	// for templateVar
	option struct {
		Text     string `json:"text"`
		Value    string `json:"value"`
		Selected bool   `json:"selected"`
	}
	// for templateVar
	current struct {
		Text  string `json:"text"`
		Value string `json:"value"`
	}
	annotation struct {
		Name       string  `json:"name"`
		Datasource *string `json:"datasource"`
		ShowLine   bool    `json:"showLine"`
		IconColor  string  `json:"iconColor"`
		LineColor  string  `json:"lineColor"`
		IconSize   uint    `json:"iconSize"`
		Enable     bool    `json:"enable"`
	}
)

type link struct {
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	Dashboard   *string  `json:"dashboard,omitempty"`
	DashURI     *string  `json:"dashUri,omitempty"`
	Params      *string  `json:"params,omitempty"`
	Icon        *string  `json:"icon,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	URL         *string  `json:"url,omitempty"`
	Tooltip     *string  `json:"tooltip,omitempty"`
	IncludeVars bool     `json:"includeVars"`
	KeepTime    *bool    `json:"keepTime,omitempty"`
}

// height of rows maybe passed as number (ex 200) or
// as string (ex "200px") or empty string
type height string

func (h *height) UnmarshalJSON(raw []byte) error {
	if raw == nil || bytes.Compare(raw, []byte(`"null"`)) == 0 {
		return nil
	}
	if raw[0] != '"' {
		tmp := []byte{'"'}
		raw = append(tmp, raw...)
		raw = append(raw, byte('"'))
	}
	var tmp string
	err := json.Unmarshal(raw, &tmp)
	*h = height(tmp)
	return err
}

func NewBoard(title string) *Board {
	boardID++
	return &Board{
		ID:           boardID,
		Title:        title,
		Style:        "dark",
		Timezone:     "browser",
		Editable:     true,
		HideControls: false,
		Rows:         []*Row{NewRow()}}
}
