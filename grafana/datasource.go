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

// Datasource as described in the doc
// http://docs.grafana.org/reference/http_api/#get-all-datasources
type Datasource struct {
	ID                uint        `json:"id"`
	OrgID             uint        `json:"orgId"`
	Name              string      `json:"name"`
	Type              string      `json:"type"`
	Access            string      `json:"access"` // direct or proxy
	URL               string      `json:"url"`
	Password          *string     `json:"password,omitempty"`
	User              *string     `json:"user,omitempty"`
	Database          *string     `json:"database,omitempty"`
	BasicAuth         *bool       `json:"basicAuth,omitempty"`
	BasicAuthUser     *string     `json:"basicAuthUser,omitempty"`
	BasicAuthPassword *string     `json:"basicAuthPassword,omitempty"`
	IsDefault         bool        `json:"isDefault"`
	JSONData          interface{} `json:"jsonData"`
}
