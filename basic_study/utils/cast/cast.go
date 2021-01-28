// Package cast implements reflection based type casting.
package cast

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

// Type convertion.
// v, err := cast.ToInt("234")

func isInt(v reflect.Value) bool {
	k := v.Kind()
	return k == reflect.Int ||
		k == reflect.Int8 ||
		k == reflect.Int16 ||
		k == reflect.Int32 ||
		k == reflect.Int64
}

func isUint(v reflect.Value) bool {
	k := v.Kind()
	return k == reflect.Uint ||
		k == reflect.Uint8 ||
		k == reflect.Uint16 ||
		k == reflect.Uint32 ||
		k == reflect.Uint64
}

func isFloat(v reflect.Value) bool {
	k := v.Kind()
	return k == reflect.Float32 || k == reflect.Float64
}

func isString(v reflect.Value) bool {
	return v.Kind() == reflect.String
}

func isBool(v reflect.Value) bool {
	return v.Kind() == reflect.Bool
}

func valueError(d interface{}, actual string) error {
	return fmt.Errorf("cast: invalid value `%v`, %s required", d, actual)
}

func badValue(d interface{}, t string) error {
	return fmt.Errorf("cast: bad value `%v`, type is %s", d, t)
}

// ToFloat converts interface to float64.
// Note this function will convert empty string `""` to 0.
func ToFloat(d interface{}) (float64, error) {
	v := reflect.ValueOf(d)
	switch {
	case isUint(v):
		return float64(v.Uint()), nil
	case isInt(v):
		return float64(v.Int()), nil
	case isFloat(v):
		return v.Float(), nil
	case isString(v):
		s := v.String()
		if s == "" {
			return 0, nil
		}
		return strconv.ParseFloat(s, 64)
	}
	return 0, valueError(d, "float")
}

// ToUint converts interface to uint64.
// Note this function will convert empty string `""` to 0.
func ToUint(d interface{}) (uint64, error) {
	v := reflect.ValueOf(d)
	switch {
	case isFloat(v):
		f := v.Float()
		n := uint64(f)
		if f != float64(n) {
			return 0, badValue(f, "uint64")
		}
		return n, nil
	case isUint(v):
		return v.Uint(), nil
	case isInt(v):
		r := v.Int()
		if r < 0 {
			return 0, badValue(r, "uint64")
		}
		return uint64(r), nil
	case isString(v):
		s := v.String()
		if s == "" {
			return 0, nil
		}
		return strconv.ParseUint(s, 10, 64)
	}
	return 0, valueError(d, fmt.Sprintf("uint(%s)", v.Kind()))
}

// ToBool converts interface to bool.
func ToBool(d interface{}) (bool, error) {
	v := reflect.ValueOf(d)
	switch {
	case isBool(v):
		return v.Bool(), nil
	case isUint(v):
		return v.Uint() != 0, nil
	case isInt(v):
		return v.Int() != 0, nil
	case isFloat(v):
		return v.Float() != 0, nil
	case isString(v):
		value := v.String()
		if value == "true" {
			return true, nil
		}
		if value == "false" {
			return false, nil
		}
	}
	return false, valueError(d, "bool")
}

// ToInt converts interface to int64.
// Note this function will convert empty string `""` to 0.
func ToInt(d interface{}) (int64, error) {
	v := reflect.ValueOf(d)
	switch {
	case isFloat(v):
		f := v.Float()
		n := int64(f)
		if f != float64(n) {
			return 0, badValue(f, "int64")
		}
		return n, nil
	case isInt(v):
		return v.Int(), nil
	case isUint(v):
		r := v.Uint()
		if r > math.MaxInt64 {
			return 0, badValue(r, "int64")
		}
		return int64(r), nil
	case isString(v):
		s := v.String()
		if s == "" {
			return 0, nil
		}
		return strconv.ParseInt(s, 10, 64)
	}
	return 0, valueError(d, "int")
}
