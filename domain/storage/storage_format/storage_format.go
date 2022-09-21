package storage_format

import (
	"github.com/recative/recative-backend/domain/storage/storage_model"
	"github.com/recative/recative-backend/spec"
)

func StorageToResponse(storage storage_model.Storage) spec.StorageResponse {
	return spec.StorageResponse{
		Comment:             storage.Comment,
		Key:                 storage.Key,
		NeedPermissionCount: storage.NeedPermissionCount,
		NeedPermissions:     storage.NeedPermissions,
		Value:               storage.Value,
	}
}

func StoragesToResponse(storage []*storage_model.Storage) []spec.StorageResponse {
	res := make([]spec.StorageResponse, 0, len(storage))
	for _, s := range storage {
		res = append(res, StorageToResponse(*s))
	}
	return res
}
