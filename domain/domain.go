package domain

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/recative/recative-backend/domain/storage"
	"github.com/recative/recative-backend/domain/storage_admin"
	"github.com/recative/recative-backend/pkg"
	"github.com/recative/recative-backend/pkg/http_engine/middleware"
	"github.com/recative/recative-backend/spec"
	"github.com/recative/recative-backend/utils/must"
)

func Init(dep *pkg.Dependence) {
	var apiSpec = func() *openapi3.T {
		swagger, err := spec.GetSwagger()
		must.Must(err)
		return swagger
	}()

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
