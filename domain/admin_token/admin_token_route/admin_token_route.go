package admin_token_route

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_controller"
	"github.com/recative/recative-service-sdk/pkg/gin_context"
)

func Init(adminGroup *gin.RouterGroup, adminTokenController admin_token_controller.Controller) {
	adminGroup.GET("/token/:token",
		gin_context.NoSecurityHandler(adminTokenController.GetInfoByToken),
	)
	adminGroup.PUT("/token/:token",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("sudo")),
		gin_context.NoSecurityHandler(adminTokenController.PutTokenInfo),
	)
	adminGroup.DELETE("/token/:token",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("sudo")),
		gin_context.NoSecurityHandler(adminTokenController.DeleteToken),
	)
	adminGroup.POST("/token",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("sudo")),
		gin_context.NoSecurityHandler(adminTokenController.CreateToken),
	)
	adminGroup.GET("/tokens",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("sudo")),
		gin_context.NoSecurityHandler(adminTokenController.GetAllTokens),
	)
	adminGroup.POST("/tokens",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("sudo")),
		gin_context.NoSecurityHandler(adminTokenController.GetSelectTokens),
	)

	//adminGroup.POST("/sudo",
	//	gin_context.NoSecurityHandler(adminTokenController.CheckRootToken()),
	//	gin_context.NoSecurityHandler(adminTokenController.GetSudoToken))

	adminGroup.GET("/temp_user_token",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("read")),
		gin_context.NoSecurityHandler(adminTokenController.GetTempToken),
	)
	adminGroup.POST("/temp_user_token",
		gin_context.NoSecurityHandler(adminTokenController.CheckAdminTokenPermission("read")),
		gin_context.NoSecurityHandler(adminTokenController.PostTempToken),
	)
}
