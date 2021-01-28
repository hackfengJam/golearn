package model

import (
	"fmt"
	"testing"
)

func TestCreateRolePermission(t *testing.T) {
	rolePermission := &RolePermission{
		PermissionId: 1,
		RoleId:       1,
	}

	if id, err := CreateRolePermission(rolePermission); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}
	return
}

func TestListRolePermission(t *testing.T) {
	listRolePermissionArgs := &ListRolePermissionArgs{
		CreateRolePermissionArgs{
			RoleId:       1,
			PermissionId: 1,
		}}
	if objectRoleList, err := ListRolePermission(listRolePermissionArgs); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(objectRoleList)
	}
	return
}
