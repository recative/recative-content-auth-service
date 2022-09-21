package storage_admin_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/domain/storage_admin/storage_admin_controller"
	"github.com/recative/recative-backend/pkg/gin_context"
)

func Init(adminGroup *gin.RouterGroup, storageAdminController storage_admin_controller.Controller) {
	adminGroup.GET("/:storage_key", gin_context.Handler(storageAdminController.GetStorageByKey))
	adminGroup.PUT("/:storage_key", gin_context.Handler(storageAdminController.PutStorageByKey))
	adminGroup.DELETE("/:storage_key", gin_context.Handler(storageAdminController.DeleteStorageByKey))

	adminGroup.POST("/storage", gin_context.Handler(storageAdminController.CreateStorage))
	adminGroup.POST("/storages", gin_context.Handler(storageAdminController.BatchGetStorage))
}
