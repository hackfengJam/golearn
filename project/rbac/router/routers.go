package router

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine) {
	{
		var v1 = r.Group("rbac")
		{
			var ipAddress = v1.Group("/permission")
			ipAddress.GET("/:id")
			ipAddress.POST("")
			ipAddress.DELETE("/:id")
		}
		{
			var role = v1.Group("/role")
			role.GET("/:id")
			role.POST("")
			role.DELETE("/:id")
		}
		{
			var objectRole = v1.Group("/object_role")
			objectRole.GET("/:id")
			objectRole.POST("")
			objectRole.DELETE("/:id")
		}
		{
			var permission = v1.Group("/permission")
			permission.GET("/:id")
			permission.POST("")
			permission.DELETE("/:id")
		}
		{
			var rolePermission = v1.Group("/role_permission")
			rolePermission.GET("/:id")
			rolePermission.POST("")
			rolePermission.DELETE("/:id")
		}
	}
}
