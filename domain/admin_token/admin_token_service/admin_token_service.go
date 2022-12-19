package admin_token_service

import (
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
	"github.com/recative/recative-backend/definition_error"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_config"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_model"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_service_public"
	"github.com/recative/recative-backend/domain/permission/permission_service_public"
	"github.com/recative/recative-service-sdk/pkg/auth"
	"github.com/recative/recative-service-sdk/util/ref"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"time"
)

type Service interface {
	admin_token_service_public.Service
	ReadTokenInfo(token string) (*admin_token_model.Token, error)
	UpdateTokenInfo(tokenRaw string, param admin_token_model.TokenParam) error
	DeleteToken(tokenRaw string) error
	CreateToken(param admin_token_model.TokenParam) (token admin_token_model.Token, err error)
	ReadAllTokens() (token []admin_token_model.Token, err error)
	ReadSelectTokens(tokenRaws []string) (token []admin_token_model.Token, err error)
	CreateSudoToken() (sudoToken string, err error)
	IsTokenExist(token string) bool
	IsSudoTokenValid(token string) bool
	GenerateTempUserToken(permissions []string, checkPermissionExist bool, expiresAt *time.Time) (token string, err error)
	GenerateTempUserTokenWithAllPermission(expiresAt *time.Time) (token string, err error)
}

type service struct {
	db    *gorm.DB
	model admin_token_model.Model
	admin_token_service_public.Service
	cache                   *cache.Cache
	permissionPublicService permission_service_public.Service
	auther                  auth.Authable
	config                  admin_token_config.Config
}

func New(db *gorm.DB, model admin_token_model.Model, publicService admin_token_service_public.Service, permissionPublicService permission_service_public.Service, auther auth.Authable, config admin_token_config.Config) Service {
	return &service{
		db:                      db,
		model:                   model,
		Service:                 publicService,
		cache:                   cache.New(5*time.Minute, 10*time.Minute),
		permissionPublicService: permissionPublicService,
		auther:                  auther,
		config:                  config,
	}
}

func (s *service) ReadTokenInfo(tokenRaw string) (*admin_token_model.Token, error) {
	if s.config.RootToken == tokenRaw {
		return ref.T(s.model.GenerateRootToken(tokenRaw)), nil
	}
	if s.IsSudoTokenValid(tokenRaw) {
		return ref.T(s.model.GenerateSudoToken(tokenRaw)), nil
	}
	return s.model.ReadTokenInfoByRaw(tokenRaw)
}

func (s *service) UpdateTokenInfo(tokenRaw string, param admin_token_model.TokenParam) error {
	return s.model.UpdateTokenInfoByRaw(tokenRaw, param)
}

func (s *service) DeleteToken(tokenRaw string) error {
	return s.model.DeleteTokenByRaw(tokenRaw)
}

func (s *service) CreateToken(param admin_token_model.TokenParam) (token admin_token_model.Token, err error) {
	return s.model.CreateToken(param)
}

func (s *service) ReadAllTokens() (token []admin_token_model.Token, err error) {
	return s.model.ReadAllTokens()
}

func (s *service) ReadSelectTokens(tokenRaws []string) (token []admin_token_model.Token, err error) {
	var res = make([]admin_token_model.Token, 0, len(tokenRaws))
	sudoToken, isExist := s.GetSudoToken()
	if isExist && lo.Contains(tokenRaws, sudoToken) {
		res = append(res, s.model.GenerateSudoToken(sudoToken))
	}
	tokens, err := s.model.ReadSelectTokens(tokenRaws)
	if err != nil {
		return nil, err
	}

	res = append(res, tokens...)
	return res, nil
}

const SudoToken = "sudo-token"

func (s *service) CreateSudoToken() (res string, err error) {
	sudoToken, isExist := s.cache.Get(SudoToken)
	if isExist {
		return sudoToken.(string), nil
	}
	var token = uuid.New().String()
	s.cache.Set(SudoToken, token, cache.DefaultExpiration)
	return token, nil
}

func (s *service) IsSudoTokenValid(token string) bool {
	sudoToken, isExist := s.cache.Get(SudoToken)
	if isExist {
		return sudoToken.(string) == token
	}
	return false
}

func (s *service) GetSudoToken() (sudoToken string, isExist bool) {
	sudoTokenInterface, isExist := s.cache.Get(SudoToken)
	if isExist {
		return sudoTokenInterface.(string), true
	}
	return "", false
}

func (s *service) IsTokenExist(token string) bool {
	return s.model.IsTokenExist(token)
}

func (s *service) GenerateTempUserTokenWithAllPermission(expiresAt *time.Time) (token string, err error) {
	if expiresAt == nil {
		expiresAt = ref.T(time.Now().Add(5 * time.Minute))
	}

	s.auther.GenJwt(map[string]any{
		"permissions": "all",
		"exp":         expiresAt.Unix(),
	})

	return "", err
}

func (s *service) GenerateTempUserToken(permissions []string, guardPermissionExist bool, expiresAt *time.Time) (token string, err error) {
	if guardPermissionExist {
		missedPermissions, ok := s.permissionPublicService.IsPermissionsExist(permissions)
		if !ok {
			return "", definition_error.PermissionNotExist.NewWithPayload("permissions not exist", missedPermissions)
		}
	}

	if expiresAt == nil {
		expiresAt = ref.T(time.Now().Add(5 * time.Minute))
	}

	token = s.auther.GenJwt(map[string]any{
		"permissions": permissions,
		"exp":         expiresAt.Unix(),
	})
	return token, nil
}
