package env

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	v := Get("A", "B")
	if v != "" {
		t.Error(v)
	}

	os.Setenv("B", "1")

	v = Get("A", "B")
	if v != "1" {
		t.Error(v)
	}

	os.Setenv("A", "2")

	v = Get("A", "B")
	if v != "2" {
		t.Error(v)
	}
}
