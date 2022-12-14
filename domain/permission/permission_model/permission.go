package permission_model

import "github.com/recative/recative-service-sdk/pkg/db/db_err"

type PermissionModel interface {
	ReadPermissionById(permissionId string) (*Permission, error)
	UpdatePermissionById(permissionId string, params PermissionParams) error
	DeletePermissionById(permissionId string) (*Permission, error)
	CreatePermission(params PermissionParams) error
	ReadPermissionsByKeys(keys []string) ([]*Permission, error)
	ReadAllPermissions() ([]*Permission, error)
	IsPermissionsExist([]string) ([]string, bool)
	ReadPermissionByQuery(ids []string, query string) ([]*Permission, error)
}

type Permission struct {
	Id      string `gorm:"not null;uniqueIndex"`
	Comment string
}

type PermissionParams struct {
	Id      string
	Comment string
}

func (m *model) ReadPermissionById(permissionId string) (*Permission, error) {
	var permission Permission
	err := m.db.First(&permission, "id = ?", permissionId).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return &permission, nil
}

func (m *model) UpdatePermissionById(permissionId string, params PermissionParams) error {
	err := m.db.Model(&Permission{}).Where("id = ?", permissionId).Updates(Permission{
		Comment: params.Comment,
	}).Error
	if err != nil {
		return db_err.Wrap(err)
	}
	return nil
}

func (m *model) DeletePermissionById(permissionId string) (*Permission, error) {
	var permission Permission
	err := m.db.Where("id = ?", permissionId).Delete(&permission).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return &permission, nil
}

func (m *model) CreatePermission(params PermissionParams) error {
	permission := Permission{
		Id:      params.Id,
		Comment: params.Comment,
	}
	err := m.db.Create(&permission).Error
	if err != nil {
		return db_err.Wrap(err)
	}

	return nil
}

func (m *model) ReadPermissionsByKeys(keys []string) ([]*Permission, error) {
	var permissions []*Permission
	err := m.db.Where("id IN ?", keys).Find(&permissions).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return permissions, nil
}

func (m *model) ReadAllPermissions() ([]*Permission, error) {
	var permissions []*Permission
	err := m.db.Find(&permissions).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return permissions, nil
}

func (m *model) IsPermissionsExist(permissionIds []string) (miss []string, ok bool) {
	var permissions []*Permission
	err := m.db.Where("id IN ?", permissionIds).Find(&permissions).Error
	if err != nil {
		return nil, false
	}
	if len(permissions) != len(permissionIds) {
		miss = make([]string, 0)
		for _, permissionId := range permissionIds {
			var found bool
			for _, permission := range permissions {
				if permission.Id == permissionId {
					found = true
					break
				}
			}
			if !found {
				miss = append(miss, permissionId)
			}
		}
		return miss, false
	}
	return nil, true
}

func (m *model) ReadPermissionByQuery(ids []string, query string) ([]*Permission, error) {
	var permissions []*Permission
	var db = m.db
	if len(ids) > 0 {
		db = db.Where("id IN ?", ids)
	}
	if query != "" {
		db = db.Where("id ~* ?", query)
	}

	err := db.Find(&permissions).Error
	if err != nil {
		return nil, db_err.Wrap(err)
	}
	return permissions, nil
}
