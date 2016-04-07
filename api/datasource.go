package api

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
	"encoding/json"

	"github.com/grafov/autograf/grafana"
)

func (r *Instance) CreateDatasource(ds grafana.Datasource) {
}

func (r *Instance) UpdateDatasource(ds grafana.Datasource) {
}

func (r *Instance) GetAllDatasources() ([]grafana.Datasource, error) {
	var (
		raw []byte
		dss []grafana.Datasource
		err error
	)
	if raw, err = r.get("api/datasources", nil); err != nil {
		return nil, err
	}
	err = json.Unmarshal(raw, &dss)
	return dss, err
}
