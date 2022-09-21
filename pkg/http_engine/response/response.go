package response

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/pkg/http_engine/http_err"
	"github.com/recative/recative-backend/pkg/logger"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func LogOkResponse(c *gin.Context, any any) {
	logger.Info("",
		zap.String("request_id", c.GetString("request_id")),
		zap.Int("user_id", c.GetInt("user_id")),
		zap.Any("response", any),
	)
}

func LogErrorResponse(c *gin.Context, err error) {
	logger.Error("",
		zap.String("request_id", c.GetString("request_id")),
		zap.Int("user_id", c.GetInt("user_id")),
		zap.Error(err),
	)
}

func Ok(c *gin.Context, data any) {
	defer func() {
		LogOkResponse(c, data)
	}()
	if data == nil {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, data)
}

func Err(c *gin.Context, err error) {
	//err = dberr.Wrap(err)
	//if dberr.DeadLockDetected.Is(err) || dberr.SerializationFailure.Is(err) || dberr.TransactionRollback.Is(err) || dberr.QueryCanceled.Is(err) {
	//	err = errs.ServiceUnavailable.Wrap(err)
	//}
	switch err.(type) {
	case *http_err.ResponseError:
		err := err.(*http_err.ResponseError)
		err.Id = c.GetString("request_id")
		c.JSON(err.ResponseStatusCode(), err)
		c.Abort()
		return
	default:
		requestId := c.GetString("request_id")

		err := http_err.UnexpectedInternalServerError.WrapAndSetId(err, requestId)

		if strings.HasPrefix(c.Request.URL.String(), "/app") {
			c.JSON(err.(*http_err.ResponseError).ResponseStatusCode(), err)
			LogErrorResponse(c, err)
			c.Abort()
			return
		}
		c.JSON(err.(*http_err.ResponseError).ResponseStatusCode(), err)
		LogErrorResponse(c, err)
		c.Abort()
		return
	}
}
