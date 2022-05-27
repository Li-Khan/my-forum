package mistake

import "errors"

var (
	ErrNotFound     = errors.New("404 not found")
	ErrBadRequest   = errors.New("400 bad request")
	ErrConflict     = errors.New("409 conflict")
	ErrUnauthorized = errors.New("401 unauthorized")
)
