package admin_token_model

import "gorm.io/gorm"

type Model interface {
	TokenModel
}

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) Model {
	return &model{
		db: db,
	}
}
