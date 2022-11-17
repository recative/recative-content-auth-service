package permission_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend-sdk/pkg/gin_context"
	"github.com/recative/recative-backend/domain/permission/permission_controller"
)

func Init(adminGroup *gin.RouterGroup, permissionController permission_controller.Controller) {
	adminGroup.GET("/:permission_id", gin_context.InternalHandler(permissionController.GetPermissionById))
	adminGroup.PUT("/:permission_id", gin_context.InternalHandler(permissionController.PutPermissionById))
	adminGroup.DELETE("/:permission_id", gin_context.InternalHandler(permissionController.DeletePermissionById))
	adminGroup.POST("/permission", gin_context.InternalHandler(permissionController.CreatePermission))

	adminGroup.POST("/storages", gin_context.InternalHandler(permissionController.BatchGetPermission))
	adminGroup.GET("/storages", gin_context.InternalHandler(permissionController.GetAllPermissions))
}
