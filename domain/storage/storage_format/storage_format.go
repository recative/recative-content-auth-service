package storage_format

import (
	"github.com/recative/recative-backend/domain/storage/storage_model"
	"github.com/recative/recative-backend/spec"
)

func StorageToResponse(storage *storage_model.Storage) spec.StorageResponse {
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
		res = append(res, StorageToResponse(s))
	}
	return res
}

func RequestStorageToStorageParam(req spec.StorageRequest) storage_model.StorageParams {
	var comment string
	if req.Comment != nil {
		comment = *req.Comment
	}

	var needPermissionCount int
	if req.NeedPermissionCount != nil {
		needPermissionCount = *req.NeedPermissionCount
	}

	var needPermission []string
	if req.NeedPermissions != nil {
		needPermission = *req.NeedPermissions
	}

	var isPublic bool
	if req.IsPublic != nil {
		isPublic = *req.IsPublic
	}

	return storage_model.StorageParams{
		Comment:             comment,
		Key:                 req.Key,
		NeedPermissionCount: needPermissionCount,
		NeedPermissions:     needPermission,
		Value:               req.Value,
		IsPublic:            isPublic,
	}
}
