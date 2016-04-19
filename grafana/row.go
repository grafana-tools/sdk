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
	Panels    []Panel `json:"panels,omitempty"`
	board     *Board
}

func (r *Row) AddDashlist(data *DashlistPanel) {
	r.board.lastPanelID++
	data.ID = r.board.lastPanelID
	panel := Panel{OfType: DashlistType, DashlistPanel: data}
	r.Panels = append(r.Panels, panel)
}

func (r *Row) AddGraph(data *GraphPanel) {
	r.board.lastPanelID++
	data.ID = r.board.lastPanelID
	panel := Panel{OfType: GraphType, GraphPanel: data}
	r.Panels = append(r.Panels, panel)
}

func (r *Row) AddTable(data *TablePanel) {
	r.board.lastPanelID++
	data.ID = r.board.lastPanelID
	panel := Panel{OfType: TableType, TablePanel: data}
	r.Panels = append(r.Panels, panel)
}

func (r *Row) AddText(data *TextPanel) {
	r.board.lastPanelID++
	data.ID = r.board.lastPanelID
	panel := Panel{OfType: TextType, TextPanel: data}
	r.Panels = append(r.Panels, panel)
}

func (r *Row) AddSinglestat(data *SinglestatPanel) {
	r.board.lastPanelID++
	data.ID = r.board.lastPanelID
	panel := Panel{OfType: SinglestatType, SinglestatPanel: data}
	r.Panels = append(r.Panels, panel)
}

func (r *Row) AddCustom(data *CustomPanel) {
	r.board.lastPanelID++
	(*data)["id"] = r.board.lastPanelID
	panel := Panel{OfType: CustomType, CustomPanel: data}
	r.Panels = append(r.Panels, panel)
}
