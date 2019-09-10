package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golearn/project/rbac/model"
	services "golearn/project/rbac/utils/bizresp"
	"golearn/project/rbac/utils/sqlhelper"
	"strconv"
)

/*
 * 添加或者编辑 ip 地址页面
 * post 处理添加或者编辑动作
 */
func SetIpAddress(c *gin.Context) {
	req := &model.SetIpAddressArgs{}
	err := BindJSON(c, req)
	if err != nil {
		fail(c, services.ErrRequestInvalid, err)
		return
	}
	var ipAddress = &model.IpAddress{
		Id:      req.Id,
		Name:    req.Name,
		Ip:      req.Ip,
		IsAdmin: req.IsAdmin,
	}
	ipAddress.Mask = sqlhelper.MaskIntToString(req.Mask)

	rev := model.IpAddress{}

	if req.Id == 0 {
		// insert
		id, err := model.CreateIpAddress(ipAddress)
		if err != nil {
			fmt.Printf("args:%v, err:%v\n", req, err)
			fail(c, services.ErrRequestInvalid, err)
			return
		}
		rev.Id = id
	} else {
		// update
		err := model.UpdateIpAddress(ipAddress)
		if err != nil {
			fmt.Printf("args:%v, err:%v\n", req, err)
			fail(c, services.ErrRequestInvalid, err)
			return
		}
		rev.Id = req.Id
	}
	success(c, rev)
	return
}

/*
 * ip 列表页
 */
func ListIpAddress(c *gin.Context) {
	// list by ip

	// TODO from header
	ip := c.Query("ip")

	var (
		ipAddressList []*model.IpAddress
		err           error
	)
	if ip != "" {
		ipAddressList, err = model.ListIpAddressByIp(ip)
		if err != nil {
			fmt.Printf("ListIpAddressByIp err:%v\n", err)
			fail(c, services.ErrRequestInvalid, err)
			return
		}
	} else {
		ipAddressList, err = model.ListIpAddress()
		if err != nil {
			fmt.Printf("ListIpAddress err:%v\n", err)
			fail(c, services.ErrRequestInvalid, err)
			return
		}
	}

	success(c, ipAddressList)
	return
}

/*
 * DeleteIpAddress
 */
func DeleteIpAddress(c *gin.Context) {
	// delete by id
	id := c.Param("id")
	if id == "" {
		fail(c, services.ErrRequestInvalid, nil)
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fail(c, services.ErrRequestInvalid, errors.New("invalid ip_address_id"))
		return
	}

	err = model.DeleteIpAddress(idUint)
	if err != nil {
		fmt.Printf("DeleteIpAddress err:%v\n", err)
		fail(c, services.ErrRequestInvalid, err)
		return
	}

	success(c, gin.H{})
	return
}
