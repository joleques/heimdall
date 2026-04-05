package shared

import "errors"

var (
	ErrInvalidCommand  = errors.New("invalid command")
	ErrInvalidTarget   = errors.New("invalid target")
	ErrInvalidArgument = errors.New("invalid argument")
)
