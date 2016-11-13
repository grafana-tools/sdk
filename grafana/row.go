package grafana

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

// Row represents single row of Grafana dashboard.
type Row struct {
	Title     string  `json:"title"`
	ShowTitle bool    `json:"showTitle"`
	Collapse  bool    `json:"collapse"`
	Editable  bool    `json:"editable"`
	Height    height  `json:"height"`
	Panels    []Panel `json:"panels"`
	board     *Board
}

func (r *Row) Add(panel *Panel) {
	r.board.lastPanelID++
	panel.ID = r.board.lastPanelID
	r.Panels = append(r.Panels, *panel)
}

func (r *Row) AddDashlist(data *DashlistPanel) {
	r.board.lastPanelID++
	panel := NewDashlist("")
	panel.ID = r.board.lastPanelID
	panel.DashlistPanel = data
	r.Panels = append(r.Panels, *panel)
}

func (r *Row) AddGraph(data *GraphPanel) {
	r.board.lastPanelID++
	panel := NewGraph("")
	panel.ID = r.board.lastPanelID
	panel.GraphPanel = data
	r.Panels = append(r.Panels, *panel)
}

func (r *Row) AddTable(data *TablePanel) {
	r.board.lastPanelID++
	panel := NewTable("")
	panel.ID = r.board.lastPanelID
	panel.TablePanel = data
	r.Panels = append(r.Panels, *panel)
}

func (r *Row) AddText(data *TextPanel) {
	r.board.lastPanelID++
	panel := NewText("")
	panel.ID = r.board.lastPanelID
	panel.TextPanel = data
	r.Panels = append(r.Panels, *panel)
}

func (r *Row) AddSinglestat(data *SinglestatPanel) {
	r.board.lastPanelID++
	panel := NewSinglestat("")
	panel.ID = r.board.lastPanelID
	panel.SinglestatPanel = data
	r.Panels = append(r.Panels, *panel)
}

func (r *Row) AddCustom(data *CustomPanel) {
	r.board.lastPanelID++
	panel := NewCustom("")
	panel.ID = r.board.lastPanelID
	panel.CustomPanel = data
	r.Panels = append(r.Panels, *panel)
}
