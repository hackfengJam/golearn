package model

import (
	"fmt"
	"testing"
)

func TestCreateRole(t *testing.T) {

	role := &Role{
		Name: "zhangsan",
	}
	if id, err := CreateRole(role); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}
	return
}

func TestGetRole(t *testing.T) {
	if role, err := GetRole(1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(role)
	}
	return
}
