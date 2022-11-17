package storage_admin_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend-sdk/pkg/gin_context"
	"github.com/recative/recative-backend/domain/storage_admin/storage_admin_controller"
)

func Init(adminGroup *gin.RouterGroup, storageAdminController storage_admin_controller.Controller) {
	adminGroup.GET("/storage/:storage_key", gin_context.NoSecurityHandler(storageAdminController.GetStorageByKey))
	adminGroup.PUT("/storage/:storage_key", gin_context.NoSecurityHandler(storageAdminController.PutStorageByKey))
	adminGroup.DELETE("/storage/:storage_key", gin_context.NoSecurityHandler(storageAdminController.DeleteStorageByKey))

	adminGroup.POST("/storage", gin_context.NoSecurityHandler(storageAdminController.CreateStorage))
	adminGroup.POST("/storages", gin_context.NoSecurityHandler(storageAdminController.BatchGetStorage))
	adminGroup.GET("/storages", gin_context.NoSecurityHandler(storageAdminController.GetAllStorages))
}
