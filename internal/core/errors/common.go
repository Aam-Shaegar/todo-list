package core_error

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrInvlalidArgument = errors.New("invalid argument")
	ErrConflict         = errors.New("conflict")
)
