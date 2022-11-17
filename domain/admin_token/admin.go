package admin_token

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_config"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_controller"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_model"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_route"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_service"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_service_public"
	"gorm.io/gorm"
)

type Dependence struct {
	Db         *gorm.DB
	AdminGroup *gin.RouterGroup
	Config     admin_token_config.Config
}

func Init(dep *Dependence) {
	model := admin_token_model.New(dep.Db)

	publicService := admin_token_service_public.New(dep.Db, model)

	service := admin_token_service.New(dep.Db, model, publicService)

	controller := admin_token_controller.New(dep.Db, service, dep.Config)

	admin_token_route.Init(dep.AdminGroup, controller)
}
