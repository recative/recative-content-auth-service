package auth

import (
	"github.com/golang-jwt/jwt"
	"github.com/recative/recative-backend/pkg/http_engine/http_err"
	"github.com/recative/recative-backend/pkg/logger"
	"go.uber.org/zap"
)

type Authable interface {
	GenJwt(mapClaims jwt.MapClaims) string
	ParseJwt(tokenStr string) (jwt.MapClaims, error)
}

type authable struct {
	JwtSecret string
}

type Config struct {
	JwtSecret string `env:"JWT_SECRET"`
}

func New(config Config) Authable {
	return authable{
		JwtSecret: config.JwtSecret,
	}
}

func (a authable) GenJwt(mapClaims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	tokenString, err := token.SignedString([]byte(a.JwtSecret))
	if err != nil {
		logger.Panic("error when sign jwt", zap.Error(err))
	}
	return tokenString
}

// ParseJwt returns user id or error if any occurs.
func (a authable) ParseJwt(tokenStr string) (jwt.MapClaims, error) {
	if tokenStr == "" {
		return nil, http_err.Unauthorized.New("invalid token header")
	}
	token, err := jwt.Parse(tokenStr,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(a.JwtSecret), nil
		})
	if err != nil {
		return nil, http_err.Unauthorized.New("invalid JWT: parsing failed" + err.Error())
	}

	return token.Claims.(jwt.MapClaims), nil

	//// get user id
	//v, ok := payload["user_id"]
	//if !ok {
	//	return 0, "", http_err.Unauthorized.New("invalid JWT: user_id in claims is required")
	//}
	//userId, ok := v.(float64)
	//if !ok || userId <= 0 {
	//	return 0, "", http_err.Unauthorized.New("invalid JWT: user_id invalid")
	//}
	//
	//// get platform
	//IPlatform, ok := payload["platform"]
	//if !ok {
	//	platform = "unknown"
	//}
	//platform, ok = IPlatform.(string)
	//if !ok {
	//	platform = "unknown"
	//}
	//
	//return int(userId), platform, nil
}
