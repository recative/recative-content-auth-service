package storage_controller

import (
	"github.com/recative/recative-backend/domain/domain_definition"
	"github.com/recative/recative-backend/domain/storage/storage_format"
	"github.com/recative/recative-backend/domain/storage/storage_service"
	"github.com/recative/recative-backend/spec"
	"github.com/recative/recative-service-sdk/pkg/gin_context"
	"github.com/recative/recative-service-sdk/pkg/http_engine/http_err"
	"github.com/recative/recative-service-sdk/pkg/http_engine/response"
	"github.com/recative/recative-service-sdk/util/ref"
	"gorm.io/gorm"
)

type Controller interface {
	PostAppStorage(c *gin_context.Context[domain_definition.JwtPayload])
}

type controller struct {
	db      *gorm.DB
	service storage_service.Service
}

func New(db *gorm.DB, service storage_service.Service) Controller {
	return &controller{
		db:      db,
		service: service,
	}
}

func (con *controller) PostAppStorage(c *gin_context.Context[domain_definition.JwtPayload]) {
	var body spec.PostAppStorageJSONRequestBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
	}

	if body.IsIncludeValue == nil {
		body.IsIncludeValue = ref.T(false)
	}

	storages, err := con.service.ReadStoragesByKeysAndPermissions(body.StorageKeys, c.Payload.Permissions, *body.IsIncludeValue)
	if err != nil {
		response.Err(c.C, http_err.InternalServerError.Wrap(err))
	}

	var res spec.StoragesResponse
	storage_format.StoragesToResponse(storages)
	response.Ok(c.C, res)
	return
}
