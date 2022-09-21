package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/recative/recative-backend/pkg/env"
	"github.com/recative/recative-backend/pkg/logger"
	"github.com/recative/recative-backend/utils/must"
	"go.uber.org/zap"
	"io"
	"strconv"
	"time"
)

func LogRequestBase(c *gin.Context) []zap.Field {
	return []zap.Field{
		zap.String("client_ip", c.ClientIP()),
		zap.String("client_port", c.Request.Host),
		zap.String("method", c.Request.Method),
		zap.String("request_id", c.GetString("request_id")),
		zap.String("uri", c.Request.RequestURI),
		zap.Any("headers", c.Request.Header),
		zap.String("hostname", c.Request.Host),
	}
}

func LogProdRequest(c *gin.Context) {
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}

	logger.Info("",
		LogRequestBase(c)...,
	)
}

// LogDevRequestCreator Use closure to get user_id from *gin.Context
func LogDevRequestCreator(c *gin.Context) func(ctx *gin.Context) {
	var requestBody string

	bodyByte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		requestBody = err.Error()
	} else {
		requestBody = string(bodyByte)
	}

	// WARN: Write it back! Read is one time!
	// Guard just in Dev mode.
	c.Request.Body = io.NopCloser(bytes.NewReader(bodyByte))

	return func(c *gin.Context) {
		logger.Info("",
			append(
				LogRequestBase(c),
				zap.String("request_body", requestBody),
			)...,
		)
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()
		c.Set("request_id", requestID)

		// magic, WARN!
		if env.Environment() != env.Prod {
			log := LogDevRequestCreator(c)
			c.Next()
			log(c)
		} else {
			LogProdRequest(c)
			c.Next()
		}

		l := float64(time.Since(start)) / float64(time.Millisecond)

		latency, err := strconv.ParseFloat(fmt.Sprintf("%.3f", l), 64)
		must.Must(err)

		logger.Info("",
			zap.String("request_id", c.GetString("request_id")),
			zap.Int64("timestamp", time.Now().Unix()),
			zap.String("remote_addr", c.ClientIP()),
			zap.String("http_referer", c.Request.Referer()),
			zap.String("request", fmt.Sprintf("%s %s %s", c.Request.Method, c.Request.RequestURI, c.Request.Proto)),
			zap.Int("status", c.Writer.Status()),
			zap.Int("body_bytes_sent", c.Writer.Size()),
			zap.String("http_user_agent", c.Request.UserAgent()),
			zap.Float64("latency", latency),
		)
	}
}
