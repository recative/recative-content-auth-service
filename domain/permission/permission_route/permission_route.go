package permission_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend-sdk/pkg/gin_context"
	"github.com/recative/recative-backend/domain/permission/permission_controller"
)

func Init(adminGroup *gin.RouterGroup, permissionController permission_controller.Controller) {
	adminGroup.GET("/:storage_key", gin_context.InternalHandler(permissionController.GetPermissionById))
	adminGroup.PUT("/:storage_key", gin_context.InternalHandler(permissionController.PutPermissionById))
	adminGroup.DELETE("/:storage_key", gin_context.InternalHandler(permissionController.DeletePermissionById))

	adminGroup.POST("/storage", gin_context.InternalHandler(permissionController.CreatePermission))
	adminGroup.POST("/storages", gin_context.InternalHandler(permissionController.BatchGetPermission))
	adminGroup.GET("/storages", gin_context.InternalHandler(permissionController.GetAllPermissions))
}
