package gin_context

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/pkg/http_engine/http_err"
	"github.com/recative/recative-backend/pkg/http_engine/response"
	"strconv"
)

type InternalContext struct {
	UserId                     int
	InternalAuthorizationToken string
	C                          *gin.Context
}

type InternalHandlerFunc func(ctx *InternalContext)

func InternalHandler(handler InternalHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		internalContext := new(InternalContext)
		internalContext.C = c

		if err := checkInternalAuthorization(internalContext); err != nil {
			response.Err(c, err)
			return
		}
		if err := fillUserId(internalContext); err != nil {
			response.Err(c, err)
			return
		}

		handler(internalContext)
	}
}

func checkInternalAuthorization(c *InternalContext) error {
	internalAuthorizationToken := c.C.GetHeader("X-InternalAuthorization")

	switch internalAuthorizationToken {
	case authorizationToken:
		c.InternalAuthorizationToken = internalAuthorizationToken
		return nil
	default:
		return http_err.Unauthorized.New("invalid internal authorization token")
	}
}

func fillUserId(c *InternalContext) error {
	userIdString := c.C.GetHeader("X-UserId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		return http_err.Unauthorized.Wrap(err)
	}
	c.UserId = userId
	return nil
}
