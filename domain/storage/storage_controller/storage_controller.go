package storage_controller

import (
	"github.com/recative/recative-backend/definition"
	"github.com/recative/recative-backend/domain/storage/storage_format"
	"github.com/recative/recative-backend/domain/storage/storage_service"
	"github.com/recative/recative-backend/spec"
	"github.com/recative/recative-service-sdk/pkg/gin_context"
	"github.com/recative/recative-service-sdk/pkg/http_engine/http_err"
	"github.com/recative/recative-service-sdk/pkg/http_engine/response"
	"gorm.io/gorm"
)

type Controller interface {
	PostAppStorage(c *gin_context.Context[definition.JwtPayload])
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

func (con *controller) PostAppStorage(c *gin_context.Context[definition.JwtPayload]) {
	var body spec.PostAppStorageJSONRequestBody
	err := c.C.ShouldBindJSON(&body)
	if err != nil {
		response.Err(c.C, http_err.InvalidArgument.Wrap(err))
	}

	storages, err := con.service.ReadStoragesByKeysAndPermissions(body, c.Payload.Permissions)
	if err != nil {
		response.Err(c.C, http_err.InternalServerError.Wrap(err))
	}

	var res spec.StoragesResponse
	storage_format.StoragesToResponse(storages)
	response.Ok(c.C, res)
	return
}
