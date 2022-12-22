package admin_token_controller

import (
	"github.com/recative/recative-backend/definition_error"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_config"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_format"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_service"
	"github.com/recative/recative-backend/spec"
	"github.com/recative/recative-service-sdk/pkg/gin_context"
	"github.com/recative/recative-service-sdk/pkg/http_engine/http_err"
	"github.com/recative/recative-service-sdk/pkg/http_engine/response"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"time"
)

type Controller interface {
	GetInfoByToken(c *gin_context.NoSecurityContext)
	PutTokenInfo(c *gin_context.NoSecurityContext)
	DeleteToken(c *gin_context.NoSecurityContext)
	CreateToken(c *gin_context.NoSecurityContext)
	GetAllTokens(c *gin_context.NoSecurityContext)
	GetSelectTokens(c *gin_context.NoSecurityContext)
	GetSudoToken(c *gin_context.NoSecurityContext)
	GetTempToken(c *gin_context.NoSecurityContext)
	PostTempToken(c *gin_context.NoSecurityContext)
	CheckAdminTokenPermission(needPermissions ...string) func(c *gin_context.NoSecurityContext)
	//CheckRootToken() func(c *gin_context.NoSecurityContext)
}

type controller struct {
	db      *gorm.DB
	service admin_token_service.Service
	config  admin_token_config.Config
}

func New(db *gorm.DB, service admin_token_service.Service, config admin_token_config.Config) Controller {
	return &controller{
		db:      db,
		service: service,
		config:  config,
	}
}

func (con *controller) GetInfoByToken(c *gin_context.NoSecurityContext) {
	tokenRaw := c.C.Param("token")
	if tokenRaw == "" {
		response.Err(c.C, http_err.InvalidArgument.New())
		return
	}

	token, err := con.service.ReadTokenInfo(tokenRaw)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res spec.TokenResponse
	res = admin_token_format.TokenToResponse(token)
	response.Ok(c.C, res)
}

func (con *controller) PutTokenInfo(c *gin_context.NoSecurityContext) {
	tokenRaw := c.C.Param("token")
	if tokenRaw == "" {
		response.Err(c.C, http_err.InvalidArgument.New())
		return
	}

	var body spec.PutAdminAuthrizationTokenJSONRequestBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
		return
	}

	tokenParam := admin_token_format.TokenRequestToTokenParam(body)

	err = con.service.UpdateTokenInfo(tokenRaw, tokenParam)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	response.Ok(c.C, nil)
}

func (con *controller) DeleteToken(c *gin_context.NoSecurityContext) {
	tokenRaw := c.C.Param("token")
	if tokenRaw == "" {
		response.Err(c.C, http_err.InvalidArgument.New())
		return
	}

	err := con.service.DeleteToken(tokenRaw)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	response.Ok(c.C, nil)
}
func (con *controller) CreateToken(c *gin_context.NoSecurityContext) {
	var body spec.PostAdmainTokenJSONRequestBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
		return
	}

	tokenParam := admin_token_format.TokenRequestToTokenParam(body)

	token, err := con.service.CreateToken(tokenParam)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res spec.RawTokenResponse
	res.Token = token.Raw
	response.Ok(c.C, res)
}

func (con *controller) GetAllTokens(c *gin_context.NoSecurityContext) {
	tokens, err := con.service.ReadAllTokens()
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res []spec.TokenResponse
	res = admin_token_format.TokensToResponses(tokens)
	response.Ok(c.C, res)
}

func (con *controller) GetSelectTokens(c *gin_context.NoSecurityContext) {
	var body spec.PostAdminAuthTokensJSONRequestBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
		return
	}

	tokens, err := con.service.ReadSelectTokens(body)

	var res []spec.TokenResponse
	res = admin_token_format.TokensToResponses(tokens)
	response.Ok(c.C, res)
}

func (con *controller) GetSudoToken(c *gin_context.NoSecurityContext) {
	sudoToken, err := con.service.CreateSudoToken()
	if err != nil {
		response.Err(c.C, err)
		return
	}

	response.Ok(c.C, spec.RawTokenResponse{
		Token: sudoToken,
	})
	return
}

func (con *controller) GetTempToken(c *gin_context.NoSecurityContext) {
	expiresAtString := c.C.Param("expires_at")

	var time_ *time.Time
	if expiresAtString != "" {
		time__, err := time.Parse(time.RFC3339, expiresAtString)
		if err != nil {
			response.Err(c.C, definition_error.TimeFormatNotSupported.Wrap(err))
			return
		}
		time_ = &time__
	}

	token, err := con.service.GenerateTempUserTokenWithAllPermission(time_)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res spec.RawTokenResponse
	res.Token = token
	response.Ok(c.C, res)
	return
}

func (con *controller) PostTempToken(c *gin_context.NoSecurityContext) {
	var body spec.PostAdminTempUserTokenJSONBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
		return
	}

	var time_ *time.Time
	if body.ExpiresAt != nil {
		time__, err := time.Parse(time.RFC3339, *body.ExpiresAt)
		if err != nil {
			response.Err(c.C, definition_error.TimeFormatNotSupported.Wrap(err))
			return
		}
		time_ = &time__
	}

	token, err := con.service.GenerateTempUserToken(body.Permissions, body.GuardPermissionExist, time_)
	var res spec.RawTokenResponse
	res.Token = token
	response.Ok(c.C, res)
	return
}

func (con *controller) CheckAdminTokenPermission(needPermissions ...string) func(c *gin_context.NoSecurityContext) {
	return func(c *gin_context.NoSecurityContext) {
		internalAuthorizationToken := c.C.GetHeader("X-InternalAuthorization")
		if con.config.RootToken == internalAuthorizationToken {
			return
		}
		ok := con.service.IsTokenExist(internalAuthorizationToken)
		if !ok {
			response.Err(c.C, http_err.Unauthorized.New())
			return
		}
		token, err := con.service.ReadTokenInfo(internalAuthorizationToken)
		if err != nil {
			response.Err(c.C, err)
			return
		}

		for _, permission := range needPermissions {
			if lo.Contains(token.AdminPermission, permission) {
				return
			}
		}

		response.Err(c.C, http_err.Forbidden.New("permission denied"))
		return
	}
}

//func (con *controller) CheckRootToken() func(c *gin_context.NoSecurityContext) {
//	return func(c *gin_context.NoSecurityContext) {
//		var body spec.PostAdminSudoJSONBody
//		err := c.C.ShouldBindJSON(&body)
//		if err != nil {
//			response.Err(c.C, http_err.InvalidArgument.Wrap(err))
//			return
//		}
//
//		if body.RootToken != con.config.RootToken {
//			response.Err(c.C, http_err.Unauthorized.New("root token in request body is invalid"))
//			return
//		}
//
//		c.C.Next()
//		return
//	}
//}
