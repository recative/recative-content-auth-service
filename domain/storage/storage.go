package storage

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend-sdk/pkg/db"
	"github.com/recative/recative-backend/domain/storage/storage_controller"
	"github.com/recative/recative-backend/domain/storage/storage_model"
	"github.com/recative/recative-backend/domain/storage/storage_route"
	"github.com/recative/recative-backend/domain/storage/storage_service"
	"github.com/recative/recative-backend/domain/storage/storage_service_public"
	"gorm.io/gorm"
)

type Dependence struct {
	Db       *gorm.DB
	AppGroup *gin.RouterGroup
	DbConfig db.Config
}

func Init(dep *Dependence) {
	model := storage_model.New(dep.Db)
	if dep.DbConfig.IsAutoMigrate {
		model.AutoMigrate()
	}

	publicService := storage_service_public.New(dep.Db, model)

	service := storage_service.New(publicService)

	controller := storage_controller.New(dep.Db, service)

	storage_route.Init(dep.AppGroup, controller)
}
