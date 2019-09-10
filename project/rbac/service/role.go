package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golearn/project/rbac/model"
	services "golearn/project/rbac/utils/bizresp"
)

/*
 * 添加或者编辑角色页面
 * post 处理添加或者编辑动作
 */
func SetRole(c *gin.Context) {
	req := &model.SetRoleArgs{}
	err := BindJSON(c, req)
	if err != nil {
		fail(c, services.ErrRequestInvalid, err)
		return
	}
	var role = &model.Role{
		Id:   req.Id,
		Name: req.Name,
	}

	rev := model.Role{}

	if req.Id == 0 {
		// insert
		id, err := model.CreateRole(role)
		if err != nil {
			fmt.Printf("args:%v, err:%v\n", req, err)
			fail(c, services.ErrRequestInvalid, err)
			return
		}
		rev.Id = id
	} else {
		// update
		err := model.UpdateRole(role)
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
 * 角色列表页
 */
func ListRole(c *gin.Context) {
	// list
	var roleList []*model.Role
	roleList, err := model.ListRole()
	if err != nil {
		fmt.Printf("ListRole err:%v\n", err)
		fail(c, services.ErrRequestInvalid, err)
		return
	}
	success(c, roleList)
	return
}
