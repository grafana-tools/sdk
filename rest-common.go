package sdk

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrNotFound        = errors.New("not found")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotAccessDenied = errors.New("access denied")
	ErrNotAuthorized   = errors.New("not authorized")
	ErrCannotCreate    = errors.New("cannot create; see body for details")
)

func httpStatusCodeError(code int, message string, raw []byte) error {
	switch code {
	case http.StatusNotFound:
		return fmt.Errorf("%s: %w", message, ErrNotFound)

	case http.StatusForbidden:
		return fmt.Errorf("%s: %w", message, ErrNotAccessDenied)

	case http.StatusUnauthorized:
		return fmt.Errorf("%s: %w", message, ErrNotAuthorized)

	case http.StatusPreconditionFailed:
		return fmt.Errorf("%s: %w", message, ErrCannotCreate)

	default:
		return fmt.Errorf("%s returned HTTP status code %d: %v", message, code, raw)
	}
}
