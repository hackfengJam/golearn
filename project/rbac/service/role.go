package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golearn/project/rbac/model"
	services "golearn/project/rbac/utils/bizresp"
	"io"
)

func BindJSON(c *gin.Context, obj interface{}) error {
	err := c.ShouldBindWith(obj, binding.JSON)
	if err == io.EOF {
		return errors.New("empty request body")
	}
	return err
}

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

	if req.Id == 0 {
		// insert
		err := model.CreateRole(role)
		if err != nil {
			fmt.Printf("args:%v, err:%v\n", req, err)
			fail(c, services.ErrRequestInvalid, err)
			return
		}
	} else {
		// update
		err := model.UpdateRole(role)
		if err != nil {
			fmt.Printf("args:%v, err:%v\n", req, err)
			fail(c, services.ErrRequestInvalid, err)
			return
		}
	}
	success(c, gin.H{})
	return
}
