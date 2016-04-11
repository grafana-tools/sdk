package client

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
	"fmt"

	"github.com/grafov/autograf/grafana"
)

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

func (r *Instance) CreateDatasource(ds grafana.Datasource) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(ds); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post("api/datasources", nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

func (r *Instance) UpdateDatasource(ds grafana.Datasource) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(ds); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post(fmt.Sprintf("api/datasources/%d", ds.ID), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

func (r *Instance) DeleteDatasource(id uint) (StatusMessage, error) {
	var (
		raw   []byte
		reply StatusMessage
		err   error
	)
	if raw, err = r.delete(fmt.Sprintf("api/datasources/%d", id)); err != nil {
		return StatusMessage{}, err
	}
	err = json.Unmarshal(raw, &reply)
	return reply, err
}
