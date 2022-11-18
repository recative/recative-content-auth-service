package storage_model

import (
	"github.com/google/uuid"
	"github.com/recative/recative-backend-sdk/pkg/db/db_err"
	"gorm.io/gorm"
	"time"
)

type StorageModel interface {
	CreateStorage(storageParams StorageParams) error
	UpdateStorageByKey(key string, storageParams StorageParams) error
	ReadStorageByKey(key string) (*Storage, error)
	ReadStorageByKeys(key []string) ([]*Storage, error)
	ReadAllStorage() ([]*Storage, error)
	DeleteStorageByKey(key string) (*Storage, error)
	IsExistStorageByKey(key string) (bool, error)
}

type Storage struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	StorageId string `gorm:"primaryKey;not null;uniqueIndex"`

	StorageParams
}

type StorageParams struct {
	Key                 string `gorm:"not null;uniqueIndex"`
	Value               string
	NeedPermissions     []string
	NeedPermissionCount int

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

func (m *model) ReadStorageByKey(key string) (*Storage, error) {
	var storage Storage
	err := m.db.Where("key = ?", key).First(&storage).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return &storage, nil
}

func (m *model) ReadStorageByKeys(key []string) ([]*Storage, error) {
	var storages []*Storage
	err := m.db.Where("key IN ?", key).Find(&storages).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return storages, nil
}

func (m *model) ReadAllStorage() ([]*Storage, error) {
	var storages []*Storage
	err := m.db.Find(&storages).Error
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
