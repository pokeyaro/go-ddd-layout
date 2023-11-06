package errors

import (
	"errors"
)

var (
	// ErrInvalidAuthorizationHeader 无效的授权头错误
	ErrInvalidAuthorizationHeader = errors.New("invalid Authorization header")
)
