package domain

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/recative/recative-backend/domain/admin_token"
	"github.com/recative/recative-backend/domain/permission"
	"github.com/recative/recative-backend/domain/permission/permission_model"
	"github.com/recative/recative-backend/domain/permission/permission_service_public"
	"github.com/recative/recative-backend/domain/storage"
	"github.com/recative/recative-backend/domain/storage_admin"
	"github.com/recative/recative-service-sdk/pkg/auth"
	"github.com/recative/recative-service-sdk/pkg/db"
	"github.com/recative/recative-service-sdk/pkg/gin_context"
	"github.com/recative/recative-service-sdk/pkg/http_engine"
	"github.com/recative/recative-service-sdk/pkg/http_engine/http_err"
	"gorm.io/gorm"
)

type Dependence struct {
	Db         *gorm.DB
	HttpEngine *http_engine.CustomHttpEngine
	Auther     auth.Authable
	DbConfig   db.Config
}

func Init(dep *Dependence, config Config) {
	//var apiSpec = func() *openapi3.T {
	//	swagger, err := spec.GetSwagger()
	//	if err != nil {
	//		panic(err)
	//	}
	//	return swagger
	//}()

	gin_context.Init(gin_context.ContextDependence{
		Auther: dep.Auther,
		CustomLogic: func(claims jwt.MapClaims, c *gin.Context) error {
			_, securityRequired := c.Get("security_required")

			userId := claims["user_id"]

			if userId == 0 && securityRequired {
				return http_err.InvalidArgument.New(fmt.Sprintf("invalid user id %#v", userId))
			}

			return nil
		},
	})

	appGroup := dep.HttpEngine.Group("/app")
	adminGroup := dep.HttpEngine.Group("/admin")
	{

		adminTokenController := admin_token.Init(&admin_token.Dependence{
			Db:                      dep.Db,
			AdminGroup:              adminGroup,
			Config:                  config.AdminTokenConfig,
			DbConfig:                dep.DbConfig,
			Auther:                  dep.Auther,
			PermissionServicePublic: permission_service_public.New(dep.Db, permission_model.New(dep.Db)),
		})

		storage.Init(&storage.Dependence{
			Db:       dep.Db,
			AppGroup: appGroup,
			DbConfig: dep.DbConfig,
		})

		storage_admin.Init(&storage_admin.Dependence{
			Db:                   dep.Db,
			AdminGroup:           adminGroup,
			AdminTokenController: adminTokenController,
		})

		permission.Init(&permission.Dependence{
			Db:                   dep.Db,
			AdminGroup:           adminGroup,
			DbConfig:             dep.DbConfig,
			AdminTokenController: adminTokenController,
		})
	}
}
