package common

import "errors"

var (
	ERR_LOCK_ALREADY_REQUIRED = errors.New("所以被占用")

	ERR_NO_LOCAL_IP_FOUND = errors.New("没找到网卡IP")
)
