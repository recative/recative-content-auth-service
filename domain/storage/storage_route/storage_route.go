package storage_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend-sdk/pkg/gin_context"
	"github.com/recative/recative-backend/domain/storage/storage_controller"
)

func Init(appGroup *gin.RouterGroup, storageController storage_controller.Controller) {
	appGroup.POST("/storage", gin_context.Handler(storageController.PostAppStorage))
}
