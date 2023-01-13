package storage_service_public

import (
	"github.com/recative/recative-backend/domain/storage/storage_model"
	"github.com/samber/lo"
	"gorm.io/gorm"
)

type Service interface {
	ReadStoragesByKeysAndPermissions(keys []string, permissions []string, isIncludeValue bool) ([]*storage_model.Storage, error)
}

type service struct {
	db    *gorm.DB
	model storage_model.Model
}

func New(db *gorm.DB, model storage_model.Model) Service {
	return &service{
		db,
		model,
	}
}

func (s *service) ReadStoragesByKeysAndPermissions(keys []string, permissions []string, isIncludeValue bool) ([]*storage_model.Storage, error) {
	res := make([]*storage_model.Storage, 0, len(keys))

	for _, key := range keys {
		storage, err := s.model.ReadStorageByKey(key, isIncludeValue)
		if err != nil {
			return nil, err
		}

		if storage.IsPublic == true {
			res = append(res, storage)
			continue
		}
		
		count := lo.CountBy(permissions, func(permission string) bool {
			return lo.Contains(storage.NeedPermissions, permission)
		})

		if count >= storage.NeedPermissionCount {
			res = append(res, storage)
		}
	}

	return res, nil
}
