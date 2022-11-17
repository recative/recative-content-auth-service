package admin_token_service

import (
	"github.com/google/uuid"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_model"
	"github.com/recative/recative-backend/domain/admin_token/admin_token_service_public"
	"gorm.io/gorm"
)

type Service interface {
	admin_token_service_public.Service
	ReadTokenInfo(token string) (*admin_token_model.Token, error)
	UpdateTokenInfo(tokenRaw string, param admin_token_model.TokenParam) error
	DeleteToken(tokenRaw string) error
	CreateToken(param admin_token_model.TokenParam) (token admin_token_model.Token, err error)
	ReadAllTokens() (token []admin_token_model.Token, err error)
	ReadSelectTokens(tokenRaws []string) (token []admin_token_model.Token, err error)
	CreateSudoToken(inputRootToken string) (sudoToken string, err error)
	IsTokenExist(token string) bool
}

type service struct {
	db    *gorm.DB
	model admin_token_model.Model
	admin_token_service_public.Service
}

func New(db *gorm.DB, model admin_token_model.Model, publicService admin_token_service_public.Service) Service {
	return &service{
		db:      db,
		model:   model,
		Service: publicService,
	}
}

func (s *service) ReadTokenInfo(tokenRaw string) (*admin_token_model.Token, error) {
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
	return s.model.ReadSelectTokens(tokenRaws)
}

var SudoToken string

func (s *service) CreateSudoToken(inputRootToken string) (sudoToken string, err error) {
	if SudoToken != "" {
		return SudoToken, nil
	}
	return uuid.New().String(), nil
}

func (s *service) IsTokenExist(token string) bool {
	return s.model.IsTokenExist(token)
}
