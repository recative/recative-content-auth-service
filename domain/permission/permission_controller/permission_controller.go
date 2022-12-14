package permission_controller

import (
	"github.com/recative/recative-backend/domain/permission/permission_format"
	"github.com/recative/recative-backend/domain/permission/permission_service"
	"github.com/recative/recative-backend/domain/utils"
	"github.com/recative/recative-backend/spec"
	"github.com/recative/recative-service-sdk/pkg/gin_context"
	"github.com/recative/recative-service-sdk/pkg/http_engine/http_err"
	"github.com/recative/recative-service-sdk/pkg/http_engine/response"
	"gorm.io/gorm"
	"net/url"
)

type Controller interface {
	GetPermissionById(c *gin_context.NoSecurityContext)
	PutPermissionById(c *gin_context.NoSecurityContext)
	DeletePermissionById(c *gin_context.NoSecurityContext)
	CreatePermission(c *gin_context.NoSecurityContext)
	//BatchGetPermission(c *gin_context.NoSecurityContext)
	//PostGetPermissionByQuery(c *gin_context.NoSecurityContext)
	GetAllPermissions(c *gin_context.NoSecurityContext)
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

func (con *controller) GetPermissionById(c *gin_context.NoSecurityContext) {
	permissionId := c.C.Param("permission_id")
	permissionId, err := url.QueryUnescape(permissionId)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	permission, err := con.service.ReadPermissionById(permissionId)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res spec.PermissionResponse
	res = permission_format.PermissionToResponse(permission)
	response.Ok(c.C, res)
}

func (con *controller) PutPermissionById(c *gin_context.NoSecurityContext) {
	permissionId := c.C.Param("permission_id")
	permissionId, err := url.QueryUnescape(permissionId)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var body spec.PutAdminPermissionJSONRequestBody
	err = c.C.ShouldBindJSON(&body)
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

func (con *controller) DeletePermissionById(c *gin_context.NoSecurityContext) {
	permissionId := c.C.Param("permission_id")
	permissionId, err := url.QueryUnescape(permissionId)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	_, err = con.service.DeletePermissionById(permissionId)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	response.Ok(c.C, nil)
}

func (con *controller) CreatePermission(c *gin_context.NoSecurityContext) {
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

//func (con *controller) BatchGetPermission(c *gin_context.NoSecurityContext) {
//	var body spec.PostAdminPermissionsJSONRequestBody
//	err := c.C.ShouldBindJSON(&body)
//	if err != nil {
//		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
//		return
//	}
//
//	permissions, err := con.service.ReadPermissionsByKeys(body)
//	if err != nil {
//		response.Err(c.C, err)
//		return
//	}
//
//	var res spec.PermissionsResponse
//	res = permission_format.PermissionsToResponse(permissions)
//	response.Ok(c.C, res)
//	return
//}

func (con *controller) GetAllPermissions(c *gin_context.NoSecurityContext) {
	ids := utils.SplitQueryParams("ids", c.C)
	regex := c.C.Query("regex")

	con.service.ReadPermissionsByQuery(ids, regex)
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

//func (con *controller) PostGetPermissionByQuery(c *gin_context.NoSecurityContext) {
//	var body spec.GetAdminPermissionsQueryJSONBody
//	err := c.C.ShouldBindJSON(&body)
//	if err != nil {
//		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
//		return
//	}
//
//	var permissions []*permission_model.Permission
//	if body.Regex == nil {
//		permissions, err = con.service.ReadAllPermissions()
//	} else {
//		permissions, err = con.service.ReadPermissionsByRegexQuery(*body.Regex)
//	}
//	if err != nil {
//		response.Err(c.C, err)
//		return
//	}
//
//	var res spec.PermissionsResponse
//	res = permission_format.PermissionsToResponse(permissions)
//	response.Ok(c.C, res)
//	return
//}
