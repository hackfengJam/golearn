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
 * 添加或者编辑 ip 地址页面
 * post 处理添加或者编辑动作
 */
func SetPermission(c *gin.Context) {
	req := &model.SetPermissionArgs{}
	err := BindJSON(c, req)
	if err != nil {
		fail(c, services.ErrRequestInvalid, err)
		return
	}
	var permission = &model.Permission{
		Id:        req.Id,
		Title:     req.Title,
		EntityKey: req.EntityKey,
	}

	rev := model.Permission{}

	if req.Id == 0 {
		// insert
		id, err := model.CreatePermission(permission)
		if err != nil {
			fmt.Printf("args:%v, err:%v\n", req, err)
			fail(c, services.ErrRequestInvalid, err)
			return
		}
		rev.Id = id
	} else {
		// update
		err := model.UpdatePermission(permission)
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
func ListPermission(c *gin.Context) {
	var (
		permissionList []*model.Permission
		err            error
	)
	permissionList, err = model.ListPermission()
	if err != nil {
		fmt.Printf("ListPermission err:%v\n", err)
		fail(c, services.ErrRequestInvalid, err)
		return
	}

	success(c, permissionList)
	return
}

/*
 * DeletePermission
 */
func DeletePermission(c *gin.Context) {
	// delete by id
	id := c.Param("id")
	if id == "" {
		fail(c, services.ErrRequestInvalid, nil)
		return
	}
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fail(c, services.ErrRequestInvalid, errors.New("invalid permission_id"))
		return
	}

	err = model.DeletePermission(idUint)
	if err != nil {
		fmt.Printf("DeletePermission err:%v\n", err)
		fail(c, services.ErrRequestInvalid, err)
		return
	}

	success(c, gin.H{})
	return
}
