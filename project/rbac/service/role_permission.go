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
 * 添加或者编辑 RolePermission
 * post 处理添加或者编辑动作
 */
func CreateRolePermission(c *gin.Context) {
	req := &model.CreateRolePermissionArgs{}
	err := BindJSON(c, req)
	if err != nil {
		fail(c, services.ErrRequestInvalid, err)
		return
	}
	var rolePermission = &model.RolePermission{
		RoleId:       req.RoleId,
		PermissionId: req.PermissionId,
	}

	// 验证 role id
	if _, err := model.GetRole(rolePermission.RoleId); err != nil {
		fmt.Printf("args:%v, err:%v\n", req, err)
		fail(c, services.ErrRoleNotFound, err)
		return
	}
	// 验证 permission id
	if _, err := model.GetPermission(rolePermission.PermissionId); err != nil {
		fmt.Printf("args:%v, err:%v\n", req, err)
		fail(c, services.ErrPermissionNotFound, err)
		return
	}

	rev := model.RolePermission{}

	// insert
	id, err := model.CreateRolePermission(rolePermission)
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
 * RolePermission 列表页
 */
func ListRolePermission(c *gin.Context) {
	param := map[string]string{
		"role_id":       c.Query("role_id"),
		"permission_id": c.Query("permission_id"),
	}

	var (
		objectRoleList []*model.RolePermission
		err            error
	)

	req := &model.ListRolePermissionArgs{}
	// object_type
	if q, ok := param["role_id"]; ok && q != "" {
		val, err := strconv.ParseUint(q, 10, 64)
		if err != nil {
			fail(c, services.ErrRequestInvalid, errors.New("invalid object_type"))
			return
		}
		req.RoleId = val
	}
	// object_id
	if q, ok := param["permission_id"]; ok && q != "" {
		val, err := strconv.ParseUint(q, 10, 64)
		if err != nil {
			fail(c, services.ErrRequestInvalid, errors.New("invalid object_id"))
			return
		}
		req.PermissionId = val
	}

	objectRoleList, err = model.ListRolePermission(req)
	if err != nil {
		fmt.Printf("ListRolePermission err:%v\n", err)
		fail(c, services.ErrRequestInvalid, err)
		return
	}

	success(c, objectRoleList)
	return
}

/*
 * 删除 RolePermission
 */
func DeleteRolePermission(c *gin.Context) {
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

	err = model.DeleteRolePermission(idUint)
	if err != nil {
		fmt.Printf("DeleteRolePermission err:%v\n", err)
		fail(c, services.ErrRequestInvalid, err)
		return
	}

	success(c, gin.H{})
	return
}
