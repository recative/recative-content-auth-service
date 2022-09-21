package middleware

import (
	"context"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	legacyrouter "github.com/getkin/kin-openapi/routers/legacy"
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/pkg/http_engine/http_err"
	"github.com/recative/recative-backend/pkg/http_engine/response"
	"github.com/recative/recative-backend/pkg/logger"
	"go.uber.org/zap"
)

func OpenapiValidator(swagger *openapi3.T) func(ginCtx *gin.Context) {
	router, err := legacyrouter.NewRouter(swagger)
	if err != nil {
		logger.Panic("failed to create openapi validator when parse input params", zap.Error(err))
	}
	return func(ginCtx *gin.Context) {
		route, m, err := router.FindRoute(ginCtx.Request)
		if err != nil {
			response.Err(ginCtx, err)
			ginCtx.Abort()
			return
		}
		input := openapi3filter.RequestValidationInput{
			Request:     ginCtx.Request,
			PathParams:  m,
			QueryParams: ginCtx.Request.URL.Query(),
			Route:       route,
			Options: &openapi3filter.Options{AuthenticationFunc: func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
				if input.SecuritySchemeName == "" {
					return nil
				}
				ginCtx.Set("security_required", input.SecuritySchemeName != "")
				return nil
			},
			},
			ParamDecoder: nil,
		}

		err = openapi3filter.ValidateRequest(context.Background(), &input)
		if err != nil {
			if _, ok := err.(*openapi3filter.SecurityRequirementsError); ok {
				response.Err(ginCtx, http_err.Unauthorized.New("authentication failed mainly cause by missing authorization header"))
			} else if err, ok := err.(*openapi3filter.RequestError); ok {
				response.Err(ginCtx, http_err.InvalidArgument.NewWithPayload("schema miss match", map[string]any{
					"Reason": err.Reason,
					"Err":    err.Err,
				}))
			} else {
				response.Err(ginCtx, http_err.InvalidArgument.Wrap(err))
			}
			ginCtx.Abort()
			return
		}
	}
}
