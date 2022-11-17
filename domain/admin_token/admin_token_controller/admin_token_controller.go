package admin_token_controller

import (
	"github.com/recative/recative-backend-sdk/pkg/gin_context"
	"github.com/recative/recative-backend-sdk/pkg/http_engine/http_err"
	"github.com/recative/recative-backend-sdk/pkg/http_engine/response"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_config"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_format"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_service"
	"github.com/recative/recative-backend/spec"
	"gorm.io/gorm"
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
	CheckAdminTokenPermission(c *gin_context.NoSecurityContext)
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
		response.Err(c.C, http_err.InvalidArgument)
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
		response.Err(c.C, http_err.InvalidArgument)
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
		response.Err(c.C, http_err.InvalidArgument)
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

	var res spec.TokenResponse
	res = admin_token_format.TokenToResponse(&token)
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
	var body spec.PostAdminSudoJSONRequestBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
		return
	}

	sudoToken, err := con.service.CreateSudoToken(body.SudoToken)
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

}

func (con *controller) CheckAdminTokenPermission(c *gin_context.NoSecurityContext) {
	internalAuthorizationToken := c.C.GetHeader("X-InternalAuthorization")
	if internalAuthorizationToken == admin_token_service.SudoToken {
		return
	}
	ok := con.service.IsTokenExist(internalAuthorizationToken)
	if ok {
		return
	}
	response.Err(c.C, http_err.Unauthorized)
	c.C.Abort()
	return
}
