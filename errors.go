package sdk

// ErrNotFound is the error returned when 404 HTTP status returned
type ErrNotFound struct {
	Message string
}

// Error implements the error interface for ErrNotFound.
func (e ErrNotFound) Error() string {
	return e.Message
}
