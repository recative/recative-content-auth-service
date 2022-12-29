package storage_admin_controller

import (
	"github.com/recative/recative-backend/domain/storage/storage_format"
	"github.com/recative/recative-backend/domain/storage_admin/storage_admin_service"
	"github.com/recative/recative-backend/spec"
	"github.com/recative/recative-service-sdk/pkg/gin_context"
	"github.com/recative/recative-service-sdk/pkg/http_engine/http_err"
	"github.com/recative/recative-service-sdk/pkg/http_engine/response"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type Controller interface {
	GetStorageByKey(c *gin_context.NoSecurityContext)
	PutStorageByKey(c *gin_context.NoSecurityContext)
	DeleteStorageByKey(c *gin_context.NoSecurityContext)
	CreateStorage(c *gin_context.NoSecurityContext)
	//PostBatchGetStorage(c *gin_context.NoSecurityContext)
	GetStoragesByQuery(c *gin_context.NoSecurityContext)
	//PostStoragesByQuery(c *gin_context.NoSecurityContext)
}

type controller struct {
	db      *gorm.DB
	service storage_admin_service.Service
}

var _ Controller = &controller{}

func New(db *gorm.DB, service storage_admin_service.Service) Controller {
	return &controller{
		db:      db,
		service: service,
	}
}

func (con *controller) GetStorageByKey(c *gin_context.NoSecurityContext) {
	storageKey := c.C.Param("storage_key")
	isIncludeValue, _ := strconv.ParseBool(c.C.Query("include_value"))

	storage, err := con.service.ReadStorageByKey(storageKey, isIncludeValue)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res spec.StorageResponse
	res = storage_format.StorageToResponse(storage)
	response.Ok(c.C, res)
}

func (con *controller) PutStorageByKey(c *gin_context.NoSecurityContext) {
	storageKey := c.C.Param("storage_key")

	var body spec.PutAdminStorageIdJSONRequestBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
		return
	}
	storageParams := storage_format.RequestStorageToStorageParam(body)

	err = con.service.UpdateStorageByKey(storageKey, storageParams)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	response.Ok(c.C, nil)
}

func (con *controller) DeleteStorageByKey(c *gin_context.NoSecurityContext) {
	storageKey := c.C.Param("storage_key")

	_, err := con.service.DeleteStorageByKey(storageKey)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	response.Ok(c.C, nil)
}

func (con *controller) CreateStorage(c *gin_context.NoSecurityContext) {
	var body spec.PostAdminStorageJSONRequestBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
		return
	}
	storageParams := storage_format.RequestStorageToStorageParam(body)

	err = con.service.CreateStorage(storageParams)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	response.Ok(c.C, nil)
}

//func (con *controller) PostBatchGetStorage(c *gin_context.NoSecurityContext) {
//	var body spec.PostAdminStoragesJSONRequestBody
//	err := c.C.ShouldBindJSON(&body)
//	if err != nil {
//		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
//		return
//	}
//
//	if body.IncludeValue == nil {
//		body.IncludeValue = ref.T(false)
//	}
//
//	storages, err := con.service.ReadStoragesByKeys(body.StorageKeys, *body.IncludeValue)
//	if err != nil {
//		response.Err(c.C, err)
//		return
//	}
//
//	var res []spec.StorageResponse
//	res = storage_format.StoragesToResponse(storages)
//	response.Ok(c.C, res)
//	return
//}

func (con *controller) GetStoragesByQuery(c *gin_context.NoSecurityContext) {
	includePermission := strings.Split(c.C.Query("include_permission"), ",")
	excludePermission := strings.Split(c.C.Query("exclude_permission"), ",")
	keys := strings.Split(c.C.Query("keys"), ",")
	isIncludeValue, _ := strconv.ParseBool(c.C.Query("include_value"))

	storages, err := con.service.ReadStoragesByQuery(keys, excludePermission, includePermission, isIncludeValue)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res []spec.StorageResponse
	res = storage_format.StoragesToResponse(storages)
	response.Ok(c.C, res)
	return
}

//func (con *controller) PostStoragesByQuery(c *gin_context.NoSecurityContext) {
//	var body spec.PostAdminStoragesQueryJSONBody
//	err := c.C.ShouldBindQuery(&body)
//	if err != nil {
//		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
//		return
//	}
//
//	if body.IncludeValue == nil {
//		body.IncludeValue = ref.T(false)
//	}
//
//	storages, err := con.service.ReadStoragesByQuery(body.IncludePermission, body.ExcludePermission, *body.IncludeValue)
//	if err != nil {
//		response.Err(c.C, err)
//		return
//	}
//
//	var res []spec.StorageResponse
//	res = storage_format.StoragesToResponse(storages)
//	response.Ok(c.C, res)
//	return
//}
