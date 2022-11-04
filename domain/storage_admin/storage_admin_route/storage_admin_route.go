package storage_admin_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend-sdk/pkg/gin_context"
	"github.com/recative/recative-backend/domain/storage_admin/storage_admin_controller"
)

func Init(adminGroup *gin.RouterGroup, storageAdminController storage_admin_controller.Controller) {
	adminGroup.GET("/:storage_key", gin_context.InternalHandler(storageAdminController.GetStorageByKey))
	adminGroup.PUT("/:storage_key", gin_context.InternalHandler(storageAdminController.PutStorageByKey))
	adminGroup.DELETE("/:storage_key", gin_context.InternalHandler(storageAdminController.DeleteStorageByKey))

	adminGroup.POST("/storage", gin_context.InternalHandler(storageAdminController.CreateStorage))
	adminGroup.POST("/storages", gin_context.InternalHandler(storageAdminController.BatchGetStorage))
	adminGroup.GET("/storages", gin_context.InternalHandler(storageAdminController.GetAllStorages))
}
