package permission_controller

import (
	"github.com/recative/recative-backend-sdk/pkg/gin_context"
	"github.com/recative/recative-backend-sdk/pkg/http_engine/http_err"
	"github.com/recative/recative-backend-sdk/pkg/http_engine/response"
	"github.com/recative/recative-backend/domain/permission/permission_format"
	"github.com/recative/recative-backend/domain/permission/permission_service"
	"github.com/recative/recative-backend/spec"
	"gorm.io/gorm"
)

type Controller interface {
	GetPermissionById(c *gin_context.InternalContext)
	PutPermissionById(c *gin_context.InternalContext)
	DeletePermissionById(c *gin_context.InternalContext)
	CreatePermission(c *gin_context.InternalContext)
	BatchGetPermission(c *gin_context.InternalContext)
	GetAllPermissions(c *gin_context.InternalContext)
}

type controller struct {
	db      *gorm.DB
	service permission_service.Service
}

func New(db *gorm.DB, service permission_service.Service) Controller {
	return &controller{
		db:      db,
		service: service,
	}
}

func (con *controller) GetPermissionById(c *gin_context.InternalContext) {
	permissionId := c.C.Param("permission_id")

	permission, err := con.service.ReadPermissionById(permissionId)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res spec.PermissionResponse
	res = permission_format.PermissionToResponse(permission)
	response.Ok(c.C, res)
}

func (con *controller) PutPermissionById(c *gin_context.InternalContext) {
	permissionId := c.C.Param("permission_id")

	var body spec.PutAdminPermissionJSONRequestBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
		return
	}

	permissionParams := permission_format.RequestPermissionToPermissionParam(body)

	err = con.service.UpdatePermissionById(permissionId, permissionParams)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	response.Ok(c.C, nil)
}

func (con *controller) DeletePermissionById(c *gin_context.InternalContext) {
	permissionId := c.C.Param("permission_id")

	_, err := con.service.DeletePermissionById(permissionId)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	response.Ok(c.C, nil)
}

func (con *controller) CreatePermission(c *gin_context.InternalContext) {
	var body spec.PostAdminPermissionJSONBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
		return
	}
	permissionParams := permission_format.RequestPermissionToPermissionParam(body)

	err = con.service.CreatePermission(permissionParams)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	response.Ok(c.C, nil)
}

func (con *controller) BatchGetPermission(c *gin_context.InternalContext) {
	var body spec.PostAdminPermissionsJSONRequestBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
		return
	}

	permissions, err := con.service.ReadPermissionsByKeys(body)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res spec.PermissionsResponse
	res = permission_format.PermissionsToResponse(permissions)
	response.Ok(c.C, res)
	return
}

func (con *controller) GetAllPermissions(c *gin_context.InternalContext) {
	permissions, err := con.service.ReadAllPermissions()
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res spec.PermissionsResponse
	res = permission_format.PermissionsToResponse(permissions)
	response.Ok(c.C, res)
	return
}