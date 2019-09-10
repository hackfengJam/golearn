package router

import (
	"github.com/gin-gonic/gin"
	"golearn/project/rbac/service"
)

func RegisterRouter(r *gin.Engine) {
	{
		var v1 = r.Group("rbac")
		{
			var ipAddress = v1.Group("/ip_address")
			ipAddress.GET("", service.ListIpAddress)
			ipAddress.POST("", service.SetIpAddress)
			ipAddress.DELETE("/:id", service.DeleteIpAddress)
		}
		{
			var role = v1.Group("/role")
			role.GET("", service.ListRole)
			role.POST("", service.SetRole)
			//role.DELETE("/:id")
		}
		{
			var objectRole = v1.Group("/object_role")
			objectRole.GET("", service.ListObjectRole)
			objectRole.POST("", service.CreateObjectRole)
			objectRole.DELETE("/:id", service.DeleteObjectRole)
		}
		{
			var permission = v1.Group("/permission")
			permission.GET("", service.ListPermission)
			permission.POST("", service.SetPermission)
			permission.DELETE("/:id", service.DeletePermission)
		}
		{
			var rolePermission = v1.Group("/role_permission")
			rolePermission.GET("", service.ListRolePermission)
			rolePermission.POST("", service.CreateRolePermission)
			rolePermission.DELETE("/:id", service.DeleteRolePermission)
		}
	}
}
