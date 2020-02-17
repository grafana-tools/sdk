package sdk

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// https://grafana.com/docs/grafana/latest/http_api/annotations/

// CreateAnnotation creates a new annotation from the annotation request
func (r *Client) CreateAnnotation(a CreateAnnotationRequest) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(a); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.post("api/annotations", nil, raw); err != nil {
		return StatusMessage{}, fmt.Errorf("failed to unmarshal response message: %w", err)
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, fmt.Errorf("failed to unmarshal response message: %w", err)
	}
	return resp, nil
}

// PatchAnnotation patches the annotation with id with the request
func (r *Client) PatchAnnotation(id uint, a PatchAnnotationRequest) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(a); err != nil {
		return StatusMessage{}, err
	}
	if raw, _, err = r.patch(fmt.Sprintf("api/annotations/%d", id), nil, raw); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// GetAnnotations gets annotations matching the annotation parameters
func (r *Client) GetAnnotations(params ...GetAnnotationsParams) ([]AnnotationResponse, error) {
	var (
		raw           []byte
		err           error
		resp          []AnnotationResponse
		requestParams = make(url.Values)
	)

	for _, p := range params {
		p(requestParams)
	}

	if raw, _, err = r.get("api/annotations", requestParams); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteAnnotation deletes the annotation with id
func (r *Client) DeleteAnnotation(id uint) (StatusMessage, error) {
	var (
		raw  []byte
		err  error
		resp StatusMessage
	)

	if raw, _, err = r.delete(fmt.Sprintf("api/annotations/%d", id)); err != nil {
		return StatusMessage{}, err
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, err
	}
	return resp, nil
}

// AnnotationOption is the type for all options implementing query parameters
// https://grafana.com/docs/grafana/latest/http_api/annotations/#find-annotations
type GetAnnotationsParams func(values url.Values)

func WithTag(tag string) GetAnnotationsParams {
	return func(v url.Values) {
		v.Add("tags", tag)
	}
}

func WithLimit(limit uint) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("limit", strconv.FormatUint(uint64(limit), 10))
	}
}

func WithAnnotationType() GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("type", "annotation")
	}
}

func WithDashboard(id uint) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("dashboardId", strconv.FormatUint(uint64(id), 10))
	}
}

func WithPanel(id uint) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("panelId", strconv.FormatUint(uint64(id), 10))
	}
}

func WithUser(id uint) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("userId", strconv.FormatUint(uint64(id), 10))
	}
}

func WithFrom(t time.Time) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("from", strconv.FormatInt(toMilliseconds(t), 10))
	}
}

func WithTo(t time.Time) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("to", strconv.FormatInt(toMilliseconds(t), 10))
	}
}

func toMilliseconds(t time.Time) int64 {
	return t.UnixNano() / 1_000_000
}
