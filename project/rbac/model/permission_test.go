package model

import (
	"fmt"
	"testing"
)

func TestCreatePermission(t *testing.T) {

	permission := &Permission{
		Title:     "zhangsan",
		EntityKey: "key-shijiu",
	}
	if id, err := CreatePermission(permission); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}
	return
}

func TestGetPermission(t *testing.T) {
	if role, err := GetPermission(1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(role)
	}
	return
}
