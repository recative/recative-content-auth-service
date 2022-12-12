package storage_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/domain/storage/storage_controller"
	"github.com/recative/recative-service-sdk/pkg/gin_context"
)

func Init(appGroup *gin.RouterGroup, storageController storage_controller.Controller) {
	appGroup.POST("/storage", gin_context.Handler(storageController.PostAppStorage))
}
