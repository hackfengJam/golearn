package common

import "errors"

var (
	ERR_LOCK_ALREADY_REQUIRED = errors.New("所以被占用")
)
