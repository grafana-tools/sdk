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

import (
	"bytes"
	"encoding/json"
	"errors"
	"strconv"
)

// Each panel may be one of these types.
const (
	CustomType panelType = iota
	DashlistType
	GraphType
	TableType
	TextType
	SinglestatType
)

const MixedSource = "-- Mixed --"

type (
	// Panel represents panels of different types defined in Grafana.
	Panel struct {
		commonPanel
		// Should be initialized only one type of panels.
		// OfType field defines which of types below will be used.
		*GraphPanel
		*TablePanel
		*TextPanel
		*SinglestatPanel
		*DashlistPanel
		*CustomPanel
	}
	commonPanel struct {
		OfType     panelType `json:"-"` // it required for defining type of the panel
		ID         uint      `json:"id"`
		Title      string    `json:"title"`                // general
		Span       float32   `json:"span"`                 // general
		Links      []link    `json:"links,omitempty"`      // general
		Datasource *string   `json:"datasource,omitempty"` // metrics
		Height     *string   `json:"height,omitempty"`     // general
		Renderer   *string   `json:"renderer,omitempty"`   // display styles
		Repeat     *string   `json:"repeat,omitempty"`     // templating options
		//RepeatIteration *int64   `json:"repeatIteration,omitempty"`
		RepeatPanelID *uint `json:"repeatPanelId,omitempty"`
		ScopedVars    map[string]struct {
			Selected bool   `json:"selected"`
			Text     string `json:"text"`
			Value    string `json:"value"`
		} `json:"scopedVars,omitempty"`
		Transparent bool     `json:"transparent"`
		MinSpan     *float32 `json:"minSpan,omitempty"` // templating options
		Type        string   `json:"type"`
		Error       bool     `json:"error"`
		IsNew       bool     `json:"isNew"`
	}
	panelType  int8
	GraphPanel struct {
		AliasColors interface{} `json:"aliasColors"` // XXX
		Bars        bool        `json:"bars"`
		Fill        int         `json:"fill"`
		Grid        grid        `json:"grid"`
		Legend      struct {
			AlignAsTable bool `json:"alignAsTable"`
			Avg          bool `json:"avg"`
			Current      bool `json:"current"`
			HideEmpty    bool `json:"hideEmpty"`
			HideZero     bool `json:"hideZero"`
			Max          bool `json:"max"`
			Min          bool `json:"min"`
			RightSide    bool `json:"rightSide"`
			Show         bool `json:"show"`
			Total        bool `json:"total"`
			Values       bool `json:"values"`
		} `json:"legend,omitempty"`
		LeftYAxisLabel  *string         `json:"leftYAxisLabel,omitempty"`
		RightYAxisLabel *string         `json:"rightYAxisLabel,omitempty"`
		Lines           bool            `json:"lines"`
		Linewidth       uint            `json:"linewidth"`
		NullPointMode   string          `json:"nullPointMode"`
		Percentage      bool            `json:"percentage"`
		Pointradius     int             `json:"pointradius"`
		Points          bool            `json:"points"`
		SeriesOverrides []serieOverride `json:"seriesOverrides"`
		Stack           bool            `json:"stack"`
		SteppedLine     bool            `json:"steppedLine"`
		Targets         []Target        `json:"targets,omitempty"`
		TimeFrom        *string         `json:"timeFrom"`
		TimeShift       *string         `json:"timeShift"`
		Tooltip         struct {
			Shared    bool   `json:"shared"`
			ValueType string `json:"value_type"`
		} `json:"tooltip"`
		XAxis    bool     `json:"x-axis"`
		YAxis    bool     `json:"y-axis"`
		YFormats []string `json:"y_formats"`
	}
	TablePanel struct {
		Columns []column `json:"columns"`
		Sort    *struct {
			Col  uint `json:"col"`
			Desc bool `json:"desc"`
		} `json:"sort,omitempty"`
		Styles    []columnStyle `json:"styles"`
		Transform string        `json:"transform"`
		Targets   []Target      `json:"targets,omitempty"`
	}
	TextPanel struct {
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
		Colors          []string `json:"colors"`
		ColorValue      bool     `json:"colorValue"`
		ColorBackground bool     `json:"colorBackground"`
		Decimals        int      `json:"decimals"`
		Format          string   `json:"format"`
		MaxDataPoints   *int     `json:"maxDataPoints,omitempty"`
		NullPointMode   string   `json:"nullPointMode"`
		Postfix         *string  `json:"postfix,omitempty"`
		Prefix          *string  `json:"prefix,omitempty"`
		PostfixFontSize *string  `json:"postfixFontSize,omitempty"`
		PrefixFontSize  *string  `json:"prefixFontSize,omitempty"`
		SparkLine       struct {
			FillColor *string `json:"fillColor,omitempty"`
			Full      bool    `json:"full,omitempty"`
			LineColor *string `json:"lineColor,omitempty"`
			Show      bool    `json:"show,omitempty"`
		} `json:"sparkline"`
		ValueFontSize string     `json:"valueFontSize"`
		ValueMaps     []valueMap `json:"valueMaps"`
		ValueName     string     `json:"valueName"`
		Targets       []Target   `json:"targets,omitempty"`
		Thresholds    string     `json:"thresholds"`
	}
	DashlistPanel struct {
		Mode  string   `json:"mode"`
		Limit uint     `json:"limit"`
		Query string   `json:"query"`
		Tags  []string `json:"tags"`
	}
	CustomPanel map[string]interface{}
)

// for a graph panel
type (
	grid struct {
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
	serieOverride struct {
		Alias         string      `json:"alias"`
		Bars          *bool       `json:"bars,omitempty"`
		Color         *string     `json:"color,omitempty"`
		Fill          *int        `json:"fill,omitempty"`
		FillBelowTo   *string     `json:"fillBelowTo,omitempty"`
		Legend        *bool       `json:"legend,omitempty"`
		Lines         *bool       `json:"lines,omitempty"`
		Stack         *BoolString `json:"stack,omitempty"`
		Transform     *string     `json:"transform,omitempty"`
		YAxis         *int        `json:"yaxis,omitempty"`
		ZIndex        *int        `json:"zindex,omitempty"`
		NullPointMode *string     `json:"nullPointMode,omitempty"`
	}
)

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

// for a table
type (
	column struct {
		TextType string `json:"text"`
		Value    string `json:"value"`
	}
	columnStyle struct {
		DateFormat *string   `json:"dateFormat,omitempty"`
		Pattern    string    `json:"pattern"`
		Type       string    `json:"type"`
		ColorMode  *string   `json:"colorMode,omitempty"`
		Colors     *[]string `json:"colors,omitempty"`
		Decimals   *uint     `json:"decimals,omitempty"`
		Thresholds *[]string `json:"thresholds,omitempty"`
		Unit       *string   `json:"unit,omitempty"`
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
	Interval       string `json:"interval"`
	Step           int    `json:"step"`
	LegendFormat   string `json:"legendFormat"`
}

// NewDashlist initializes panel with a dashlist panel.
func NewDashlist(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		commonPanel: commonPanel{
			OfType:   DashlistType,
			Title:    title,
			Type:     "dashlist",
			Renderer: &render,
			IsNew:    true},
		DashlistPanel: &DashlistPanel{}}
}

// NewGraph initializes panel with a graph panel.
func NewGraph(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		commonPanel: commonPanel{
			OfType:   GraphType,
			Title:    title,
			Type:     "graph",
			Renderer: &render,
			Span:     12,
			IsNew:    true},
		GraphPanel: &GraphPanel{
			NullPointMode: "connected",
			Pointradius:   5,
			XAxis:         true,
			YAxis:         true,
		}}
}

// NewTable initializes panel with a table panel.
func NewTable(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		commonPanel: commonPanel{
			OfType:   TableType,
			Title:    title,
			Type:     "table",
			Renderer: &render,
			IsNew:    true},
		TablePanel: &TablePanel{}}
}

// NewText initializes panel with a text panel.
func NewText(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		commonPanel: commonPanel{
			OfType:   TextType,
			Title:    title,
			Type:     "text",
			Renderer: &render,
			IsNew:    true},
		TextPanel: &TextPanel{}}
}

// NewSinglestat initializes panel with a singlestat panel.
func NewSinglestat(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		commonPanel: commonPanel{
			OfType:   SinglestatType,
			Title:    title,
			Type:     "singlestat",
			Renderer: &render,
			IsNew:    true},
		SinglestatPanel: &SinglestatPanel{}}
}

// NewCustom initializes panel with a singlestat panel.
func NewCustom(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		commonPanel: commonPanel{
			OfType:   CustomType,
			Title:    title,
			Type:     "singlestat",
			Renderer: &render,
			IsNew:    true},
		CustomPanel: &CustomPanel{}}
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
	setTarget := func(t *Target, targets *[]Target) {
		for i, target := range *targets {
			if t.RefID == target.RefID {
				(*targets)[i] = *t
				return
			}
		}
		(*targets) = append((*targets), *t)
	}
	switch p.OfType {
	case GraphType:
		setTarget(t, &p.GraphPanel.Targets)
	case SinglestatType:
		setTarget(t, &p.SinglestatPanel.Targets)
	case TableType:
		setTarget(t, &p.TablePanel.Targets)
	}
}

// MapDatasources on all existing targets for the panel.
func (p *Panel) RepeatDatasourcesForEachTarget(dsNames ...string) {
	repeatDS := func(dsNames []string, targets *[]Target) {
		var refID = "A"
		originalTargets := *targets
		cleanedTargets := make([]Target, 0, len(originalTargets)*len(dsNames))
		*targets = cleanedTargets
		for _, target := range originalTargets {
			for _, ds := range dsNames {
				newTarget := target
				newTarget.RefID = refID
				newTarget.Datasource = ds
				refID = incRefID(refID)
				*targets = append(*targets, newTarget)
			}
		}
	}
	switch p.OfType {
	case GraphType:
		repeatDS(dsNames, &p.GraphPanel.Targets)
	case SinglestatType:
		repeatDS(dsNames, &p.SinglestatPanel.Targets)
	case TableType:
		repeatDS(dsNames, &p.TablePanel.Targets)
	}
}

// RepeatTargetsForDatasources repeats all existing targets for a panel
// for all provided in the argument datasources. Existing datasources of
// targets are ignored.
func (p *Panel) RepeatTargetsForDatasources(dsNames ...string) {
	repeatTarget := func(dsNames []string, targets *[]Target) {
		var lastRefID string
		lenTargets := len(*targets)
		for i, name := range dsNames {
			if i < lenTargets {
				(*targets)[i].Datasource = name
				lastRefID = (*targets)[i].RefID
			} else {
				newTarget := (*targets)[i%lenTargets]
				lastRefID = incRefID(lastRefID)
				newTarget.RefID = lastRefID
				newTarget.Datasource = name
				*targets = append(*targets, newTarget)
			}
		}
	}
	switch p.OfType {
	case GraphType:
		repeatTarget(dsNames, &p.GraphPanel.Targets)
	case SinglestatType:
		repeatTarget(dsNames, &p.SinglestatPanel.Targets)
	case TableType:
		repeatTarget(dsNames, &p.TablePanel.Targets)
	}
}

// GetTargets is iterate over all panel targets. It just returns nil if
// no targets defined for panel of concrete type.
func (p *Panel) GetTargets() *[]Target {
	switch p.OfType {
	case GraphType:
		return &p.GraphPanel.Targets
	case SinglestatType:
		return &p.SinglestatPanel.Targets
	case TableType:
		return &p.TablePanel.Targets
	default:
		return nil
	}
}

type probePanel struct {
	commonPanel
	//	json.RawMessage
}

func (p *Panel) UnmarshalJSON(b []byte) (err error) {
	var probe probePanel
	if err = json.Unmarshal(b, &probe); err == nil {
		p.commonPanel = probe.commonPanel
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
	switch p.OfType {
	case GraphType:
		var outGraph = struct {
			commonPanel
			GraphPanel
		}{p.commonPanel, *p.GraphPanel}
		return json.Marshal(outGraph)
	case TableType:
		var outTable = struct {
			commonPanel
			TablePanel
		}{p.commonPanel, *p.TablePanel}
		return json.Marshal(outTable)
	case TextType:
		var outText = struct {
			commonPanel
			TextPanel
		}{p.commonPanel, *p.TextPanel}
		return json.Marshal(outText)
	case SinglestatType:
		var outSinglestat = struct {
			commonPanel
			SinglestatPanel
		}{p.commonPanel, *p.SinglestatPanel}
		return json.Marshal(outSinglestat)
	case DashlistType:
		var outDashlist = struct {
			commonPanel
			DashlistPanel
		}{p.commonPanel, *p.DashlistPanel}
		return json.Marshal(outDashlist)
	case CustomType:
		var outCustom = struct {
			commonPanel
			CustomPanel
		}{p.commonPanel, *p.CustomPanel}
		return json.Marshal(outCustom)
	}
	return nil, errors.New("can't marshal unknown panel type")
}

func incRefID(refID string) string {
	firstLetter := refID[0]
	ordinal := int(firstLetter)
	ordinal++
	return string(rune(ordinal))
}
