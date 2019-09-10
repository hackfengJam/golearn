package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"golearn/project/rbac/model"
	services "golearn/project/rbac/utils/bizresp"
	"strconv"
)

/*
 * 添加或者编辑 ObjectRole
 * post 处理添加或者编辑动作
 */
func CreateObjectRole(c *gin.Context) {
	req := &model.CreateObjectRoleArgs{}
	err := BindJSON(c, req)
	if err != nil {
		fail(c, services.ErrRequestInvalid, err)
		return
	}
	var objectRole = &model.ObjectRole{
		ObjectType: req.ObjectType,
		ObjectId:   req.ObjectId,
		RoleId:     req.RoleId,
	}

	// 默认 1: Ip Whitelisting
	if objectRole.ObjectType <= 0 {
		objectRole.ObjectType = 1
	}

	// 验证 ip address id
	if objectRole.ObjectType == 1 {
		if _, err := model.GetIpAddress(objectRole.ObjectId); err != nil {
			fmt.Printf("args:%v, err:%v\n", req, err)
			fail(c, services.ErrIpAddressNotFound, err)
			return
		}
	}

	// 验证 role id
	if _, err := model.GetRole(objectRole.RoleId); err != nil {
		fmt.Printf("args:%v, err:%v\n", req, err)
		fail(c, services.ErrRoleNotFound, err)
		return
	}

	rev := model.ObjectRole{}

	// insert
	id, err := model.CreateObjectRole(objectRole)
	if err != nil {
		fmt.Printf("args:%v, err:%v\n", req, err)
		fail(c, services.ErrRequestInvalid, err)
		return
	}
	rev.Id = id

	success(c, rev)
	return
}

/*
 * ObjectRole 列表页
 */
func ListObjectRole(c *gin.Context) {
	param := map[string]string{
		"object_type": c.Query("object_type"),
		"object_id":   c.Query("object_id"),
		"role_id":     c.Query("role_id"),
	}

	var (
		objectRoleList []*model.ObjectRole
		err            error
	)

	req := &model.ListObjectRoleArgs{}
	// object_type
	if q, ok := param["object_type"]; ok && q != "" {
		val, err := strconv.ParseInt(q, 10, 64)
		if err != nil {
			fail(c, services.ErrRequestInvalid, errors.New("invalid object_type"))
			return
		}
		req.ObjectType = int(val)
	}
	// object_id
	if q, ok := param["object_id"]; ok && q != "" {
		val, err := strconv.ParseUint(q, 10, 64)
		if err != nil {
			fail(c, services.ErrRequestInvalid, errors.New("invalid object_id"))
			return
		}
		req.ObjectId = val
	}
	// role_id
	if q, ok := param["role_id"]; ok && q != "" {
		val, err := strconv.ParseUint(q, 10, 64)
		if err != nil {
			fail(c, services.ErrRequestInvalid, errors.New("invalid role_id"))
			return
		}
		req.RoleId = val
	}

	objectRoleList, err = model.ListObjectRole(req)
	if err != nil {
		fmt.Printf("ListObjectRole err:%v\n", err)
		fail(c, services.ErrRequestInvalid, err)
		return
	}

	success(c, objectRoleList)
	return
}

/*
 * 删除 ObjectRole
 */
func DeleteObjectRole(c *gin.Context) {
	// delete by id
	id := c.Param("id")
	if id == "" {
		fail(c, services.ErrRequestInvalid, nil)
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fail(c, services.ErrRequestInvalid, errors.New("invalid object_role"))
		return
	}

	err = model.DeleteObjectRole(idUint)
	if err != nil {
		fmt.Printf("DeleteObjectRole err:%v\n", err)
		fail(c, services.ErrRequestInvalid, err)
		return
	}

	success(c, gin.H{})
	return
}
