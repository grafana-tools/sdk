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
