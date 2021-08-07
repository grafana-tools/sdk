package sdk

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotAccessDenied = errors.New("access denied")
	ErrNotAuthorized   = errors.New("not authorized")
	ErrCannotCreate    = errors.New("cannot create; see body for details")
)
