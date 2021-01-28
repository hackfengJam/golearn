package cast

import (
	"errors"
	"testing"
)

func TestToMap(t *testing.T) {
	sample := struct {
		A int32   `redis:"a"`
		B int64   `redis:"-"`
		C float64 `redis:"cccc"`
		D uint16
		E string `redis:"zzz"`
		F bool
	}{
		A: 12,
		B: 13,
		C: 12.22,
		D: 1232,
		E: "fasdfads",
		F: true,
	}

	_, err := ToMap(nil)
	if err == nil {
		t.Fatal(errors.New("nil src should return errror"))
	}
	m, err := ToMapWithTag(&sample, "redis")
	if err != nil {
		t.Fatal(err)
	}
	if m["a"].(int32) != 12 || m["zzz"].(string) != "fasdfads" || m["D"].(uint16) != 1232 {
		t.Errorf("wrong map value %+v", m)
	}
}

func TestScan(t *testing.T) {

	var a struct {
		A int32
		B int64
		C float64
		D uint16
		E string
		F bool
	}
	m := map[string]string{"A": "12", "B": "23423423", "C": "23.33", "D": "34", "E": "asdfasd", "F": "1"}

	err := ScanStruct(m, &a)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("a is %#v", a)

	var b struct {
		A *int
	}
	err = ScanStruct(m, &b)
	if err == nil {
		t.Fail()
	}
	t.Log(err)

	var c struct {
		A struct {
			A int32
		}
	}
	err = ScanStruct(m, &c)
	if err == nil {
		t.Fail()
	}
	t.Log(err)

	var tg struct {
		A int32   `redis:"a"`
		B int64   `cast:"bb"`
		C float64 `redis:"cccc"`
		D uint16  `redis:"-"`
		E string  `redis:"zzz"`
		F bool
	}
	err = ScanStructWithTag(map[string]string{
		"a":    "123",
		"B":    "123",
		"cccc": "23.2",
		"D":    "12",
		"zzz":  "asfasd",
		"F":    "1",
	}, &tg, "redis")
	if err != nil {
		t.Fatal(err)
	}
	if tg.A != 123 || tg.B != 123 || tg.E != "asfasd" || !tg.F {
		t.Errorf("wrong struct value: %+v", tg)
	}
	ScanStruct(map[string]string{"bb": "90234"}, &tg)
	if tg.B != 90234 {
		t.Errorf("default tag cast failed")
	}
}
