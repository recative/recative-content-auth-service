package mock_data

import (
	"github.com/recative/recative-backend/domain/admin_token/admin_token_model"
	"github.com/recative/recative-backend/domain/permission/permission_model"
	"github.com/recative/recative-backend/domain/storage/storage_model"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) {
	autoMigrate(db)

	InitPermission(db)
}

func autoMigrate(db *gorm.DB) {
	permission_model.New(db).AutoMigrate()
	admin_token_model.New(db).AutoMigrate()
	storage_model.New(db).AutoMigrate()
}
