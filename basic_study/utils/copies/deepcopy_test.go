package copies

import (
	"fmt"
	"reflect"
	"testing"
)

type subType struct {
	Name string
	Age  int
}

type dataType struct {
	F9  subType
	F2  string
	F3  []int
	F4  map[string]string
	F6  []byte
	F10 int64
	F15 uint64
	F17 float64
	F11 uint
	F1  int
	F16 float32
	F14 uint32
	F5  int32
	F13 uint16
	F8  int16
	F12 uint8
	F7  int8
	F18 bool
}

func TestDeepCopy(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		var d1 int = 234

		var d2 int

		DeepCopy(d1, &d2)

		if d2 != 234 {
			t.Error(d2)
		}
	})

	t.Run("uint16", func(t *testing.T) {
		var d1 uint16 = 1999

		var d2 uint16

		DeepCopy(d1, &d2)

		if d2 != 1999 {
			t.Error(d2)
		}
	})

	t.Run("int-slice", func(t *testing.T) {
		d1 := []int{4, 5, 6}

		var d2 []int

		DeepCopy(d1, &d2)

		if v := fmt.Sprintf("%v", d2); v != "[4 5 6]" {
			t.Error(v)
		}
	})

	t.Run("int-slice2", func(t *testing.T) {
		d1 := [][]int{{1, 2, 34}, {5, 6, 7}}

		var d2 [][]int

		DeepCopy(d1, &d2)
		if v := fmt.Sprintf("%v", d2); v != "[[1 2 34] [5 6 7]]" {
			t.Error(v)
		}

		d1[1][1] = 2345

		if v := fmt.Sprintf("%v", d1); v != "[[1 2 34] [5 2345 7]]" {
			t.Error(v)
		}

		if v := fmt.Sprintf("%v", d2); v != "[[1 2 34] [5 6 7]]" {
			t.Error(v)
		}
	})

	t.Run("map", func(t *testing.T) {
		d1 := map[int]uint32{
			34:  234,
			-12: 2356,
			567: 1299,
		}

		d2 := map[int]uint32{}

		DeepCopy(d1, &d2)

		if d2[-12] != 2356 || d2[34] != 234 || d2[567] != 1299 {
			t.Error(d2)
		}
	})

	t.Run("map2", func(t *testing.T) {
		d1 := map[string][]int{
			"hello": {4, 5, 6},
			"ss":    {},
		}

		d2 := map[string][]int{"hello": {4}, "ss": nil}

		DeepCopy(d1, &d2)

		if d2["hello"][1] != 5 || len(d2["ss"]) != 0 {
			t.Error(d2)
		}
	})

	t.Run("struct", func(t *testing.T) {
		d1 := dataType{
			F1:  123,
			F3:  []int{4, 5},
			F4:  map[string]string{"hello": "1"},
			F2:  "say",
			F6:  []byte("no"),
			F5:  823,
			F7:  8,
			F8:  -123,
			F9:  subType{"pp", 84},
			F10: 88323423,
			F11: 234,
			F12: 12,
			F13: 1234,
			F14: 8834,
			F15: 12359884,
			F16: 4,
			F17: 9,
			F18: true,
		}

		f3 := make([]int, 2)

		var d2 dataType

		d2.F3 = f3

		DeepCopy(d1, &d2)

		if !reflect.DeepEqual(d1, d2) {
			t.Error(d2)
		}

		f3[1] = 901

		if d2.F3[1] != 901 {
			t.Error(d2.F3)
		}

		d1.F3[0] = 40

		if d2.F3[0] != 4 {
			t.Error(d2.F3)
		}

		d3 := map[string]dataType{"dd": d1}

		var d4 map[string]dataType

		DeepCopy(d3, &d4)

		if !reflect.DeepEqual(d3, d4) {
			t.Error(d4)
		}

		d3["dd"].F4["pp"] = "qq"

		if _, ok := d4["dd"].F4["pp"]; ok {
			t.Error(d4)
		}
	})

	t.Run("interface", func(t *testing.T) {
		a := map[string]interface{}{
			"hello": "world",
			"agent": 901,
			"foo":   []int{4, 5, 67},
			"bar":   map[string]string{"aa": "11"},
		}

		var b map[string]interface{}

		DeepCopy(a, &b)
		if !reflect.DeepEqual(a, b) {
			t.Error(b)
		}

		var m int = 234
		var n int
		var p interface{} = n

		DeepCopy(m, &p)
		if p != 234 {
			t.Error(p)
		}
	})

	t.Run("ptr", func(t *testing.T) {
		type ptrData struct {
			If  interface{}
			If2 interface{}
			Mf  map[string]string
			Ff  *int
			Fg  string
		}

		ff := 344

		a := ptrData{
			If2: []byte("yy"),
			Ff:  &ff,
			Fg:  "hello",
		}

		var p int

		b := ptrData{Ff: &p}

		DeepCopy(a, &b)
		if !reflect.DeepEqual(a, b) {
			t.Error(b)
		}

		p = 90

		if *b.Ff != 90 {
			t.Error(*b.Ff)
		}

		a = ptrData{Fg: "j"}

		DeepCopy(a, &b)
		if !reflect.DeepEqual(a, b) {
			t.Error(b)
		}
	})

	t.Run("nil", func(t *testing.T) {
		v := dataType{F9: subType{"aa", 90}, F4: map[string]string{"hello": "i"}}

		var a interface{}

		DeepCopy(v, &a)
		if !reflect.DeepEqual(v, a) {
			t.Error(a)
		}
	})

	t.Run("pointer", func(t *testing.T) {
		a := 90

		var b *int = &a

		var c interface{}

		DeepCopy(b, &c)
		if c != 90 {
			t.Error(c)
		}
	})
}
