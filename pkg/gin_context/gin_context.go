package gin_context

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/recative/recative-backend/pkg/auth"
)

var _auther auth.Authable
var authorizationToken string

type CustomLogic = func(claims jwt.Claims, c *gin.Context) error

var customLogic CustomLogic

func Init(auther auth.Authable, authorizationToken_ string, _customLogic CustomLogic) {
	_auther = auther
	authorizationToken = authorizationToken_
	customLogic = _customLogic
}
