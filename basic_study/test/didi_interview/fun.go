package didi

import (
	"errors"
	"strings"
)

/*
example:
- -1
- abc
- 123
- +123
*/
const MaxInt32 = "2147483648"

func T(str string, base int32) (ret int32, err error) {
	b := []byte(str)
	// 2147483648
	if len(b) > len(MaxInt32) {
		return 0, errors.New("invalid string number, overflow")
	}
	if len(b) == len(MaxInt32) && strings.Compare(str, MaxInt32) > 0 {
		return 0, errors.New("invalid string number, overflow")
	}

	ret = 0
	for _, b := range []byte(str) {
		if b < '0' || b > '9' {
			return 0, errors.New("invalid string number")
		}
		loc := int(b - '0')
		ret = ret*base + loc
	}
	return
}
