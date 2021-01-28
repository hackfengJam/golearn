package model

import (
	"fmt"
	"golearn/project/rbac/utils/sqlhelper"
	"testing"
)

func TestCreateIpAddress(t *testing.T) {

	var mask uint32 = 8

	ipAddress := &IpAddress{
		Name: "shijiu-pc",
		Ip:   "10.105.16.214",
		//Mask:    "",
		IsAdmin: true,
		Status:  false,
	}
	// mask int to string  16 -> 255.255.0.0
	ipAddress.Mask = sqlhelper.MaskIntToString(mask)

	if id, err := CreateIpAddress(ipAddress); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(id)
	}
	return
}

func TestGetIpAddress(t *testing.T) {

	if ipAddress, err := GetIpAddress(1); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ipAddress)
	}
	return
}

func TestListIpAddress(t *testing.T) {

	if ipAddressList, err := ListIpAddress(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ipAddressList)
	}
	return
}

func TestListIpAddressByIp(t *testing.T) {

	if ipAddressList, err := ListIpAddressByIp("10.105.16.214"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ipAddressList)
	}
	return
}
