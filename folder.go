package sdk

/*
author: bilc_dev@163.com
*/

import (
	"encoding/json"
	"fmt"
	"time"
)

type Folder struct {
	UID   string `json:"uid"`
	Title string `json:"title"`
}

type FolderProperties struct {
	ID        int       `json:"id"`
	UID       string    `json:"uid"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	HasACL    bool      `json:"hasAcl"`
	CanSave   bool      `json:"canSave"`
	CanEdit   bool      `json:"canEdit"`
	CanAdmin  bool      `json:"canAdmin"`
	CreatedBy string    `json:"createdBy"`
	Created   time.Time `json:"created"`
	UpdatedBy string    `json:"updatedBy"`
	Updated   time.Time `json:"updated"`
	Version   int       `json:"version"`
}

func NewFolder(uid, title string) Folder {
	return Folder{
		UID:   uid,
		Title: title,
	}
}

func (r *Client) CreateFolder(f Folder) (resp FolderProperties, err error) {
	raw, _ := json.Marshal(f)
	if raw, _, err = r.post("api/folders", nil, raw); err != nil {
		return resp, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return resp, fmt.Errorf("%v %s", err, string(raw))
	}
	return resp, nil
}

func (r *Client) GetFolder(uid string) (resp FolderProperties, err error) {

	var (
		raw  []byte
		code int
	)
	if raw, code, err = r.get("api/folders/"+uid, nil); err != nil {
		return resp, err
	}
	if code != 200 {
		return resp, fmt.Errorf("HTTP error %d: returns %s", code, raw)
	}
	err = json.Unmarshal(raw, &resp)
	return resp, err
}
