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
	Title    string  `json:"title"`
	Collapse bool    `json:"collapse"`
	Editable bool    `json:"editable"`
	Height   height  `json:"height"`
	Panels   []Panel `json:"panels,omitempty"`
}

func NewRow() *Row {
	return &Row{
		Title:    "New row",
		Collapse: false,
		Editable: true,
		Height:   "250px"}
}

func (r *Row) AddPanel(oftype panelType, title string) {
	panel := Panel{OfType: oftype}
	switch oftype {
	case DashlistPanel:
		panel.dashlistPanel = &dashlistPanel{}
		panel.dashlistPanel.Title = title
	case GraphPanel:
		panel.graphPanel = &graphPanel{}
		panel.graphPanel.Title = title
	case TablePanel:
		panel.tablePanel = &tablePanel{}
		panel.tablePanel.Title = title
	case TextPanel:
		panel.textPanel = &textPanel{}
		panel.textPanel.Title = title
	case SinglestatPanel:
		panel.singlestatPanel = &singlestatPanel{}
		panel.singlestatPanel.Title = title
	case CustomPanel:
		panel.customPanel = &customPanel{}
	}
	r.Panels = append(r.Panels, panel)
}
