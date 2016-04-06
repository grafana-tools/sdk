package grafana

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

import "encoding/json"

type (
	// Panel represents panels of different types defined in Grafana.
	// Panels declared in external plugins maybe parsed too but
	// without manipulating them. They just keeped as is for output
	// in a resulting JSON.
	Panel struct {
		*graphPanel
		*tablePanel
		*textPanel
		*singlestatPanel
		*dashlistPanel
		*unknownPanel
	}
	commonPanel struct {
		ID         int     `json:"id"`
		Title      string  `json:"title"`
		Type       string  `json:"type"`
		Error      bool    `json:"error"`
		IsNew      bool    `json:"isNew"`
		Span       uint    `json:"span"`
		Links      []link  `json:"links,omitempty"`
		Datasource *string `json:"datasource,omitempty"`
		Renderer   string  `json:"renderer"`
	}
	graphPanel struct {
		commonPanel
		AliasColors interface{} `json:"aliasColors"` // XXX
		Bars        bool        `json:"bars"`
		Fill        int         `json:"fill"`
		Grid        grid        `json:"grid"`
		Legend      struct {
			Avg     bool `json:"avg"`
			Current bool `json:"current"`
			Max     bool `json:"max"`
			Min     bool `json:"min"`
			Show    bool `json:"show"`
			Total   bool `json:"total"`
			Values  bool `json:"values"`
		} `json:"legend"`
		Lines           bool          `json:"lines"`
		Linewidth       uint          `json:"linewidth"`
		NullPointMode   string        `json:"nullPointMode"`
		Percentage      bool          `json:"percentage"`
		Pointradius     int           `json:"pointradius"`
		Points          bool          `json:"points"`
		seriesOverrides []interface{} // TODO
		Span            int           `json:"span"`
		Stack           bool          `json:"stack"`
		SteppedLine     bool          `json:"steppedLine"`
		Targets         []Target      `json:"targets,omitempty"`
		TimeFrom        *string       `json:"timeFrom"`
		TimeShift       *string       `json:"timeShift"`
		Tooltip         struct {
			Shared    bool   `json:"shared"`
			ValueType string `json:"value_type"`
		} `json:"tooltip"`
		XAxis    bool     `json:"x-axis"`
		YAxis    bool     `json:"y-axis"`
		YFormats []string `json:"y_formats"`
	}
	tablePanel struct {
		commonPanel
		Columns   []string `json:"columns"`
		Transform string   `json:"transform"`
	}
	textPanel struct {
		commonPanel
		Content string `json:"content"`
		Mode    string `json:"mode"`
	}
	singlestatPanel struct {
		commonPanel
		ValueFontSize string     `json:"valueFontSize"`
		ValueMaps     []valueMap `json:"valueMaps"`
		ValueName     string     `json:"valueName"`
	}
	dashlistPanel struct {
		commonPanel
		Mode  string   `json:"mode"`
		Limit uint     `json:"limit"`
		Query string   `json:"query"`
		Tags  []string `json:"tags"`
	}
	unknownPanel map[string]interface{}
)

// for graph panel
type grid struct {
	LeftLogBase     *int   `json:"leftLogBase"`
	LeftMin         *int   `json:"leftMin"`
	RightLogBase    *int   `json:"rightLogBase"`
	RightMax        *int   `json:"rightMax"`
	RightMin        *int   `json:"rightMin"`
	Threshold1      *int   `json:"threshold1"`
	Threshold1Color string `json:"threshold1Color"`
	Threshold2      *int   `json:"threshold2"`
	Threshold2Color string `json:"threshold2Color"`
}

// for singlestat
type valueMap struct {
	Op    string `json:"op"`
	Text  string `json:"text"`
	Value string `json:"value"`
}

// for any panel
type Target struct {
	RefID          string `json:"refId"`
	Datasource     string `json:"datasource"`
	Expr           string `json:"expr"`
	IntervalFactor int    `json:"intervalFactor"`
	Step           int    `json:"step"`
	LegendFormat   string `json:"legendFormat"`
}

func (p *Panel) UnmarshalJSON(b []byte) (err error) {
	var common commonPanel
	if err = json.Unmarshal(b, &common); err == nil {
		switch common.Type {
		case "graph":
			var graph graphPanel
			if err = json.Unmarshal(b, &graph); err == nil {
				p.graphPanel = &graph
			}
		case "table":
			var table tablePanel
			if err = json.Unmarshal(b, &table); err == nil {
				p.tablePanel = &table
			}
		case "text":
			var text textPanel
			if err = json.Unmarshal(b, &text); err == nil {
				p.textPanel = &text
			}
		case "singlestat":
			var singlestat singlestatPanel
			if err = json.Unmarshal(b, &singlestat); err == nil {
				p.singlestatPanel = &singlestat
			}
		case "dashlist":
			var dashlist dashlistPanel
			if err = json.Unmarshal(b, &dashlist); err == nil {
				p.dashlistPanel = &dashlist
			}
		default:
			var unknown = make(unknownPanel)
			if err = json.Unmarshal(b, &unknown); err == nil {
				p.unknownPanel = &unknown
			}
		}
	}
	return
}

func (p *Panel) MarshalJSON() ([]byte, error) {
	if p.graphPanel != nil {
		return json.Marshal(*p.graphPanel)
	}
	if p.tablePanel != nil {
		return json.Marshal(*p.tablePanel)
	}
	if p.textPanel != nil {
		return json.Marshal(*p.textPanel)
	}
	if p.singlestatPanel != nil {
		return json.Marshal(*p.singlestatPanel)
	}
	if p.dashlistPanel != nil {
		return json.Marshal(*p.dashlistPanel)
	}
	return json.Marshal(*p.unknownPanel)
}

// ResetTargets delete all targets defined for a panel.
func (p *Panel) ResetTargets() {
	p.Targets = []Target{}
}

// AddTarget adds a new target as defined in the argument
// but with refId letter incremented. Value of refID from
// the argument will be used only if no target with such
// value already exists.
func (p *Panel) AddTarget(t *Target) {

}

// AddTarget adds a new target as defined in the argument.
// If old target with such refID exists it will be replaced
// with a new one.
func (p *Panel) SetTarget(t *Target) {

}

// RepeatTargetsForDatasources repeats all existing targets for a panel
// for all provided in the argument datasources. Existing datasources of
// targets are ignored.
func (p *Panel) RepeatTargetsForDatasources(ds []Datasource) {
	// XXX
}
