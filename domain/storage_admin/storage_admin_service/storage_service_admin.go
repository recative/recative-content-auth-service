package storage_admin_service

import (
	"github.com/recative/recative-backend/domain/storage/storage_model"
	"github.com/recative/recative-backend/domain/storage/storage_service_public"
	"gorm.io/gorm"
)

type Service interface {
	storage_service_public.Service
	ReadStorageByKey(key string) (*storage_model.Storage, error)
	UpdateStorageByKey(key string, storageParams storage_model.StorageParams) error
	DeleteStorageByKey(key string) (*storage_model.Storage, error)
	CreateStorage(storageParams storage_model.StorageParams) error
	ReadStoragesByKeys(keys []string) ([]*storage_model.Storage, error)
	ReadAllStorages() ([]*storage_model.Storage, error)
}

type service struct {
	db    *gorm.DB
	model storage_model.Model
	storage_service_public.Service
}

func New(db *gorm.DB, model storage_model.Model, publicService storage_service_public.Service) Service {
	return &service{
		db:      db,
		model:   model,
		Service: publicService,
	}
}

func (s *service) ReadStorageByKey(key string) (*storage_model.Storage, error) {
	return s.model.ReadStorageByKey(key)
}

func (s *service) UpdateStorageByKey(key string, storageParams storage_model.StorageParams) error {
	return s.model.UpdateStorageByKey(key, storageParams)
}

func (s *service) DeleteStorageByKey(key string) (*storage_model.Storage, error) {
	return s.model.DeleteStorageByKey(key)
}

func (s *service) CreateStorage(storageParams storage_model.StorageParams) error {
	return s.model.CreateStorage(storageParams)
}

func (s *service) ReadStoragesByKeys(keys []string) ([]*storage_model.Storage, error) {
	return s.model.ReadStorageByKeys(keys)
}

func (s *service) ReadAllStorages() ([]*storage_model.Storage, error) {
	return s.model.ReadAllStorage()
}
