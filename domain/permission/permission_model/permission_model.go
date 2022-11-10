package permission_model

import "gorm.io/gorm"

type Model interface {
	PermissionModel
}

type model struct {
	*gorm.DB
}

func New(db *gorm.DB) Model {
	return &model{
		db,
	}
}
