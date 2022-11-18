package permission_model

import (
	"github.com/recative/recative-backend-sdk/pkg/db"
	"github.com/recative/recative-backend-sdk/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Model interface {
	PermissionModel
	db.AutoMigrater
}

type model struct {
	db *gorm.DB
}

func (m *model) AutoMigrate() {
	err := m.db.AutoMigrate(
		&Permission{},
	)

	if err != nil {
		logger.Fatal("permission model auto migrate error", zap.Error(err))
	}
}

func New(db *gorm.DB) Model {
	return &model{
		db,
	}
}
