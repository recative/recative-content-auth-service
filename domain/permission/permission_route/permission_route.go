package permission_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_controller"
	"github.com/recative/recative-backend/domain/permission/permission_controller"
	"github.com/recative/recative-service-sdk/pkg/gin_context"
)

func Init(adminGroup *gin.RouterGroup, permissionController permission_controller.Controller, adminTokenController admin_token_controller.Controller) {
	adminGroup.GET("/permission/:permission_id",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("read")),
		gin_context.NoSecurityHandler(permissionController.GetPermissionById),
	)
	adminGroup.PUT("/permission/:permission_id",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("write")),
		gin_context.NoSecurityHandler(permissionController.PutPermissionById),
	)
	adminGroup.DELETE("/permission/:permission_id",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("write")),
		gin_context.NoSecurityHandler(permissionController.DeletePermissionById),
	)
	adminGroup.POST("/permission",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("write")),
		gin_context.NoSecurityHandler(permissionController.CreatePermission),
	)

	adminGroup.GET("/permissions",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("read")),
		gin_context.NoSecurityHandler(permissionController.GetAllPermissions),
	)
	adminGroup.POST("/permissions",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("read")),
		gin_context.NoSecurityHandler(permissionController.BatchGetPermission),
	)
	adminGroup.POST("/permissions/query",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("read")),
		gin_context.NoSecurityHandler(permissionController.PostGetPermissionByQuery),
	)
}
