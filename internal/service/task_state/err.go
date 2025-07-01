package taskstate

import "errors"

var (
	ErrInvalidTransition = errors.New("invalid status transition")
)
