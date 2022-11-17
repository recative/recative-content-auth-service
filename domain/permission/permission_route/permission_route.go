package permission_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend-sdk/pkg/gin_context"
	"github.com/recative/recative-backend/domain/permission/permission_controller"
)

func Init(adminGroup *gin.RouterGroup, permissionController permission_controller.Controller) {
	adminGroup.GET("/:permission_id", gin_context.NoSecurityHandler(permissionController.GetPermissionById))
	adminGroup.PUT("/:permission_id", gin_context.NoSecurityHandler(permissionController.PutPermissionById))
	adminGroup.DELETE("/:permission_id", gin_context.NoSecurityHandler(permissionController.DeletePermissionById))
	adminGroup.POST("/permission", gin_context.NoSecurityHandler(permissionController.CreatePermission))

	adminGroup.POST("/permissions", gin_context.NoSecurityHandler(permissionController.GetAllPermissions))
	adminGroup.POST("/permissions/query", gin_context.NoSecurityHandler(permissionController.BatchGetPermission))
}
