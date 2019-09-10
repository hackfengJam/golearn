package model

import (
	"fmt"
	"golearn/project/rbac/utils/sqlhelper"
)

// Requests
type SetAkConfigArgs struct {
	Id        uint64 `json:"id" binding:"omitempty,gte=0"`
	Namespace string `json:"name" binding:"required,lte=64"`  // 命名空间
	Name      string `json:"title" binding:"required,lte=64"` // name 唯一
	Title     string `json:"remark"`                          // 标题
	Remark    string `json:"value"`                           // 备注
}

func CreateAkConfig(akConfig *AkConfig) (rid uint64, err error) {
	var (
		id int64
	)

	tableName := "ak_config"
	insertAkConfigColumns := "namespace, name, title, remark, value"

	insertAkConfigStmt, err := db.PrepareNamed(fmt.Sprintf("INSERT INTO %s(%s) VALUES(:namespace,:name,:title,:remark,:value)", tableName, insertAkConfigColumns))
	if err != nil {
		return
	}

	rev, err := insertAkConfigStmt.Exec(akConfig)
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

func UpdateAkConfig(akConfig *AkConfig) error {
	tableName := "ak_config"

	updateAkConfigStmt, err := db.Preparex(fmt.Sprintf("UPDATE %s SET name=? where id=?", tableName))
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = updateAkConfigStmt.Exec(akConfig.Name, akConfig.Id)
	return err
}

func GetAkConfig(id uint64) (akConfig AkConfig, err error) {
	tableName := "ak_config"
	columns := "id, name, status, updated_time, created_time"

	getAkConfigStmt, err := db.Preparex(fmt.Sprintf(
		"SELECT %s from %s "+
			" where "+
			" id = ?;", columns, tableName))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = getAkConfigStmt.Get(&akConfig, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func ListAkConfig() (akConfigList []*AkConfig, err error) {
	tableName := "ak_config"
	columns := "id, name, status, updated_time, created_time"
	pageMarker := 0
	pageSize := 1024

	listAkConfigStmt, err := db.Preparex(fmt.Sprintf("select %s from %s where id > ? order by id asc limit ?", columns, tableName))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = listAkConfigStmt.Select(&akConfigList, pageMarker, pageSize)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
