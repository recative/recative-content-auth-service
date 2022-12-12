package admin_token_model

import (
	"github.com/recative/recative-service-sdk/pkg/db"
	"github.com/recative/recative-service-sdk/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Model interface {
	TokenModel
	db.AutoMigrater
}

type model struct {
	db *gorm.DB
}

func (m *model) AutoMigrate() {
	err := m.db.AutoMigrate(
		&Token{},
	)

	if err != nil {
		logger.Fatal("admin token model auto migrate error", zap.Error(err))
	}
}

func New(db *gorm.DB) Model {
	return &model{
		db: db,
	}
}
