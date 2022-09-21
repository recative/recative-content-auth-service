package storage_service

import "github.com/recative/recative-backend/domain/storage/storage_service_public"

type Service interface {
	storage_service_public.Service
}

type service struct {
	storage_service_public.Service
}
