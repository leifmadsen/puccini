package common

import (
	"unsafe"
)

// See:
// https://github.com/golang/go/issues/25484
// https://github.com/golang/go/issues/19367
// https://golang.org/src/strings/builder.go#L45
func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
