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
		return StatusMessage{}, fmt.Errorf("failed to marshal request: %w", err)
	}
	if raw, _, err = r.post("api/annotations", nil, raw); err != nil {
		return StatusMessage{}, fmt.Errorf("failed to create annotation: %w", err)
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
		return StatusMessage{}, fmt.Errorf("failed to marshal request: %w", err)
	}
	if raw, _, err = r.patch(fmt.Sprintf("api/annotations/%d", id), nil, raw); err != nil {
		return StatusMessage{}, fmt.Errorf("failed to patch annotation: %w", err)
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, fmt.Errorf("failed to unmarshal response message: %w", err)
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
		return nil, fmt.Errorf("failed to get annotations: %w", err)
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response message: %w", err)
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
		return StatusMessage{}, fmt.Errorf("failed to delete annotation: %w", err)
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, fmt.Errorf("failed to unmarshal response message: %w", err)
	}
	return resp, nil
}

// AnnotationOption is the type for all options implementing query parameters
// https://grafana.com/docs/grafana/latest/http_api/annotations/#find-annotations
type GetAnnotationsParams func(values url.Values)

// WithTag adds the tag to the
func WithTag(tag string) GetAnnotationsParams {
	return func(v url.Values) {
		v.Add("tags", tag)
	}
}

// WithLimit sets the max number of alerts to return
func WithLimit(limit uint) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("limit", strconv.FormatUint(uint64(limit), 10))
	}
}

// WithAnnotationType filters the type to annotations
func WithAnnotationType() GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("type", "annotation")
	}
}

// WithAlertType filters the type to alerts
func WithAlertType() GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("type", "alert")
	}
}

// WithDashboard filters the response to the specified dashboard ID
func WithDashboard(id uint) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("dashboardId", strconv.FormatUint(uint64(id), 10))
	}
}

// WithPanel filters the response to the specified panel ID
func WithPanel(id uint) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("panelId", strconv.FormatUint(uint64(id), 10))
	}
}

// WithUser filters the annotations to only be made by the specified user ID
func WithUser(id uint) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("userId", strconv.FormatUint(uint64(id), 10))
	}
}

// WithStartTime filters the annotations to after the specified time
func WithStartTime(t time.Time) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("from", strconv.FormatInt(toMilliseconds(t), 10))
	}
}

// WithEndTime filters the annotations to before the specified time
func WithEndTime(t time.Time) GetAnnotationsParams {
	return func(v url.Values) {
		v.Set("to", strconv.FormatInt(toMilliseconds(t), 10))
	}
}

func toMilliseconds(t time.Time) int64 {
	return t.UnixNano() / 1_000_000
}
