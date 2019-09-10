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

func CreateRole(role *Role) (rid uint64, err error) {
	var (
		id int64
	)

	tableName := "role"
	insertRoleColumns := "name"

	insertRoleStmt, err := db.PrepareNamed(fmt.Sprintf("INSERT INTO %s(%s) VALUES(:name)", tableName, insertRoleColumns))
	if err != nil {
		return
	}

	rev, err := insertRoleStmt.Exec(role)
	if err != nil {
		if sqlhelper.IsDup(err) {
			return
		}
		return
	}
	if id, err = rev.LastInsertId(); err != nil {
		return
	} else {
		rid = uint64(id)
	}
	return
}

func UpdateRole(role *Role) error {
	tableName := "role"

	updateRoleStmt, err := db.Preparex(fmt.Sprintf("UPDATE %s SET name=? where id=?", tableName))
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = updateRoleStmt.Exec(role.Name, role.Id)
	return err
}

func GetRole(id uint64) (role Role, err error) {
	tableName := "role"
	columns := "id, name, status, updated_time, created_time"

	getRoleStmt, err := db.Preparex(fmt.Sprintf(
		"SELECT %s from %s "+
			" where "+
			" id = ?;", columns, tableName))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = getRoleStmt.Get(&role, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func ListRole() (roleList []*Role, err error) {
	tableName := "role"
	columns := "id, name, status, updated_time, created_time"
	pageMarker := 0
	pageSize := 1024

	listRoleStmt, err := db.Preparex(fmt.Sprintf("select %s from %s where id > ? order by id asc limit ?", columns, tableName))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = listRoleStmt.Select(&roleList, pageMarker, pageSize)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
