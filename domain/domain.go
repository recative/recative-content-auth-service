package domain

import (
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/recative/recative-backend-sdk/pkg/auth"
	"github.com/recative/recative-backend-sdk/pkg/gin_context"
	"github.com/recative/recative-backend-sdk/pkg/http_engine"
	"github.com/recative/recative-backend-sdk/pkg/http_engine/http_err"
	"github.com/recative/recative-backend/domain/admin_token"
	"github.com/recative/recative-backend/domain/permission"
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

func Init(dep *Dependence, config Config) {
	var apiSpec = func() *openapi3.T {
		swagger, err := spec.GetSwagger()
		if err != nil {
			panic(err)
		}
		return swagger
	}()

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
		storage.Init(&storage.Dependence{
			Db:       dep.Db,
			AppGroup: appGroup,
		})

		storage_admin.Init(&storage_admin.Dependence{
			Db:         dep.Db,
			AdminGroup: adminGroup,
		})

		permission.Init(&permission.Dependence{
			Db:         dep.Db,
			AdminGroup: adminGroup,
		})

		admin_token.Init(&admin_token.Dependence{
			Db:         dep.Db,
			AdminGroup: adminGroup,
			Config:     config.AdminTokenConfig,
		})
	}
}
