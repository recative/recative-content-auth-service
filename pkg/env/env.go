package env

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"os"
	"reflect"
	"strings"
)

type EnvConfig struct {
	RuntimeEnvironmentConfig RuntimeEnvironmentConfig
	DatabaseConfig           DatabaseConfig
	ServerConfig             ServerConfig
	CacheConfig              CacheConfig
	LogConfig                LogConfig
	AliyunConfig             AliyunConfig
	JwtConfig                JwtConfig
	WeChatConfig             WeChatConfig
	AlipayConfig             AlipayConfig
	PolyvConfig              PolyvConfig
	PaymentConfig            PaymentConfig
	CrossMicroServiceConfig  CrossMicroServiceConfig
	OauthConfig              OauthConfig
}

type EnvironmentType string

const (
	Debug EnvironmentType = "DEBUG"
	Test  EnvironmentType = "TEST"
	Prod  EnvironmentType = "PROD"
)

type EnvironmentConfig struct {
	Environment EnvironmentType `env:"ENVIRONMENT"`
}

func Parse(structPointer any) {
	err := env.ParseWithFuncs(structPointer, map[reflect.Type]env.ParserFunc{
		reflect.TypeOf("string"): func(s string) (interface{}, error) {
			// This is fit docker-compose env file parse which will save \"xxx\" in .env file
			s = strings.Trim(s, `"`)
			return s, nil
		},
	}, env.Options{
		RequiredIfNoDef: true,
	})
	if err != nil {
		panic(err)
	}
}

func Environment() EnvironmentType {
	return environmentConfig.Environment
}

var environmentConfig EnvironmentConfig

func init() {
	if os.Getenv("ENVIRONMENT") != string(Prod) {
		err := godotenv.Load(".env")
		if err != nil {
			panic(fmt.Errorf(
				"error loading .env file: %w when ENVIRONMENT is %s",
				err,
				os.Getenv("ENVIRONMENT"),
			))
		}
	}

	Parse(&environmentConfig)

	if environmentConfig.Environment != Debug &&
		environmentConfig.Environment != Test &&
		environmentConfig.Environment != Prod {
		panic("unknown ENVIRONMENT: " + environmentConfig.Environment)
	}
}

type RuntimeEnvironmentConfig struct {
	GoEnv string `env:"ENVIRONMENT"`
}

type DatabaseConfig struct {
	IsAutoMigrate bool   `env:"IS_AUTO_MIGRATE" envDefault:"true"`
	PgsqlUri      string `env:"PGSQL_URI"`
	PgsqlMaxIdle  int    `env:"PGSQL_MAX_IDLE" envDefault:"50"`
	PgsqlMaxOpen  int    `env:"PGSQL_MAX_OPEN" envDefault:"50"`
}

type CacheConfig struct {
	RedisUri string `env:"REDIS_URI"`
}

type ServerConfig struct {
	ServerHost string `env:"SERVER_HOST"`
	ListenAddr string `env:"LISTEN_ADDR"`
}

type LogConfig struct {
	LogDir string `env:"LOG_DIR" envDefault:"/app"`
}

type AliyunConfig struct {
	AliAccessKeyID        string `env:"ALI_ACCESS_KEY_ID"`
	AliAccessKeySecret    string `env:"ALI_ACCESS_KEY_SECRET"`
	AliOssAccessKeyID     string `env:"ALI_OSS_ACCESS_KEY_ID"`
	AliOssAccessKeySecret string `env:"ALI_OSS_ACCESS_KEY_SECRET"`
}

type JwtConfig struct {
	JwtSecret        string `env:"JWT_SECRET" envDefault:"azm47twl0e1sy96ic2xr3gko1nud5v"`
	AdminJwtSecret   string `env:"ADMIN_JWT_SECRET" envDefault:"ak9wjd5nkssefseteonoqwpsea5jfkjbd64"`
	AdminJwtIssuer   string `env:"ADMIN_JWT_ISSUER" envDefault:"admin_site"`
	AdminJwtAudience string `env:"ADMIN_JWT_AUDIENCE" envDefault:"admin_site"`
}

type WeChatConfig struct {
	WechatSerialNo       string `env:"WECHAT_SERIAL_NO"`
	WechatMchID          string `env:"WECHAT_MCH_ID"`
	WechatApiV3Key       string `env:"WECHAT_API_V3_KEY"`
	WechatPrivateKey     string `env:"WECHAT_PRIVATE_KEY"`
	WechatOAAppID        string `env:"WECHAT_OA_APP_ID"` // 公众号 AppID
	WechatOASecret       string `env:"WECHAT_OA_SECRET"`
	WechatOAToken        string `env:"WECHAT_OA_TOKEN"`
	WechatWebAppID       string `env:"WECHAT_WEB_APP_ID"` // 网站应用 AppID
	WechatWebSecret      string `env:"WECHAT_WEB_SECRET"`
	WechatAppAppID       string `env:"WECHAT_APP_APP_ID"` // 移动app AppID
	WechatAppSecret      string `env:"WECHAT_APP_SECRET"`
	WechatEncodingAESKey string `env:"WECHAT_ENCODING_AES_KEY"`
}

type AlipayConfig struct {
	AlipayAppID      string `env:"ALIPAY_APP_ID"`
	AlipayPublicKey  string `env:"ALIPAY_PUBLIC_KEY"`
	AlipayPrivateKey string `env:"ALIPAY_PRIVATE_KEY"`
}

type PolyvConfig struct {
	PolyvUserId string `env:"POLYV_USER_ID"`
	PolyvSecret string `env:"POLYV_SECRET"`
}

type PaymentConfig struct {
	StripeSecretKey string `env:"STRIPE_SECRET_KEY"`
	// 前端处理成功的URL，可以使用参数
	StripePayHandleSuccessUrl string `env:"STRIPE_PAY_HANDLE_SUCCESS_URL"`
	// 前端处理取消的URL，可以使用参数
	StripePayHandleCancelUrl string `env:"STRIPE_PAY_HANDLE_CANCEL_URL"`
	// webhook secert
	StripeEndpointSecret string `env:"STRIPE_ENDPOINT_SECRET"`

	GooglePlayStoreKey string `env:"GOOGLE_PLAY_STORE_KEY"`

	// The app-specific shared secret is a unique code to receive receipts for only this app’s auto-renewable subscriptions.
	AppSpecificSharedSecret string `env:"APP_SPECIFIC_SHARED_SECRET"`
}

type CrossMicroServiceConfig struct {
	HarmonyAuthorizationToken string `env:"HARMONY_AUTHORIZATION_TOKEN"`
	HarmonyBaseUrl            string `env:"HARMONY_BASE_URL"`
}

type OauthConfig struct {
	TwitterKey           string `env:"TWITTER_KEY"`
	TwitterSecret        string `env:"TWITTER_SECRET"`
	TwitterCallback      string `env:"TWITTER_CALLBACK"`
	GithubKey            string `env:"GITHUB_KEY"`
	GithubSecret         string `env:"GITHUB_SECRET"`
	GithubCallback       string `env:"GITHUB_CALLBACK"`
	FacebookKey          string `env:"FACEBOOK_KEY"`
	FacebookSecret       string `env:"FACEBOOK_SECRET"`
	FacebookCallback     string `env:"FACEBOOK_CALLBACK"`
	AppleKey             string `env:"APPLE_KEY"`
	AppleSecret          string `env:"APPLE_SECRET"`
	AppleCallback        string `env:"APPLE_CALLBACK"`
	GoogleKey            string `env:"GOOGLE_KEY"`
	GoogleSecret         string `env:"GOOGLE_SECRET"`
	GoogleCallback       string `env:"GOOGLE_CALLBACK"`
	OauthFilterRedirect  string `env:"OAUTH_FILTER_REDIRECT"`
	OauthSuccessRedirect string `env:"OAUTH_SUCCESS_REDIRECT"`
	OauthFailRedirect    string `env:"OAUTH_FAIL_REDIRECT"`
}

var Config EnvConfig

//func Init() {
//	err := godotenv.Load(".env")
//	if err != nil && os.Getenv("ENVIRONMENT") != Prod {
//		log.Println("Error loading .env file", zap.Error(err))
//	}
//
//	var config EnvConfig
//	err = env.ParseWithFuncs(&config, map[reflect.Type]env.ParserFunc{
//		reflect.TypeOf("string"): func(s string) (interface{}, error) {
//			// This is fit docker-compose env file parse which will save \"xxx\" in .env fil
//			s = strings.TrimPrefix(s, `"`)
//			s = strings.TrimSuffix(s, `"`)
//			return s, nil
//		},
//	}, env.Options{
//		RequiredIfNoDef: true,
//	})
//	if err != nil {
//		panic(err)
//	}
//	if config.RuntimeEnvironmentConfig.GoEnv != Debug && config.RuntimeEnvironmentConfig.GoEnv != Test && config.RuntimeEnvironmentConfig.GoEnv != Prod {
//		panic("wrong ENVIRONMENT")
//	}
//
//	Config = config
//}
