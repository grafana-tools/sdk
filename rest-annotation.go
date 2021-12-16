package sdk

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// https://grafana.com/docs/grafana/latest/http_api/annotations/

// CreateAnnotation creates a new annotation from the annotation request
func (r *Client) CreateAnnotation(ctx context.Context, a CreateAnnotationRequest) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(a); err != nil {
		return StatusMessage{}, errors.Wrap(err, "marshal request")
	}
	if raw, _, err = r.post(ctx, "api/annotations", raw); err != nil {
		return StatusMessage{}, errors.Wrap(err, "create annotation")
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, errors.Wrap(err, "unmarshal response message")
	}
	return resp, nil
}

// PatchAnnotation patches the annotation with id with the request
func (r *Client) PatchAnnotation(ctx context.Context, id uint, a PatchAnnotationRequest) (StatusMessage, error) {
	var (
		raw  []byte
		resp StatusMessage
		err  error
	)
	if raw, err = json.Marshal(a); err != nil {
		return StatusMessage{}, errors.Wrap(err, "marshal request")
	}
	if raw, _, err = r.patch(ctx, fmt.Sprintf("api/annotations/%d", id), raw); err != nil {
		return StatusMessage{}, errors.Wrap(err, "patch annotation")
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, errors.Wrap(err, "unmarshal response message")
	}
	return resp, nil
}

// GetAnnotations gets annotations matching the annotation parameters
func (r *Client) GetAnnotations(ctx context.Context, params ...GetAnnotationsParams) ([]AnnotationResponse, error) {
	var (
		raw       []byte
		err       error
		resp      []AnnotationResponse
		modifiers []APIRequestModifier
	)

	for _, params := range params {
		modifiers = append(modifiers, APIRequestModifier(params))
	}

	if raw, _, err = r.get(ctx, "api/annotations", modifiers...); err != nil {
		return nil, errors.Wrap(err, "get annotations")
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return nil, errors.Wrap(err, "unmarshal response message")
	}
	return resp, nil
}

// DeleteAnnotation deletes the annotation with id
func (r *Client) DeleteAnnotation(ctx context.Context, id uint) (StatusMessage, error) {
	var (
		raw  []byte
		err  error
		resp StatusMessage
	)

	if raw, _, err = r.delete(ctx, fmt.Sprintf("api/annotations/%d", id)); err != nil {
		return StatusMessage{}, errors.Wrap(err, "delete annotation")
	}
	if err = json.Unmarshal(raw, &resp); err != nil {
		return StatusMessage{}, errors.Wrap(err, "unmarshal response message")
	}
	return resp, nil
}

// GetAnnotationsParams is the type for all options implementing query parameters
// https://grafana.com/docs/grafana/latest/http_api/annotations/#find-annotations
type GetAnnotationsParams APIRequestModifier

// WithTag adds the tag to the
func WithTag(tag string) GetAnnotationsParams {
	return func(req *http.Request) {
		values := req.URL.Query()
		values.Add("tags", tag)
		req.URL.RawQuery = values.Encode()
	}
}

// WithLimit sets the max number of alerts to return
func WithLimit(limit uint) GetAnnotationsParams {
	return func(req *http.Request) {
		values := req.URL.Query()
		values.Set("limit", strconv.FormatUint(uint64(limit), 10))
		req.URL.RawQuery = values.Encode()
	}
}

// WithAnnotationType filters the type to annotations
func WithAnnotationType() GetAnnotationsParams {
	return func(req *http.Request) {
		values := req.URL.Query()
		values.Set("type", "annotation")
		req.URL.RawQuery = values.Encode()
	}
}

// WithAlertType filters the type to alerts
func WithAlertType() GetAnnotationsParams {
	return func(req *http.Request) {
		values := req.URL.Query()
		values.Set("type", "alert")
		req.URL.RawQuery = values.Encode()
	}
}

// WithDashboard filters the response to the specified dashboard ID
func WithDashboard(id uint) GetAnnotationsParams {
	return func(req *http.Request) {
		values := req.URL.Query()
		values.Set("dashboardId", strconv.FormatUint(uint64(id), 10))
		req.URL.RawQuery = values.Encode()
	}
}

// WithPanel filters the response to the specified panel ID
func WithPanel(id uint) GetAnnotationsParams {
	return func(req *http.Request) {
		values := req.URL.Query()
		values.Set("panelId", strconv.FormatUint(uint64(id), 10))
		req.URL.RawQuery = values.Encode()
	}
}

// WithUser filters the annotations to only be made by the specified user ID
func WithUser(id uint) GetAnnotationsParams {
	return func(req *http.Request) {
		values := req.URL.Query()
		values.Set("userId", strconv.FormatUint(uint64(id), 10))
		req.URL.RawQuery = values.Encode()
	}
}

// WithStartTime filters the annotations to after the specified time
func WithStartTime(t time.Time) GetAnnotationsParams {
	return func(req *http.Request) {
		values := req.URL.Query()
		values.Set("from", strconv.FormatInt(toMilliseconds(t), 10))
		req.URL.RawQuery = values.Encode()
	}
}

// WithEndTime filters the annotations to before the specified time
func WithEndTime(t time.Time) GetAnnotationsParams {
	return func(req *http.Request) {
		values := req.URL.Query()
		values.Set("to", strconv.FormatInt(toMilliseconds(t), 10))
		req.URL.RawQuery = values.Encode()
	}
}

func toMilliseconds(t time.Time) int64 {
	return t.UnixNano() / 1000000
}
