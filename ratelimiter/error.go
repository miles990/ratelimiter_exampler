package ratelimiter

import "errors"

var (
	// ErrOverLimit 超過限制
	ErrOverLimit = errors.New("over limit error")
)
