package mock_data

import (
	"github.com/recative/recative-backend/domain/permission/permission_model"
	"github.com/recative/recative-service-sdk/pkg/db/mock"
	"gorm.io/gorm"
)

func InitPermission(db *gorm.DB) {
	err := mock.Mock(db, false,
		testPermission1,
		testPermission2,
	)
	if err != nil {
		panic(err)
	}
}

var testPermission1 = permission_model.Permission{
	Id:      "test_permission_1",
	Comment: "this is test permission",
}

var testPermission2 = permission_model.Permission{
	Id:      "test_permission_2",
	Comment: "this is test permission",
}
