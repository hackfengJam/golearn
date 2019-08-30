package model

import (
	"fmt"
	"golearn/project/rbac/utils/sqlhelper"
)

// Requests
type SetRoleArgs struct {
	Id   uint64 `json:"id" binding:"omitempty,gte=0"`
	Name string `json:"name" binding:"required,lte=64"` // 角色名称
}

func CreateRole(role *Role) error {
	tableName := "role"
	insertRoleColumns := "name"

	insertRoleStmt, err := db.PrepareNamed(fmt.Sprintf("INSERT INTO %s(%s) VALUES(:name)", tableName, insertRoleColumns))
	if err != nil {
		return err
	}

	_, err = insertRoleStmt.Exec(role)
	if err != nil {
		if sqlhelper.IsDup(err) {
			return nil
		}
		return err
	}
	return err
}

func UpdateRole(role *Role) error {
	tableName := "role"

	updateRoleStmt, err := db.Preparex(fmt.Sprintf("UPDATE %s SET name=? where id=?", tableName))
	if err != nil {
		return err
	}

	_, err = updateRoleStmt.Exec(role.Name, role.Id)
	return err
}
