package domain

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/recative/recative-backend-sdk/pkg/auth"
	"github.com/recative/recative-backend-sdk/pkg/gin_context"
	"github.com/recative/recative-backend-sdk/pkg/http_engine"
	"github.com/recative/recative-backend-sdk/pkg/http_engine/middleware"
	"github.com/recative/recative-backend/domain/storage"
	"github.com/recative/recative-backend/domain/storage_admin"
	"github.com/recative/recative-backend/spec"
	"gorm.io/gorm"
)

type Dependence struct {
	Db         *gorm.DB
	HttpEngine *http_engine.CustomHttpEngine
	Auther     auth.Authable
}

func Init(dep *Dependence) {
	var apiSpec = func() *openapi3.T {
		swagger, err := spec.GetSwagger()
		if err != nil {
			panic(err)
		}
		return swagger
	}()

	gin_context.Init(gin_context.ContextDependence{
		Auther:      dep.Auther,
		CustomLogic: nil,
	})

	appGroup := dep.HttpEngine.Group("/app", middleware.OpenapiValidator(apiSpec))
	adminGroup := dep.HttpEngine.Group("/admin", middleware.OpenapiValidator(apiSpec))
	{
		storage.Init(&storage.Dependence{
			Db:       dep.Db,
			AppGroup: appGroup,
		})

		storage_admin.Init(&storage_admin.Dependence{
			Db:         dep.Db,
			AdminGroup: adminGroup,
		})
	}
}
