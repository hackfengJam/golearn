package cast

import (
	"reflect"
	"unsafe"
)

// StringToBytes converts a string to bytes without copy.
func StringToBytes(s string) (b []byte) {
	h := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	h.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	h.Len = len(s)
	h.Cap = len(s)
	return
}

// BytesToString converts a byte array to string without copy.
func BytesToString(b []byte) (s string) {
	h := (*reflect.StringHeader)(unsafe.Pointer(&s))
	h.Data = (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data
	h.Len = len(b)
	return
}
