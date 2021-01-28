package cast

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// ScanStruct is a specialized form of ScanStructWithTag with struct field tag name`cast`.
func ScanStruct(m map[string]string, dst interface{}) error {
	return ScanStructWithTag(m, dst, "cast")
}

// ScanStructWithTag scans a string map and cast to a struct.
// The supported types:
//  int int8 int16 int32 int64
//  uint uint8 uint16 uint32 uint64
//  float32 float64
//  bool
//  string
func ScanStructWithTag(m map[string]string, dst interface{}, tagClass string) error {
	getValue := reflect.ValueOf(dst)
	if getValue.Kind() != reflect.Ptr || getValue.IsNil() {
		return errors.New("cast: dst should be pointer")
	}
	setValue := getValue.Elem()
	if !setValue.CanSet() {
		return errors.New("cast: dst can not set")
	}
	if setValue.Kind() != reflect.Struct {
		return errors.New("cast: dst should be a struct")
	}
	getType := setValue.Type()

	for idx := 0; idx < getType.NumField(); idx++ {
		field := getType.Field(idx)
		name := field.Name
		if tag, ok := field.Tag.Lookup(tagClass); ok {
			if tag == "-" {
				continue
			}
			// Use tag name
			name = tag
		}
		if v, ok := m[name]; ok {
			switch field.Type.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				i, err := strconv.ParseInt(v, 10, field.Type.Bits())
				if err != nil {
					return fmt.Errorf("cast: scan field %v type %v %v", name, field.Type, err)
				}
				setValue.Field(idx).SetInt(i)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				i, err := strconv.ParseUint(v, 10, field.Type.Bits())
				if err != nil {
					return fmt.Errorf("cast: scan field %v type %v %v", name, field.Type, err)
				}
				setValue.Field(idx).SetUint(i)
			case reflect.String:
				setValue.Field(idx).SetString(v)
			case reflect.Float32, reflect.Float64:
				fval, err := strconv.ParseFloat(v, field.Type.Bits())
				if err != nil {
					return fmt.Errorf("cast: scan field %v type %v %v", name, field.Type, err)
				}
				setValue.Field(idx).SetFloat(fval)
			case reflect.Bool:
				setValue.Field(idx).SetBool(v == "1")
			default:
				return fmt.Errorf("cast: unsupport Field %v Type %v", name, field.Type.Kind())
			}
		}
	}
	return nil
}

// ToMap is a specialized form of ToMapWithTag with the fixed tag name `cast`.
func ToMap(src interface{}) (map[string]interface{}, error) {
	return ToMapWithTag(src, "cast")
}

// ToMapWithTag casts a struct to a map.
func ToMapWithTag(src interface{}, tagClass string) (map[string]interface{}, error) {
	getValue := reflect.Indirect(reflect.ValueOf(src))
	if !getValue.IsValid() {
		return nil, errors.New("cast: src is nil")
	}
	if getValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf("cast: src type is not struct %v", getValue.Type())
	}
	getType := getValue.Type()
	m := make(map[string]interface{})
	for i := 0; i < getValue.NumField(); i++ {
		if val, ok := getType.Field(i).Tag.Lookup(tagClass); ok {
			// use tag first
			if val == "-" {
				// ignore
				continue
			}
			m[val] = getValue.Field(i).Interface()
		} else {
			m[getType.Field(i).Name] = getValue.Field(i).Interface()
		}
	}
	return m, nil
}

// MapToSlice casts a map to slice
func MapToSlice(src map[string]interface{}) []interface{} {
	args := make([]interface{}, len(src)*2)
	idx := 0
	for k, v := range src {
		args[idx] = k
		args[idx+1] = v
		idx += 2
	}
	return args
}
