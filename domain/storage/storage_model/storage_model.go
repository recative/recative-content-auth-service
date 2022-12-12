package storage_model

import (
	"github.com/recative/recative-service-sdk/pkg/db"
	"github.com/recative/recative-service-sdk/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Model interface {
	StorageModel
	db.AutoMigrater
}

type model struct {
	db *gorm.DB
}

func (m *model) AutoMigrate() {
	err := m.db.AutoMigrate(
		&Storage{},
	)

	if err != nil {
		logger.Fatal("storage model auto migrate error", zap.Error(err))
	}
}

func New(db *gorm.DB) Model {
	return &model{
		db,
	}
}
