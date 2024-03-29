package storage_model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/recative/recative-service-sdk/pkg/db/db_err"
	"gorm.io/gorm"
	"time"
)

type StorageModel interface {
	CreateStorage(storageParams StorageParams) error
	UpdateStorageByKey(key string, storageParams StorageParams) error
	ReadStorageByKey(key string, isIncludeValue bool) (*Storage, error)
	ReadStorageByKeys(key []string, isIncludeValue bool) ([]*Storage, error)
	ReadAllStorage(isIncludeValue bool) ([]*Storage, error)
	DeleteStorageByKey(key string) (*Storage, error)
	IsExistStorageByKey(key string) (bool, error)
	ReadStoragesByQuery(keys, excludePermission, includePermission []string, isIncludeValue bool) ([]*Storage, error)
}

type Storage struct {
	CreatedAt time.Time
	UpdatedAt time.Time

	StorageId string `gorm:"primaryKey;not null;uniqueIndex"`

	StorageParams `gorm:"embedded"`
}

func isIncludeValue(isIncludeValue bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if !isIncludeValue {
			return db.Model(&Storage{}).Select([]string{
				"created_at",
				"updated_at",
				"storage_id",
				"key",
				"need_permissions",
				"need_permission_count",
				"comment",
			})
		} else {
			return db.Model(&Storage{})
		}
	}
}

type StorageParams struct {
	Key                 string `gorm:"not null;uniqueIndex"`
	Value               string
	NeedPermissions     pq.StringArray `gorm:"type:varchar[]"`
	NeedPermissionCount int
	IsPublic            bool

	Comment string
}

func (s *Storage) BeforeCreate(tx *gorm.DB) (err error) {
	s.StorageId = uuid.New().String()
	return
}

func (m *model) CreateStorage(storageParams StorageParams) error {
	storage := Storage{
		StorageParams: storageParams,
	}
	err := m.db.Create(&storage).Error
	if err != nil {
		return db_err.Wrap(err)
	}
	return nil
}

func (m *model) UpdateStorageByKey(key string, storageParams StorageParams) error {
	err := m.db.Where("key = ?", key).Updates(Storage{StorageParams: storageParams}).Error
	if err != nil {
		return db_err.Wrap(err)
	}
	return nil
}

func (m *model) ReadStorageByKey(key string, isIncludeValue_ bool) (*Storage, error) {
	var storage Storage

	err := m.db.Scopes(isIncludeValue(isIncludeValue_)).Where("key = ?", key).First(&storage).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return &storage, nil
}

func (m *model) ReadStorageByKeys(key []string, isIncludeValue_ bool) ([]*Storage, error) {
	var storages []*Storage
	err := m.db.Scopes(isIncludeValue(isIncludeValue_)).Where("key IN ?", key).Find(&storages).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return storages, nil
}

func (m *model) ReadAllStorage(isIncludeValue_ bool) ([]*Storage, error) {
	var storages []*Storage
	err := m.db.Scopes(isIncludeValue(isIncludeValue_)).Find(&storages).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return storages, nil
}

func (m *model) DeleteStorageByKey(key string) (*Storage, error) {
	var storage Storage
	err := m.db.Where("key = ?", key).Delete(&storage).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return &storage, nil
}

func (m *model) IsExistStorageByKey(key string) (bool, error) {
	var storage Storage
	err := m.db.Where("key = ?", key).First(&storage).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, db_err.Wrap(err)
	}
	return true, nil
}

func (m *model) ReadStoragesByQuery(keys, excludePermission, includePermission []string, isIncludeValue_ bool) ([]*Storage, error) {
	var storages []*Storage
	var db = m.db

	db = db.Scopes(isIncludeValue(isIncludeValue_))
	if len(keys) > 0 {
		db = db.Where("key IN ?", keys)
	}
	if len(excludePermission) > 0 {
		db = db.Not("need_permissions && ?", pq.StringArray(excludePermission))
	}
	if len(includePermission) > 0 {
		db = db.Where("need_permissions && ?", pq.StringArray(includePermission))
	}

	err := db.Find(&storages).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return storages, nil
}
