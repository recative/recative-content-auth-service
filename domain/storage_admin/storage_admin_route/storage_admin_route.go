package storage_admin_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_controller"
	"github.com/recative/recative-backend/domain/storage_admin/storage_admin_controller"
	"github.com/recative/recative-service-sdk/pkg/gin_context"
)

func Init(
	adminGroup *gin.RouterGroup,
	storageAdminController storage_admin_controller.Controller,
	adminTokenController admin_token_controller.Controller,
) {
	adminGroup.GET("/storage/:storage_key",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission()),
		gin_context.NoSecurityHandler(storageAdminController.GetStorageByKey),
	)

	adminGroup.PUT("/storage/:storage_key",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission()),
		gin_context.NoSecurityHandler(storageAdminController.PutStorageByKey),
	)
	adminGroup.DELETE("/storage/:storage_key",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission()),
		gin_context.NoSecurityHandler(storageAdminController.DeleteStorageByKey),
	)

	adminGroup.POST("/storage",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission()),
		gin_context.NoSecurityHandler(storageAdminController.CreateStorage),
	)

	adminGroup.POST("/storages",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission()),
		gin_context.NoSecurityHandler(storageAdminController.BatchGetStorage),

	)
	adminGroup.GET("/storages",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission()),
		gin_context.NoSecurityHandler(storageAdminController.GetAllStorages),
	)

	adminGroup.POST("/storages/query",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission()),
		gin_context.NoSecurityHandler(storageAdminController.GetAllStorages),
	)
}
