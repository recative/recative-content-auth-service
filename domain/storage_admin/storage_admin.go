package storage_admin

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_controller"
	"github.com/recative/recative-backend/domain/storage/storage_model"
	"github.com/recative/recative-backend/domain/storage/storage_service_public"
	"github.com/recative/recative-backend/domain/storage_admin/storage_admin_controller"
	"github.com/recative/recative-backend/domain/storage_admin/storage_admin_route"
	"github.com/recative/recative-backend/domain/storage_admin/storage_admin_service"
	"gorm.io/gorm"
)

type Dependence struct {
	Db                   *gorm.DB
	AdminGroup           *gin.RouterGroup
	AdminTokenController admin_token_controller.Controller
}

func Init(dep *Dependence) {
	model := storage_model.New(dep.Db)

	publicService := storage_service_public.New(dep.Db, model)

	service := storage_admin_service.New(dep.Db, model, publicService)

	controller := storage_admin_controller.New(dep.Db, service)

	storage_admin_route.Init(dep.AdminGroup, controller, dep.AdminTokenController)
}
