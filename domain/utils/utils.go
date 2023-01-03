package utils

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func SplitQueryParams(key string, ctx *gin.Context) []string {
	value := ctx.Query(key)
	if value == "" {
		return []string{}
	}
	return strings.Split(value, ",")
}
