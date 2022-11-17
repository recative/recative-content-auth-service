package admin_token_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend-sdk/pkg/gin_context"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_controller"
)

func Init(adminGroup *gin.RouterGroup, adminTokenController admin_token_controller.Controller) {
	adminGroup.GET("/admin/token/:token", gin_context.InternalHandler(adminTokenController.GetInfoByToken))
	adminGroup.PUT("/admin/token/:token", gin_context.InternalHandler(adminTokenController.PutTokenInfo))
	adminGroup.DELETE("/admin/token/:token", gin_context.InternalHandler(adminTokenController.DeleteToken))
	adminGroup.POST("/admin/token", gin_context.InternalHandler(adminTokenController.CreateToken))
	adminGroup.GET("/admin/tokens", gin_context.InternalHandler(adminTokenController.GetAllTokens))
	adminGroup.POST("/admin/tokens", gin_context.InternalHandler(adminTokenController.GetSelectTokens))

	adminGroup.POST("/sudo", gin_context.InternalHandler(adminTokenController.GetSudoToken))
	adminGroup.GET("/temp_user_token", gin_context.InternalHandler(adminTokenController.GetTempToken))
}
