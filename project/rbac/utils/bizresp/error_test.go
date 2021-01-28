package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	e := NewError("")
	assert.Equal(t, "enigma.system.UnknownError", e.Error())

	e = NewError(ErrRequestInvalid)
	assert.Equal(t, ErrRequestInvalid, e.Error())
}
