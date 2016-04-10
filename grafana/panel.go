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

// Each panel may be one of these types.
const (
	CustomType panelType = iota
	DashlistType
	GraphType
	TableType
	TextType
	SinglestatType
)

type (
	// Panel represents panels of different types defined in Grafana.
	Panel struct {
		OfType panelType
		*GraphPanel
		*TablePanel
		*TextPanel
		*SinglestatPanel
		*DashlistPanel
		*CustomPanel
	}
	panelType   int8
	commonPanel struct {
		ID         int     `json:"id"`
		Title      string  `json:"title"`
		Type       string  `json:"type"`
		Error      bool    `json:"error"`
		IsNew      bool    `json:"isNew"`
		Span       float32 `json:"span"`
		Links      []link  `json:"links,omitempty"`
		Datasource *string `json:"datasource,omitempty"`
		Renderer   string  `json:"renderer"`
	}
	GraphPanel struct {
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
		Span            float32       `json:"span"`
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
	TablePanel struct {
		commonPanel
		Columns   []column `json:"columns"`
		Transform string   `json:"transform"`
		Targets   []Target `json:"targets,omitempty"`
	}
	TextPanel struct {
		commonPanel
		Content    string `json:"content"`
		Mode       string `json:"mode"`
		PageSize   uint   `json:"pageSize"`
		Scroll     bool   `json:"scroll"`
		ShowHeader bool   `json:"showHeader"`
		Sort       struct {
			Col  int  `json:"col"`
			Desc bool `json:"desc"`
		} `json:"sort"`
		Styles []columnStyle `json:"styles"`
	}
	SinglestatPanel struct {
		commonPanel
		ValueFontSize string     `json:"valueFontSize"`
		ValueMaps     []valueMap `json:"valueMaps"`
		ValueName     string     `json:"valueName"`
		Targets       []Target   `json:"targets,omitempty"`
	}
	DashlistPanel struct {
		commonPanel
		Mode  string   `json:"mode"`
		Limit uint     `json:"limit"`
		Query string   `json:"query"`
		Tags  []string `json:"tags"`
	}
	CustomPanel map[string]interface{}
)

// for a graph panel
type grid struct {
	LeftLogBase     *int     `json:"leftLogBase"`
	LeftMax         *int     `json:"leftMax"`
	LeftMin         *int     `json:"leftMin"`
	RightLogBase    *int     `json:"rightLogBase"`
	RightMax        *int     `json:"rightMax"`
	RightMin        *int     `json:"rightMin"`
	Threshold1      *float64 `json:"threshold1"`
	Threshold1Color string   `json:"threshold1Color"`
	Threshold2      *float64 `json:"threshold2"`
	Threshold2Color string   `json:"threshold2Color"`
	ThresholdLine   bool     `json:"thresholdLine"`
}

// for a table
type (
	column struct {
		TextType string `json:"text"`
		Value    string `json:"value"`
	}
	columnStyle struct {
		DateFormat string `json:"dateFormat"`
		Pattern    string `json:"pattern"`
		Type       string `json:"type"`
	}
)

// for a singlestat
type valueMap struct {
	Op       string `json:"op"`
	TextType string `json:"text"`
	Value    string `json:"value"`
}

// for an any panel
type Target struct {
	RefID          string `json:"refId"`
	Datasource     string `json:"datasource"`
	Expr           string `json:"expr"`
	IntervalFactor int    `json:"intervalFactor"`
	Step           int    `json:"step"`
	LegendFormat   string `json:"legendFormat"`
}

func (p *Panel) UnmarshalJSON(b []byte) (err error) {
	var probe commonPanel
	if err = json.Unmarshal(b, &probe); err == nil {
		switch probe.Type {
		case "graph":
			var graph GraphPanel
			p.OfType = GraphType
			if err = json.Unmarshal(b, &graph); err == nil {
				p.GraphPanel = &graph
			}
		case "table":
			var table TablePanel
			p.OfType = TableType
			if err = json.Unmarshal(b, &table); err == nil {
				p.TablePanel = &table
			}
		case "text":
			var text TextPanel
			p.OfType = TextType
			if err = json.Unmarshal(b, &text); err == nil {
				p.TextPanel = &text
			}
		case "singlestat":
			var singlestat SinglestatPanel
			p.OfType = SinglestatType
			if err = json.Unmarshal(b, &singlestat); err == nil {
				p.SinglestatPanel = &singlestat
			}
		case "dashlist":
			var dashlist DashlistPanel
			p.OfType = DashlistType
			if err = json.Unmarshal(b, &dashlist); err == nil {
				p.DashlistPanel = &dashlist
			}
		default:
			var custom = make(CustomPanel)
			p.OfType = CustomType
			if err = json.Unmarshal(b, &custom); err == nil {
				p.CustomPanel = &custom
			}
		}
	}
	return
}

func (p *Panel) MarshalJSON() ([]byte, error) {
	if p.GraphPanel != nil {
		return json.Marshal(*p.GraphPanel)
	}
	if p.TablePanel != nil {
		return json.Marshal(*p.TablePanel)
	}
	if p.TextPanel != nil {
		return json.Marshal(*p.TextPanel)
	}
	if p.SinglestatPanel != nil {
		return json.Marshal(*p.SinglestatPanel)
	}
	if p.DashlistPanel != nil {
		return json.Marshal(*p.DashlistPanel)
	}
	return json.Marshal(*p.CustomPanel)
}

// ResetTargets delete all targets defined for a panel.
func (p *Panel) ResetTargets() {
	switch p.OfType {
	case GraphType:
		p.GraphPanel.Targets = nil
	case SinglestatType:
		p.SinglestatPanel.Targets = nil
	case TableType:
		p.TablePanel.Targets = nil
	}
}

// AddTarget adds a new target as defined in the argument
// but with refId letter incremented. Value of refID from
// the argument will be used only if no target with such
// value already exists.
func (p *Panel) AddTarget(t *Target) {
	switch p.OfType {
	case GraphType:
		p.GraphPanel.Targets = append(p.GraphPanel.Targets, *t)
	case SinglestatType:
		p.SinglestatPanel.Targets = append(p.SinglestatPanel.Targets, *t)
	case TableType:
		p.TablePanel.Targets = append(p.TablePanel.Targets, *t)
	}
	// TODO check for existing refID
}

// SetTarget updates a target if target with such refId exists
// or creates a new one.
func (p *Panel) SetTarget(t *Target) {
	setTarget := func(t *Target, targets []Target) {
		for i := range targets {
			if t.RefID == targets[i].RefID {
				targets[i] = *t
				return
			}
		}
		targets = append(targets, *t)
	}
	switch p.OfType {
	case GraphType:
		setTarget(t, p.GraphPanel.Targets)
	case SinglestatType:
		setTarget(t, p.SinglestatPanel.Targets)
	case TableType:
		setTarget(t, p.TablePanel.Targets)
	}
}

// RepeatTargetsForDatasources repeats all existing targets for a panel
// for all provided in the argument datasources. Existing datasources of
// targets are ignored.
func (p *Panel) RepeatTargetsForDatasources(ds []Datasource) {
	// XXX
}

// Targets is iterate over all panel targets. It just returns nil if
// no targets defined for panel of concrete type.
func (p *Panel) Targets() []Target {
	switch p.OfType {
	case GraphType:
		return p.GraphPanel.Targets
	case SinglestatType:
		return p.SinglestatPanel.Targets
	case TableType:
		return p.TablePanel.Targets
	default:
		return nil
	}
}
