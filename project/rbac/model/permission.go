package model

import (
	"errors"
	"fmt"
	"golearn/project/rbac/utils/sqlhelper"
	"strings"
)

// Requests
type SetPermissionArgs struct {
	Id        uint64 `json:"id" binding:"omitempty,gte=0"`
	Title     string `json:"title" binding:"required,lte=64"`       // 权限名称
	EntityKey string `json:"entity_key" binding:"required,lte=128"` // 可访问的 entity_key
}

func CreatePermission(permission *Permission) (rid uint64, err error) {
	var (
		id int64
	)

	tableName := "permission"
	insertPermissionColumns := "title, entity_key"

	insertPermissionStmt, err := db.PrepareNamed(fmt.Sprintf("INSERT INTO %s(%s) VALUES(:title,:entity_key)", tableName, insertPermissionColumns))
	if err != nil {
		return
	}

	rev, err := insertPermissionStmt.Exec(permission)
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

func UpdatePermission(permission *Permission) error {
	tableName := "permission"

	if permission == nil {
		permission = &Permission{}
	}

	// 拼接 sql
	var sqlHandler = &SqlHandler{}
	sqlHandler.Segment = append(sqlHandler.Segment, fmt.Sprintf(
		`
		UPDATE %s SET
		`, tableName))

	// update set
	var updateSegment []string
	if permission.Title != "" {
		updateSegment = append(updateSegment, " title = ? ")
		sqlHandler.Param = append(sqlHandler.Param, permission.Title)
	}
	if permission.EntityKey != "" {
		updateSegment = append(updateSegment, " entity_key = ? ")
		sqlHandler.Param = append(sqlHandler.Param, permission.EntityKey)
	}
	if len(updateSegment) == 0 {
		return errors.New("BadRequest")
	}
	sqlHandler.Segment = append(sqlHandler.Segment, strings.Join(updateSegment, ", "))

	// where
	sqlHandler.Segment = append(sqlHandler.Segment, " where id=? ")
	sqlHandler.Param = append(sqlHandler.Param, permission.Id)

	updatePermissionStmt, err := db.Preparex(strings.Join(sqlHandler.Segment, " "))
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = updatePermissionStmt.Exec(sqlHandler.Param...)
	return err
}

func GetPermission(id uint64) (permission Permission, err error) {
	tableName := "permission"
	columns := "id, title, entity_key, status, updated_time, created_time"

	getPermissionStmt, err := db.Preparex(fmt.Sprintf(
		"SELECT %s from %s "+
			" where "+
			" id = ?;", columns, tableName))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = getPermissionStmt.Get(&permission, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func ListPermission() (permissionList []*Permission, err error) {
	tableName := "permission"
	columns := "id, title, entity_key, status, updated_time, created_time"
	pageMarker := 0
	pageSize := 1024

	listPermissionStmt, err := db.Preparex(fmt.Sprintf("select %s from %s where id > ? AND status=1 order by id asc limit ?", columns, tableName))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = listPermissionStmt.Select(&permissionList, pageMarker, pageSize)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func DeletePermission(id uint64) (err error) {
	tableName := "permission"

	deleteIpAddressStmt, err := db.Preparex(fmt.Sprintf(
		"UPDATE %s "+
			"SET "+
			"status=0 "+
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
