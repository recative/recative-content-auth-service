package permission_format

import (
	"github.com/recative/recative-backend/domain/permission/permission_model"
	"github.com/recative/recative-backend/spec"
)

func PermissionToResponse(permission *permission_model.Permission) spec.PermissionResponse {
	return spec.PermissionResponse{
		Comment: &permission.Comment,
		Id:      permission.Id,
	}
}

func PermissionsToResponse(permissions []*permission_model.Permission) []spec.PermissionResponse {
	res := make([]spec.PermissionResponse, 0, len(permissions))
	for _, p := range permissions {
		res = append(res, PermissionToResponse(p))
	}
	return res
}

func RequestPermissionToPermissionParam(req spec.PermissionRequest) permission_model.PermissionParams {
	var comment string
	if req.Comment != nil {
		comment = *req.Comment
	}

	return permission_model.PermissionParams{
		Id:      req.Id,
		Comment: comment,
	}
}
