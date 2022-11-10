package permission_service

import (
	"github.com/recative/recative-backend/domain/permission/permission_service_public"
)

type Service interface {
	permission_service_public.Service
}

type service struct {
	permission_service_public.Service
}

func New(permissionServicePublic permission_service_public.Service) Service {
	return &service{
		permissionServicePublic,
	}
}
