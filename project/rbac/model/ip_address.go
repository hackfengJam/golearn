package model

import (
	"fmt"
	"golearn/project/rbac/utils/sqlhelper"
)

const defaultMask string = "255.255.255.255"

// Requests
type SetIpAddressArgs struct {
	Id      uint64 `json:"id" binding:"omitempty,gte=0"`
	Name    string `json:"name" binding:"required,lte=64"`       // 角色名称
	Ip      string `json:"ip"`                                   // ip: ipv4
	Mask    uint32 `json:"mask" binding:"required,gte=0,lte=32"` // 掩码
	IsAdmin bool   `json:"is_admin"`                             // 是否是 admin 级
}

func CreateIpAddress(ipAddress *IpAddress) (rid uint64, err error) {
	var (
		id int64
	)

	tableName := "ip_address"
	insertIpAddressColumns := "name, ip, mask, is_admin"
	if ipAddress.Mask == "" {
		ipAddress.Mask = defaultMask
	}

	insertIpAddressStmt, err := db.PrepareNamed(fmt.Sprintf("INSERT INTO %s(%s) VALUES(:name,:ip,:mask,:is_admin)", tableName, insertIpAddressColumns))
	if err != nil {
		fmt.Println(err)
		return
	}

	rev, err := insertIpAddressStmt.Exec(ipAddress)
	if err != nil {
		if sqlhelper.IsDup(err) {
			fmt.Println(err)
			return
		}
		fmt.Println(err)
		return
	}

	if id, err = rev.LastInsertId(); err != nil {
		fmt.Println(err)
		return
	} else {
		rid = uint64(id)
	}
	return
}

func UpdateIpAddress(ipAddress *IpAddress) (err error) {
	tableName := "ip_address"

	updateIpAddressStmt, err := db.Preparex(fmt.Sprintf(
		"UPDATE %s "+
			"SET "+
			"name=:name,ip=:ip,mask=:mask,is_admin=:is_admin "+
			"where id=?", tableName))

	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = updateIpAddressStmt.Exec(ipAddress)
	if err != nil {
		fmt.Println(err)
		return
	}
	return err
}

func GetIpAddress(id uint64) (ipAddress IpAddress, err error) {
	tableName := "ip_address"
	columns := "id, name, ip, mask, is_admin, status, updated_time, created_time"

	getIpAddressStmt, err := db.Preparex(fmt.Sprintf(
		"SELECT %s from %s "+
			" where "+
			" id = ?;", columns, tableName))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = getIpAddressStmt.Get(&ipAddress, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func ListIpAddress() (ipAddressList []*IpAddress, err error) {
	tableName := "ip_address"
	columns := "id, name, ip, mask, is_admin, updated_time, created_time"
	pageMarker := 0
	pageSize := 1024

	listIpAddressStmt, err := db.Preparex(fmt.Sprintf(
		"SELECT %s from %s "+
			" where "+
			" id > ? "+
			" AND status=1 "+
			" order by id asc limit ?", columns, tableName))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = listIpAddressStmt.Select(&ipAddressList, pageMarker, pageSize)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func ListIpAddressByIp(ip string) (ipAddressList []*IpAddress, err error) {
	tableName := "ip_address"
	columns := "id, name, ip, mask, is_admin, updated_time, created_time"
	pageSize := 1024

	listIpAddressByIpStmt, err := db.Preparex(fmt.Sprintf(
		"SELECT %s from %s as ip_a"+
			" where "+
			" inet_ntoa(inet_aton(?) & inet_aton(ip_a.mask))=ip_a.ip "+
			" AND ip_a.status=1 "+
			" order by ip_a.id asc limit ?;", columns, tableName))
	if err != nil {
		fmt.Println(err)
		return
	}

	err = listIpAddressByIpStmt.Select(&ipAddressList, ip, pageSize)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func DeleteIpAddress(id uint64) (err error) {
	tableName := "ip_address"

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
