package model

import (
	"fmt"
	"golearn/project/rbac/utils/sqlhelper"
	"strings"
)

// Requests
type CreateObjectRoleArgs struct {
	ObjectType int    `json:"object_type" binding:"omitempty,gt=0"`
	ObjectId   uint64 `json:"object_id" binding:"required,gte=0"` // 对象id
	RoleId     uint64 `json:"role_id" binding:"required,gte=0"`   // 角色id
}

// Requests
type ListObjectRoleArgs struct {
	CreateObjectRoleArgs
}

func CreateObjectRole(objectRole *ObjectRole) (rid uint64, err error) {
	var (
		id int64
	)

	tableName := "object_role"
	insertObjectRoleColumns := "object_type, object_id, role_id"

	insertObjectRoleStmt, err := db.PrepareNamed(fmt.Sprintf(
		"INSERT INTO "+
			"%s(%s) "+
			"VALUES(:object_type,:object_id,:role_id)", tableName, insertObjectRoleColumns))
	if err != nil {
		return
	}

	rev, err := insertObjectRoleStmt.Exec(objectRole)
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

func ListObjectRole(listObjectRoleArgs *ListObjectRoleArgs) (objectRoleList []*ObjectRole, err error) {
	tableName := "object_role"
	columns := "id, object_type, object_id, role_id, created_time"
	pageSize := 1024

	// 拼接 sql
	var sqlHandler = &SqlHandler{}

	sqlHandler.Segment = append(sqlHandler.Segment, fmt.Sprintf(
		`
		SELECT %s FROM %s 
		WHERE
		1=1 `, columns, tableName))

	if listObjectRoleArgs == nil {
		listObjectRoleArgs = &ListObjectRoleArgs{}
	}
	if listObjectRoleArgs.ObjectType != 0 {
		sqlHandler.Segment = append(sqlHandler.Segment, " AND object_type = ? ")
		sqlHandler.Param = append(sqlHandler.Param, listObjectRoleArgs.ObjectType)
	}
	if listObjectRoleArgs.ObjectId != 0 {
		sqlHandler.Segment = append(sqlHandler.Segment, " AND object_id = ? ")
		sqlHandler.Param = append(sqlHandler.Param, listObjectRoleArgs.ObjectId)
	}
	if listObjectRoleArgs.RoleId != 0 {
		sqlHandler.Segment = append(sqlHandler.Segment, " AND role_id = ? ")
		sqlHandler.Param = append(sqlHandler.Param, listObjectRoleArgs.RoleId)
	}

	sqlHandler.Segment = append(sqlHandler.Segment, " order by id asc limit ?;")
	sqlHandler.Param = append(sqlHandler.Param, pageSize)

	listObjectRoleStmt, err := db.Preparex(strings.Join(sqlHandler.Segment, " "))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = listObjectRoleStmt.Select(&objectRoleList, sqlHandler.Param...)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func DeleteObjectRole(id uint64) (err error) {
	tableName := "object_role"

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
