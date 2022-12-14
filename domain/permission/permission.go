package permission

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_controller"
	"github.com/recative/recative-backend/domain/permission/permission_controller"
	"github.com/recative/recative-backend/domain/permission/permission_model"
	"github.com/recative/recative-backend/domain/permission/permission_route"
	"github.com/recative/recative-backend/domain/permission/permission_service"
	"github.com/recative/recative-backend/domain/permission/permission_service_public"
	"github.com/recative/recative-service-sdk/pkg/db"
	"gorm.io/gorm"
)

type Dependence struct {
	Db                   *gorm.DB
	AdminGroup           *gin.RouterGroup
	DbConfig             db.Config
	AdminTokenController admin_token_controller.Controller
}

func Init(dep *Dependence) {
	model := permission_model.New(dep.Db)
	if dep.DbConfig.IsAutoMigrate {
		model.AutoMigrate()
	}

	publicService := permission_service_public.New(dep.Db, model)

	service := permission_service.New(publicService)

	controller := permission_controller.New(dep.Db, service)

	permission_route.Init(dep.AdminGroup, controller, dep.AdminTokenController)
}
