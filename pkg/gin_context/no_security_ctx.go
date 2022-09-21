package gin_context

import (
	"github.com/gin-gonic/gin"
)

type NoSecurityContext struct {
	C *gin.Context
}

type NoSecurityHandlerFunc func(ctx *NoSecurityContext)

func NoSecurityHandler(handler NoSecurityHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		internalContext := new(NoSecurityContext)
		internalContext.C = c

		handler(internalContext)
	}
}
