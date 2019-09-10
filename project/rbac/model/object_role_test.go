package model

import (
	"fmt"
	"testing"
)

func TestCreateObjectRole(t *testing.T) {
	objectRole := &ObjectRole{
		ObjectType: 1,
		ObjectId:   1,
		RoleId:     1,
	}

	if id, err := CreateObjectRole(objectRole); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}
	return
}

func TestListObjectRole(t *testing.T) {
	listObjectRoleArgs := &ListObjectRoleArgs{
		CreateObjectRoleArgs{
			ObjectType: 1,
			ObjectId:   1,
			RoleId:     1,
		}}
	if objectRoleList, err := ListObjectRole(listObjectRoleArgs); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(objectRoleList)
	}
	return
}
