package cast

import (
	"encoding/json"
	"math"
	"strconv"
	"testing"
)

func TestToInt(t *testing.T) {
	const expected = 2354
	v, _ := ToInt("2354")
	if v != expected {
		t.Error()
	}
	v, _ = ToInt(expected)
	if v != expected {
		t.Error()
	}
	if v, _ := ToInt(""); v != 0 {
		t.Error()
	}
	if _, err := ToInt(`""`); err == nil {
		t.Error()
	}
	if v, _ := ToInt(uint16(123)); v != 123 {
		t.Error()
	}
	_, err := ToInt([]int{3})
	if err.Error() != "cast: invalid value `[3]`, int required" {
		t.Error()
	}
}

func TestToUint(t *testing.T) {
	const expected = 2354
	if v, _ := ToUint(uint16(expected)); v != expected {
		t.Error()
	}
	v, _ := ToUint("2354")
	if v != expected {
		t.Error()
	}
	v, _ = ToUint(expected)
	if v != expected {
		t.Error()
	}
	if v, _ := ToUint(""); v != 0 {
		t.Error()
	}
	if _, err := ToUint(`""`); err == nil {
		t.Error()
	}
	if v, _ := ToUint(float32(expected)); v != expected {
		t.Error()
	}
	_, err := ToUint(1.12)
	if err.Error() != "cast: bad value `1.12`, type is uint64" {
		t.Error()
	}
	_, err = ToUint([]int{3})
	if err.Error() != "cast: invalid value `[3]`, uint(slice) required" {
		t.Error()
	}
	_, err = ToUint(-12)
	if err.Error() != "cast: bad value `-12`, type is uint64" {
		t.Error()
	}
}

func TestToFloat(t *testing.T) {
	const e = 123
	v, _ := ToFloat("123")
	if v != e {
		t.Error()
	}
	v, _ = ToFloat(123)
	if v != e {
		t.Error()
	}
	if v, _ := ToFloat(uint16(123)); v != e {
		t.Error()
	}
	if v, _ := ToFloat(1.23); v != 1.23 {
		t.Error()
	}
	if v, _ := ToFloat(""); v != 0 {
		t.Error()
	}
	if _, err := ToFloat(`""`); err == nil {
		t.Error()
	}
	_, err := ToFloat([]int{1})
	if err.Error() != "cast: invalid value `[1]`, float required" {
		t.Error()
	}
}

func TestToBool(t *testing.T) {
	if v, _ := ToBool(true); !v {
		t.Error()
	}
	if v, _ := ToBool(false); v {
		t.Error()
	}
	if v, _ := ToBool(1); !v {
		t.Error()
	}
	if v, _ := ToBool(0); v {
		t.Error()
	}
	if v, _ := ToBool("true"); !v {
		t.Error()
	}
	if v, _ := ToBool("false"); v {
		t.Error()
	}
	if _, err := ToBool("hello"); err == nil {
		t.Error()
	}
	if v, _ := ToBool(uint8(1)); !v {
		t.Error()
	}
	if v, _ := ToBool(1.34); !v {
		t.Error()
	}
	if v, _ := ToBool(float32(0)); v {
		t.Error()
	}
	_, err := ToBool([]int{0})
	if err.Error() != "cast: invalid value `[0]`, bool required" {
		t.Error()
	}
}

func TestOverflowInt(t *testing.T) {
	if i, _ := ToInt(uint16(90)); i != 90 {
		t.Error()
	}
	_, err := ToInt(uint64(12345994566764549642))
	if err.Error() != "cast: bad value `12345994566764549642`, type is int64" {
		t.Errorf("err: %s", err)
	}
}

func TestBadUint(t *testing.T) {
	if i, _ := ToUint(int(234)); i != 234 {
		t.Error()
	}
	_, err := ToUint(int(-234))
	if err.Error() != "cast: bad value `-234`, type is uint64" {
		t.Error()
	}
}

func TestFloatToUint(t *testing.T) {
	_, e := ToUint(float32(-7))
	if e.Error() != "cast: bad value `-7`, type is uint64" {
		t.Error()
	}
	m, _ := ToUint(float32(234995))
	if m != 234995 {
		t.Error()
	}
	_, e = ToUint(float32(4.12))
	if e.Error() != "cast: bad value `4.119999885559082`, type is uint64" {
		t.Error()
	}
}

func TestFloatToInt(t *testing.T) {
	_, e := ToInt(float64(235342343245346456456456))
	if e.Error() != "cast: bad value `2.3534234324534644e+23`, type is int64" {
		t.Error()
	}
	v, _ := ToInt(float64(90.0))
	if v != 90 {
		t.Error()
	}
	v, _ = ToInt(float64(-23543))
	if v != -23543 {
		t.Error()
	}
	_, e = ToInt(float64(4.125))
	if e.Error() != "cast: bad value `4.125`, type is int64" {
		t.Error()
	}
	_, e = ToInt(-1e+34)
	if e.Error() != "cast: bad value `-1e+34`, type is int64" {
		t.Error()
	}
}

func TestJSON(t *testing.T) {
	var v map[string]interface{}
	json.Unmarshal([]byte(`{"A": 3242342312353426}`), &v)
	u, _ := ToUint(v["A"])
	if u != 3242342312353426 {
		t.Error("json convert")
	}
}

func toIntAssert(d interface{}) (int64, error) {
	switch v := d.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		if v > math.MaxInt64 {
			return 0, badValue(v, "int64")
		}
		return int64(v), nil
	case float32:
		n := int64(v)
		if v != float32(n) {
			return 0, badValue(v, "int64")
		}
		return n, nil
	case float64:
		n := int64(v)
		if v != float64(n) {
			return 0, badValue(v, "int64")
		}
		return n, nil
	case string:
		if v == "" {
			return 0, nil
		}
		return strconv.ParseInt(v, 10, 64)
	}
	return 0, valueError(d, "int")
}

func BenchmarkToInt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ToInt("1234")
	}
}

func BenchmarkToIntAssert(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		toIntAssert("1234")
	}
}
