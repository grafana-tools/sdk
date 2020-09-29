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
	"bytes"
	"encoding/json"
	"errors"
)

// Each panel may be one of these types.
const (
	CustomType panelType = iota
	DashlistType
	GraphType
	TableType
	TextType
	PluginlistType
	AlertlistType
	SinglestatType
	RowType
)

const MixedSource = "-- Mixed --"

type (
	// Panel represents panels of different types defined in Grafana.
	Panel struct {
		CommonPanel
		// Should be initialized only one type of panels.
		// OfType field defines which of types below will be used.
		*RowPanel
		*CustomPanel
	}
	panelType   int8
	CommonPanel struct {
		// Default fields.
		GridPos struct {
			H *float32 `json:"h,omitempty"`
			W *float32 `json:"w,omitempty"`
			X *float32 `json:"x,omitempty"`
			Y *float32 `json:"y,omitempty"`
		} `json:"gridPos"`
		ID    uint   `json:"id"`
		Title string `json:"title"`
		Type  string `json:"type"`

		// Optional fields.
		Collapsed        *bool          `json:"collapsed,omitempty"`
		Datasource       *string        `json:"datasource,omitempty"`
		Description      string         `json:"description,omitempty"`
		Editable         *bool          `json:"editable,omitempty"`
		Error            *bool          `json:"error,omitempty"`
		Height           *FloatOrString `json:"height,omitempty"`
		HideTimeOverride *bool          `json:"hideTimeOverride,omitempty"`
		IsNew            *bool          `json:"isNew,omitempty"`
		Links            []Link         `json:"links,omitempty"`
		MinSpan          *float32       `json:"minSpan,omitempty"`  // templating options
		OfType           panelType      `json:"-"`                  // it is required for defining type of the panel
		Renderer         *string        `json:"renderer,omitempty"` // display styles
		Repeat           *string        `json:"repeat,omitempty"`   // templating options
		// RepeatIteration *int64   `json:"repeatIteration,omitempty"`
		RepeatPanelID *uint `json:"repeatPanelId,omitempty"`
		RepeatedByRow bool  `json:"repeatedByRow,omitempty"`
		ScopedVars    map[string]struct {
			Selected bool   `json:"selected"`
			Text     string `json:"text"`
			Value    string `json:"value"`
		} `json:"scopedVars,omitempty"`
		Span        float32  `json:"span,omitempty"`
		Transparent bool     `json:"transparent,omitempty"`
		Alert       *Alert   `json:"alert,omitempty"`
		Targets     []Target `json:"targets,omitempty"`
	}
	AlertEvaluator struct {
		Params []float64 `json:"params,omitempty"`
		Type   string    `json:"type,omitempty"`
	}
	AlertOperator struct {
		Type string `json:"type,omitempty"`
	}
	AlertQuery struct {
		Params []string `json:"params,omitempty"`
	}
	AlertReducer struct {
		Params []string `json:"params,omitempty"`
		Type   string   `json:"type,omitempty"`
	}
	AlertCondition struct {
		Evaluator AlertEvaluator `json:"evaluator,omitempty"`
		Operator  AlertOperator  `json:"operator,omitempty"`
		Query     AlertQuery     `json:"query,omitempty"`
		Reducer   AlertReducer   `json:"reducer,omitempty"`
		Type      string         `json:"type,omitempty"`
	}
	Alert struct {
		Conditions          []AlertCondition    `json:"conditions,omitempty"`
		ExecutionErrorState string              `json:"executionErrorState,omitempty"`
		Frequency           string              `json:"frequency,omitempty"`
		Handler             int                 `json:"handler,omitempty"`
		Name                string              `json:"name,omitempty"`
		NoDataState         string              `json:"noDataState,omitempty"`
		Notifications       []AlertNotification `json:"notifications,omitempty"`
		Message             string              `json:"message,omitempty"`
		For                 string              `json:"for,omitempty"`
	}
	Threshold struct {
		// the alert threshold value, we do not omitempty, since 0 is a valid
		// threshold
		Value float32 `json:"value"`
		// critical, warning, ok, custom
		ColorMode string `json:"colorMode,omitempty"`
		// gt or lt
		Op   string `json:"op,omitempty"`
		Fill bool   `json:"fill"`
		Line bool   `json:"line"`
		// hexadecimal color (e.g. #629e51, only when ColorMode is "custom")
		FillColor string `json:"fillColor,omitempty"`
		// hexadecimal color (e.g. #629e51, only when ColorMode is "custom")
		LineColor string `json:"lineColor,omitempty"`
		// left or right
		Yaxis string `json:"yaxis,omitempty"`
	}

	Tooltip struct {
		Shared       bool   `json:"shared"`
		ValueType    string `json:"value_type"`
		MsResolution bool   `json:"msResolution,omitempty"` // was added in Grafana 3.x
		Sort         int    `json:"sort,omitempty"`
	}
	RowPanel struct {
		Panels []*Panel `json:"panels,omitempty"`
	}
	CustomPanel map[string]json.RawMessage
)

// for a graph panel
type (
	Axis struct {
		Decimals int            `json:"decimals,omitempty"`
		Format   string         `json:"format,omitempty"`
		LogBase  int            `json:"logBase,omitempty"`
		Max      *FloatOrString `json:"max,omitempty"`
		Min      *FloatOrString `json:"min,omitempty"`
		Mode     string         `json:"mode,omitempty"`
		Show     bool           `json:"show"`
		Label    string         `json:"label,omitempty"`
	}
	SeriesOverride struct {
		Alias         string      `json:"alias"`
		Bars          *bool       `json:"bars,omitempty"`
		Color         *string     `json:"color,omitempty"`
		Dashes        bool        `json:"dashes,omitempty"`
		Fill          *int        `json:"fill,omitempty"`
		FillBelowTo   *string     `json:"fillBelowTo,omitempty"`
		Legend        *bool       `json:"legend,omitempty"`
		Lines         *bool       `json:"lines,omitempty"`
		Linewidth     *int        `json:"linewidth,omitempty"`
		Stack         *BoolString `json:"stack,omitempty"`
		Transform     *string     `json:"transform,omitempty"`
		YAxis         *int        `json:"yaxis,omitempty"`
		ZIndex        *int        `json:"zindex,omitempty"`
		NullPointMode *string     `json:"nullPointMode,omitempty"`
	}
	Sort struct {
		Col  int  `json:"col"`
		Desc bool `json:"desc"`
	}
	Legend struct {
		AlignAsTable bool   `json:"alignAsTable"`
		Avg          bool   `json:"avg"`
		Current      bool   `json:"current"`
		HideEmpty    bool   `json:"hideEmpty"`
		HideZero     bool   `json:"hideZero"`
		Max          bool   `json:"max"`
		Min          bool   `json:"min"`
		RightSide    bool   `json:"rightSide"`
		Show         bool   `json:"show"`
		SideWidth    *uint  `json:"sideWidth,omitempty"`
		Sort         string `json:"sort,omitempty"`
		SortDesc     bool   `json:"sortDesc,omitempty"`
		Total        bool   `json:"total"`
		Values       bool   `json:"values"`
	}
)

// for a table
type (
	Column struct {
		TextType string `json:"text"`
		Value    string `json:"value"`
	}
	ColumnStyle struct {
		Alias           *string   `json:"alias"`
		ColorMode       *string   `json:"colorMode,omitempty"`
		Colors          *[]string `json:"colors,omitempty"`
		DateFormat      *string   `json:"dateFormat,omitempty"`
		Decimals        *int      `json:"decimals,omitempty"`
		Link            *bool     `json:"link,omitempty"`
		LinkTargetBlank *bool     `json:"linkTargetBlank,omitempty"`
		LinkTooltip     *string   `json:"linkTooltip,omitempty"`
		LinkURL         *string   `json:"linkUrl,omitempty"`
		Pattern         string    `json:"pattern"`
		Thresholds      *[]string `json:"thresholds,omitempty"`
		Type            string    `json:"type"`
		Unit            *string   `json:"unit,omitempty"`
	}
)

// for a singlestat
type (
	ValueMap struct {
		Op       string `json:"op"`
		TextType string `json:"text"`
		Value    string `json:"value"`
	}
	Gauge struct {
		MaxValue         float32 `json:"maxValue"`
		MinValue         float32 `json:"minValue"`
		Show             bool    `json:"show"`
		ThresholdLabels  bool    `json:"thresholdLabels"`
		ThresholdMarkers bool    `json:"thresholdMarkers"`
	}
	SparkLine struct {
		FillColor *string  `json:"fillColor,omitempty"`
		Full      bool     `json:"full,omitempty"`
		LineColor *string  `json:"lineColor,omitempty"`
		Show      bool     `json:"show,omitempty"`
		YMin      *float64 `json:"ymin,omitempty"`
		YMax      *float64 `json:"ymax,omitempty"`
	}
)

// for an any panel
type Target struct {
	RefID      string `json:"refId"`
	Datasource string `json:"datasource,omitempty"`
	Hide       *bool  `json:"hide,omitempty"`

	// For Prometheus
	Expr           string `json:"expr,omitempty"`
	IntervalFactor int    `json:"intervalFactor,omitempty"`
	Interval       string `json:"interval,omitempty"`
	Step           int    `json:"step,omitempty"`
	LegendFormat   string `json:"legendFormat,omitempty"`
	Instant        bool   `json:"instant,omitempty"`
	Format         string `json:"format,omitempty"`

	catchall map[string]json.RawMessage
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Target) UnmarshalJSON(data []byte) error {
	catchall := make(map[string]json.RawMessage)
	err := json.Unmarshal(data, &catchall)
	if err != nil {
		return err
	}
	t.catchall = catchall
	if v, ok := t.catchall["refId"]; ok {
		vt := t.RefID
		err := json.Unmarshal(v, &vt)
		if err != nil {
			return err
		}
		delete(t.catchall, "refId")
		t.RefID = vt
	}
	if v, ok := t.catchall["datasource"]; ok {
		vt := t.Datasource
		err := json.Unmarshal(v, &vt)
		if err != nil {
			return err
		}
		delete(t.catchall, "datasource")
		t.Datasource = vt
	}
	if v, ok := t.catchall["hide"]; ok {
		vt := t.Hide
		err := json.Unmarshal(v, &vt)
		if err != nil {
			return err
		}
		delete(t.catchall, "hide")
		t.Hide = vt
	}
	if v, ok := t.catchall["expr"]; ok {
		vt := t.Expr
		err := json.Unmarshal(v, &vt)
		if err != nil {
			return err
		}
		delete(t.catchall, "expr")
		t.Expr = vt
	}
	if v, ok := t.catchall["intervalFactor"]; ok {
		vt := t.IntervalFactor
		err := json.Unmarshal(v, &vt)
		if err != nil {
			return err
		}
		delete(t.catchall, "intervalFactor")
		t.IntervalFactor = vt
	}
	if v, ok := t.catchall["interval"]; ok {
		vt := t.Interval
		err := json.Unmarshal(v, &vt)
		if err != nil {
			return err
		}
		delete(t.catchall, "interval")
		t.Interval = vt
	}
	if v, ok := t.catchall["step"]; ok {
		vt := t.Step
		err := json.Unmarshal(v, &vt)
		if err != nil {
			return err
		}
		delete(t.catchall, "step")
		t.Step = vt
	}
	if v, ok := t.catchall["legendFormat"]; ok {
		vt := t.LegendFormat
		err := json.Unmarshal(v, &vt)
		if err != nil {
			return err
		}
		delete(t.catchall, "legendFormat")
		t.LegendFormat = vt
	}
	if v, ok := t.catchall["instant"]; ok {
		vt := t.Instant
		err := json.Unmarshal(v, &vt)
		if err != nil {
			return err
		}
		delete(t.catchall, "instant")
		t.Instant = vt
	}
	if v, ok := t.catchall["format"]; ok {
		vt := t.Format
		err := json.Unmarshal(v, &vt)
		if err != nil {
			return err
		}
		delete(t.catchall, "format")
		t.Format = vt
	}
	return nil
}

// MarshalJSON implements json.Marshaler.
func (t *Target) MarshalJSON() ([]byte, error) {
	// Marshal struct without custom fields.
	tmp := struct {
		RefID      string `json:"refId"`
		Datasource string `json:"datasource,omitempty"`
		Hide       *bool  `json:"hide,omitempty"`

		// For Prometheus
		Expr           string `json:"expr,omitempty"`
		IntervalFactor int    `json:"intervalFactor,omitempty"`
		Interval       string `json:"interval,omitempty"`
		Step           int    `json:"step,omitempty"`
		LegendFormat   string `json:"legendFormat,omitempty"`
		Instant        bool   `json:"instant,omitempty"`
		Format         string `json:"format,omitempty"`
	}{
		RefID:          t.RefID,
		Datasource:     t.Datasource,
		Hide:           t.Hide,
		Expr:           t.Expr,
		IntervalFactor: t.IntervalFactor,
		Interval:       t.Interval,
		Step:           t.Step,
		LegendFormat:   t.LegendFormat,
		Instant:        t.Instant,
		Format:         t.Format,
	}
	b, err := json.Marshal(tmp)
	if err != nil {
		return b, err
	}
	// Append custom keys to marshalled Target.
	buf := bytes.NewBuffer(b[:len(b)-1])
	for k, v := range t.catchall {
		buf.WriteString(`,"`)
		buf.WriteString(k)
		buf.WriteString(`":`)
		b, err := json.Marshal(v)
		if err != nil {
			return b, err
		}
		buf.Write(b)
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

// StackdriverAlignOptions defines the list of alignment options shown in
// Grafana during query configuration.
type StackdriverAlignOptions struct {
	Expanded bool                     `json:"expanded"`
	Label    string                   `json:"label"`
	Options  []StackdriverAlignOption `json:"options"`
}

// StackdriverAlignOption defines a single alignment option shown in Grafana
// during query configuration.
type StackdriverAlignOption struct {
	Label       string   `json:"label"`
	MetricKinds []string `json:"metricKinds"`
	Text        string   `json:"text"`
	Value       string   `json:"value"`
	ValueTypes  []string `json:"valueTypes"`
}

type MapType struct {
	Name  *string `json:"name,omitempty"`
	Value *int    `json:"value,omitempty"`
}

type RangeMap struct {
	From *string `json:"from,omitempty"`
	Text *string `json:"text,omitempty"`
	To   *string `json:"to,omitempty"`
}

func newTrue() *bool {
	b := true
	return &b
}

func newFalse() *bool {
	b := false
	return &b
}

// NewCustom initializes panel with a singlestat panel.
func NewCustom(title string) *Panel {
	if title == "" {
		title = "Panel Title"
	}
	render := "flot"
	return &Panel{
		CommonPanel: CommonPanel{
			OfType:   CustomType,
			Title:    title,
			Type:     "singlestat",
			Renderer: &render,
			IsNew:    newTrue()},
		CustomPanel: &CustomPanel{}}
}

// ResetTargets delete all targets defined for a panel.
func (p *Panel) ResetTargets() {
	p.CommonPanel.Targets = nil
}

// AddTarget adds a new target as defined in the argument
// but with refId letter incremented. Value of refID from
// the argument will be used only if no target with such
// value already exists.
func (p *Panel) AddTarget(t *Target) {
	p.CommonPanel.Targets = append(p.CommonPanel.Targets, *t)
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
	setTarget(t, &p.CommonPanel.Targets)
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
	repeatDS(dsNames, &p.CommonPanel.Targets)
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
	repeatTarget(dsNames, &p.CommonPanel.Targets)
}

// GetTargets is iterate over all panel targets. It just returns nil if
// no targets defined for panel of concrete type.
func (p *Panel) GetTargets() *[]Target {
	return &p.CommonPanel.Targets
}

type probePanel struct {
	CommonPanel
	//	json.RawMessage
}

func (p *Panel) UnmarshalJSON(b []byte) (err error) {
	var probe probePanel
	if err = json.Unmarshal(b, &probe); err == nil {
		p.CommonPanel = probe.CommonPanel
		switch probe.Type {
		case "row":
			var row RowPanel
			p.OfType = RowType
			if err = json.Unmarshal(b, &row); err == nil {
				p.RowPanel = &row
			}
		default:
			var custom = make(CustomPanel)
			p.OfType = CustomType
			if err = json.Unmarshal(b, &custom); err == nil {
				p.CustomPanel = &custom
			}
			for _, key := range []string{
				"gridPos",
				"collapsed",
				"datasource",
				"description",
				"editable",
				"error",
				"height",
				"hideTimeOverride",
				"isNew",
				"links",
				"minSpan",
				"renderer",
				"repeat",
				"repeatPanelId",
				"repeatedByRow",
				"scopedVars",
				"span",
				"transparent",
				"alert",
				"targets",
			} {
				delete(*p.CustomPanel, key)
			}
		}
	}
	return
}

func (p *Panel) MarshalJSON() ([]byte, error) {
	switch p.OfType {
	case RowType:
		var outRow = struct {
			CommonPanel
			RowPanel
		}{p.CommonPanel, *p.RowPanel}
		return json.Marshal(outRow)
	case CustomType:
		var outCustom = customPanelOutput{
			p.CommonPanel,
			*p.CustomPanel,
		}
		return json.Marshal(outCustom)
	}
	return nil, errors.New("can't marshal unknown panel type")
}

type customPanelOutput struct {
	CommonPanel
	CustomPanel
}

func (c customPanelOutput) MarshalJSON() ([]byte, error) {
	b, err := json.Marshal(c.CommonPanel)
	if err != nil {
		return b, err
	}
	// Append custom keys to marshalled CommonPanel.
	buf := bytes.NewBuffer(b[:len(b)-1])

	for k, v := range c.CustomPanel {
		buf.WriteString(`,"`)
		buf.WriteString(k)
		buf.WriteString(`":`)
		b, err := json.Marshal(v)
		if err != nil {
			return b, err
		}
		buf.Write(b)
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

func incRefID(refID string) string {
	firstLetter := refID[0]
	ordinal := int(firstLetter)
	ordinal++
	return string(rune(ordinal))
}
