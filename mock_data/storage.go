package mock_data

import (
	"github.com/recative/recative-backend/domain/storage/storage_model"
	"github.com/recative/recative-service-sdk/pkg/db/mock"
	"gorm.io/gorm"
)

func InitStorage(db *gorm.DB) {
	err := mock.Mock(db, false,
		testStorage1,
		testStorage2,
	)
	if err != nil {
		panic(err)
	}
}

var testStorage1 = storage_model.Storage{
	StorageParams: storage_model.StorageParams{
		Key:                 "a",
		Value:               "b",
		NeedPermissions:     []string{"test_permission_1", "test_permission_2"},
		NeedPermissionCount: 2,
		Comment:             "test1",
	},
}

var testStorage2 = storage_model.Storage{
	StorageParams: storage_model.StorageParams{
		Key:                 "aa",
		Value:               "bb",
		NeedPermissions:     []string{"test_permission_1", "test_permission_2"},
		NeedPermissionCount: 2,
		Comment:             "test2",
	},
}
