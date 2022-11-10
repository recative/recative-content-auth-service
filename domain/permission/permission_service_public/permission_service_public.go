package permission_service_public

import (
	"github.com/recative/recative-backend/domain/permission/permission_model"
	"gorm.io/gorm"
)

type Service interface {
	ReadPermissionById(permissionId string) (*permission_model.Permission, error)
	UpdatePermissionById(permissionId string, params permission_model.PermissionParams) error
	DeletePermissionById(permissionId string) (*permission_model.Permission, error)
	CreatePermission(params permission_model.PermissionParams) error
	ReadPermissionsByKeys(keys []string) ([]*permission_model.Permission, error)
	ReadAllPermissions() ([]*permission_model.Permission, error)
}

type service struct {
	db    *gorm.DB
	model permission_model.Model
}

func New(db *gorm.DB, model permission_model.Model) Service {
	return &service{
		db,
		model,
	}
}

func (s *service) ReadPermissionById(permissionId string) (*permission_model.Permission, error) {
	return s.model.ReadPermissionById(permissionId)
}

func (s *service) UpdatePermissionById(permissionId string, params permission_model.PermissionParams) error {
	return s.model.UpdatePermissionById(permissionId, params)
}

func (s *service) DeletePermissionById(permissionId string) (*permission_model.Permission, error) {
	return s.model.DeletePermissionById(permissionId)
}

func (s *service) CreatePermission(params permission_model.PermissionParams) error {
	return s.model.CreatePermission(params)
}

func (s *service) ReadPermissionsByKeys(keys []string) ([]*permission_model.Permission, error) {
	return s.model.ReadPermissionsByKeys(keys)
}

func (s *service) ReadAllPermissions() ([]*permission_model.Permission, error) {
	return s.model.ReadAllPermissions()
}
