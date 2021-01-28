package sqlhelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	Input  uint32
	Expect string
}

func TestMaskIntToString(t *testing.T) {
	// mask int to string  16 -> 255.255.0.0
	cases := []testCase{
		{
			Input:  16,
			Expect: "255.255.0.0",
		},
		{
			Input:  8,
			Expect: "255.0.0.0",
		},
		{
			Input:  31,
			Expect: "255.255.255.254",
		},
		{ // invalid
			Input:  33,
			Expect: "",
		},
	}
	for _, caze := range cases {
		assert.EqualValues(t, caze.Expect, MaskIntToString(caze.Input))
	}
	return
}

func TestMaskStringToInt(t *testing.T) {
	cases := []testCase{
		{
			Input:  16,
			Expect: "255.255.0.0",
		}, {
			Input:  8,
			Expect: "255.0.0.0",
		}, {
			Input:  0,
			Expect: "0.0.0.0",
		}, {
			Input:  31,
			Expect: "255.255.255.254",
		}, {
			Input:  32,
			Expect: "255.255.255.255",
		}, { // invalid
			Input:  0,
			Expect: "0.0.0.0.0",
		}, { // invalid
			Input:  0,
			Expect: "",
		},
	}
	for _, caze := range cases {
		assert.EqualValues(t, caze.Input, MaskStringToInt(caze.Expect))
	}
	return
}
