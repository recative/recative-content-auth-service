package storage_admin_controller

import (
	"github.com/recative/recative-backend/domain/admin_token/admin_token_controller"
	"github.com/recative/recative-backend/domain/storage/storage_format"
	"github.com/recative/recative-backend/domain/storage_admin/storage_admin_service"
	"github.com/recative/recative-backend/spec"
	"github.com/recative/recative-service-sdk/pkg/gin_context"
	"github.com/recative/recative-service-sdk/pkg/http_engine/http_err"
	"github.com/recative/recative-service-sdk/pkg/http_engine/response"
	"gorm.io/gorm"
)

type Controller interface {
	GetStorageByKey(c *gin_context.NoSecurityContext)
	PutStorageByKey(c *gin_context.NoSecurityContext)
	DeleteStorageByKey(c *gin_context.NoSecurityContext)
	CreateStorage(c *gin_context.NoSecurityContext)
	BatchGetStorage(c *gin_context.NoSecurityContext)
	GetAllStorages(c *gin_context.NoSecurityContext)
}

type controller struct {
	db                   *gorm.DB
	service              storage_admin_service.Service
	adminTokenController admin_token_controller.Controller
}

func New(db *gorm.DB, service storage_admin_service.Service, adminTokenController admin_token_controller.Controller) Controller {
	return &controller{
		db:                   db,
		service:              service,
		adminTokenController: adminTokenController,
	}
}

func (con *controller) GetStorageByKey(c *gin_context.NoSecurityContext) {
	storageKey := c.C.Param("storage_key")

	storage, err := con.service.ReadStorageByKey(storageKey)
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

func (con *controller) BatchGetStorage(c *gin_context.NoSecurityContext) {
	var body spec.PostAdminStoragesJSONRequestBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
		return
	}

	storages, err := con.service.ReadStoragesByKeys(body)
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res []spec.StorageResponse
	res = storage_format.StoragesToResponse(storages)
	response.Ok(c.C, res)
	return
}

func (con *controller) GetAllStorages(c *gin_context.NoSecurityContext) {
	storages, err := con.service.ReadAllStorages()
	if err != nil {
		response.Err(c.C, err)
		return
	}

	var res []spec.StorageResponse
	res = storage_format.StoragesToResponse(storages)
	response.Ok(c.C, res)
	return
}
