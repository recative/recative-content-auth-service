package admin_token_service_public

import (
	"github.com/recative/recative-backend/domain/admin_token/admin_token_model"
	"gorm.io/gorm"
)

type Service interface {
}

type service struct {
	db    *gorm.DB
	model admin_token_model.Model
}

func New(db *gorm.DB, model admin_token_model.Model) Service {
	return &service{
		db:    db,
		model: model,
	}
}
