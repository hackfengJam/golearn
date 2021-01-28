package model

import (
	"fmt"
	"golearn/project/rbac/utils/sqlhelper"
	"strings"
)

// Requests
type CreateRolePermissionArgs struct {
	RoleId       uint64 `json:"role_id" binding:"required,gte=0"`       // 角色id
	PermissionId uint64 `json:"permission_id" binding:"required,gte=0"` // 对象id
}

// Requests
type ListRolePermissionArgs struct {
	CreateRolePermissionArgs
}

func CreateRolePermission(rolePermission *RolePermission) (rid uint64, err error) {
	var (
		id int64
	)

	tableName := "role_permission"
	insertRolePermissionColumns := "role_id, permission_id"

	insertRolePermissionStmt, err := db.PrepareNamed(fmt.Sprintf(
		"INSERT INTO "+
			"%s(%s) "+
			"VALUES(:role_id,:permission_id)", tableName, insertRolePermissionColumns))
	if err != nil {
		return
	}

	rev, err := insertRolePermissionStmt.Exec(rolePermission)
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

func ListRolePermission(listRolePermissionArgs *ListRolePermissionArgs) (rolePermissionList []*RolePermission, err error) {
	tableName := "role_permission"
	columns := "id, role_id, permission_id, created_time"
	pageSize := 1024

	// 拼接 sql
	var sqlHandler = &SqlHandler{}

	sqlHandler.Segment = append(sqlHandler.Segment, fmt.Sprintf(
		`
		SELECT %s FROM %s 
		WHERE
		1=1 `, columns, tableName))

	if listRolePermissionArgs == nil {
		listRolePermissionArgs = &ListRolePermissionArgs{}
	}
	if listRolePermissionArgs.RoleId != 0 {
		sqlHandler.Segment = append(sqlHandler.Segment, " AND role_id = ? ")
		sqlHandler.Param = append(sqlHandler.Param, listRolePermissionArgs.RoleId)
	}
	if listRolePermissionArgs.PermissionId != 0 {
		sqlHandler.Segment = append(sqlHandler.Segment, " AND permission_id = ? ")
		sqlHandler.Param = append(sqlHandler.Param, listRolePermissionArgs.PermissionId)
	}

	sqlHandler.Segment = append(sqlHandler.Segment, " order by id asc limit ?;")
	sqlHandler.Param = append(sqlHandler.Param, pageSize)

	listRolePermissionStmt, err := db.Preparex(strings.Join(sqlHandler.Segment, " "))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = listRolePermissionStmt.Select(&rolePermissionList, sqlHandler.Param...)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func DeleteRolePermission(id uint64) (err error) {
	tableName := "role_permission"

	deleteIpAddressStmt, err := db.Preparex(fmt.Sprintf(
		"DELETE FROM %s "+
			"where id=?", tableName))

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = deleteIpAddressStmt.Exec(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
