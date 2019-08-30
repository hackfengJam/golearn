package service

import (
	"github.com/gin-gonic/gin"
	services "golearn/project/rbac/utils/bizresp"
	"net/http"
)

func GetPing(c *gin.Context) {
	success(c, gin.H{"message": "pong"})
}

// 成功响应
func success(c *gin.Context, d interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errorCode": "",
		"data":      d,
	})
}

// 失败响应
func fail(c *gin.Context, code string, err error) {
	// 应统一要求，错误码返回非200
	if e, ok := err.(*services.Error); ok {
		// 如果是app内的错误类型
		if code == services.ErrAuthRequired {
			// 如果是未登录，返回401登录要求
			c.JSON(http.StatusUnauthorized, gin.H{
				"errorCode": e.Error(),
				"message":   e.Error(),
				"_e":        "app",
			})
		} else {
			// 其他类型错误
			c.JSON(http.StatusBadRequest, gin.H{
				"errorCode": e.Error(),
				"message":   e.Error(),
				"_e":        "app",
			})
		}
	} else {
		// 其他错误类型
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode": services.NewError(code).Error(),
			"message":   err.Error(),
			"_e":        "general",
		})
	}
}
