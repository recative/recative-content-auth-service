package gin_context

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/recative/recative-backend/pkg/http_engine/http_err"
	"github.com/recative/recative-backend/pkg/http_engine/response"
	"github.com/recative/recative-backend/utils/array"
	"strings"
)

type Context[T any] struct {
	C         *gin.Context
	Languages []string
	Jwt       string
	// UserId 0 indicates current session is anonymous
	Payload T
}

func (c *Context[T]) Language() string {
	return c.Languages[0]
}

type HandlerFunc[T any] func(*Context[T])

func fillLanguage[T any](c *Context[T]) error {
	languageHeader := c.C.GetHeader(
		"X-Accept-Language")
	if strings.TrimSpace(languageHeader) == "" {
		languageHeader = c.C.GetHeader("Accept-Language")
	}
	c.Languages = MatchLanguage(languageHeader, []string{"zh-Hans", "zh-Hant", "en"})
	// use "en", "zh-Hans", "zh-Hant" as the fallback language
	c.Languages = append(c.Languages, "en", "zh-Hans", "zh-Hant")
	array.DistinctStringArray(c.Languages)
	return nil
}

func authN[T any](c *Context[T]) error {
	// get credential from header or query
	// TODO(Hexagram): why we need this?
	authHeader := c.C.Query("authorization_header")
	if len(authHeader) == 0 {
		authHeader = c.C.GetHeader("Authorization")
	}
	// authenticate
	jwtMapClaims, err := _auther.ParseJwt(ParseAuthorizationBearerHeader(authHeader))
	if err != nil {
		return http_err.Unauthorized.Wrap(err)
	}

	var t T
	err = mapstructure.Decode(jwtMapClaims, &t)
	if err != nil {
		return http_err.Unauthorized.Wrap(err)
	}

	c.Payload = t
	c.Jwt = authHeader

	err = customLogic(jwtMapClaims, c.C)
	if err != nil {
		return err
	}

	//_, securityRequired := c.C.Get("security_required")
	//must.True(err == nil || userId == 0)
	//if (userId == 0 || err != nil) && securityRequired {
	//	if err != nil {
	//		return err
	//	} else {
	//		return http_err.Unauthorized.New()
	//	}
	//}
	//c.C.Set("user_id", userId)
	//c.UserId = userId

	return nil
}

func Handler[T any](handler HandlerFunc[T]) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context[T])
		context.C = c
		if e := fillLanguage(context); e != nil {
			response.Err(c, e)
			return
		}

		if e := authN(context); e != nil {
			response.Err(c, e)
			return
		}

		handler(context)
	}
}

// ParseAuthorizationBearerHeader Parse HTTP Authorization header with Bearer credentials.
// Return an empty string when parsing fails.
//
// See also: https://datatracker.ietf.org/doc/html/rfc6750#section-2.1
func ParseAuthorizationBearerHeader(authHeader string) string {
	h := strings.TrimSpace(authHeader)
	if strings.HasPrefix(h, "Bearer ") {
		return strings.TrimPrefix(h, "Bearer ")
	}
	return ""
}
