package model

import (
	"fmt"
	"testing"
)

func TestAppendParseTime(t *testing.T) {
	dsn := "root:123345@tcp(127.0.0.1:3307)/abac_test?charset=utf8mb4&collation=utf8mb4_unicode_ci"
	fmt.Println(AppendParseTime(dsn))
}
