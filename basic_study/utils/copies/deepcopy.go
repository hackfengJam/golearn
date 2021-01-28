// Package copies implements utilities about data copy.
package copies

import (
	"reflect"
)

// DeepCopy deeply copies src to dest.
// If type mismatch, that field will be ignored.
//
// The unexported fields are ignored.
//
// Example:
//  type dataType struct {
//    F1 int
//    F2 string
//    F3 []int
//    F4 map[string]string
//  }
//
//  d1 := dataType{F3: []int{4, 5}, F4: map[string]string{"hello": "1"}, F2: "say"}
//
//  var v dataType
//  copies.DeepCopy(d, &v)
func DeepCopy(src interface{}, dest interface{}) {
	deepCopy(reflect.ValueOf(src), reflect.ValueOf(dest))
}

func isNil(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice:
		if v.IsNil() {
			return true
		}
	}

	return false
}

func deepCopy(src reflect.Value, dst reflect.Value) {
	// src is nil.
	if isNil(src) {
		dst.Set(src)
		return
	}

	v := reflect.Indirect(src)

	// dst is nil.
	if dst.Kind() == reflect.Ptr {
		if dst.IsNil() {
			dst.Set(reflect.New(v.Type()))
		}

		dst = dst.Elem()
	}

	if dst.Kind() == reflect.Interface && v.Kind() != reflect.Interface {
		e := reflect.New(v.Type())
		deepCopy(src, e)
		dst.Set(e.Elem())

		return
	}

	d := reflect.Indirect(dst)

	if v.Kind() != d.Kind() || v.Type() != d.Type() {
		return
	}

	switch v.Kind() {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Float32, reflect.Float64, reflect.Bool, reflect.String:
		d.Set(v)
	case reflect.Slice, reflect.Array:
		vt := v.Type().Elem()
		dt := d.Type().Elem()

		if vt != dt {
			return
		}

		if v.Len() != d.Len() {
			d.Set(reflect.MakeSlice(v.Type(), v.Len(), v.Cap()))
		}

		for i := 0; i < v.Len(); i++ {
			deepCopy(v.Index(i), d.Index(i))
		}
	case reflect.Map:
		vkt := v.Type().Key()
		vvt := v.Type().Elem()

		dkt := d.Type().Key()
		dvt := d.Type().Elem()

		if vkt != dkt || vvt != dvt {
			return
		}

		d.Set(reflect.MakeMapWithSize(v.Type(), v.Len()))

		vIter := v.MapRange()

		for vIter.Next() {
			value := vIter.Value()

			n := reflect.New(value.Type())

			deepCopy(value, n)
			d.SetMapIndex(vIter.Key(), n.Elem())
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			f := v.Field(i)
			t := v.Type().Field(i)

			n := d.FieldByName(t.Name)

			if n.CanSet() {
				deepCopy(f, n)
			}
		}
	case reflect.Interface:
		deepCopy(v.Elem(), d)
	}
}
