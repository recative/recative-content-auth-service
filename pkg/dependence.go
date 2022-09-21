package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/recative/recative-backend/pkg/auth"
	"github.com/recative/recative-backend/pkg/cache"
	"github.com/recative/recative-backend/pkg/cronjob"
	"github.com/recative/recative-backend/pkg/cross_micro_service"
	"github.com/recative/recative-backend/pkg/db"
	"github.com/recative/recative-backend/pkg/env"
	"github.com/recative/recative-backend/pkg/gin_context"
	"github.com/recative/recative-backend/pkg/http_engine"
	"github.com/recative/recative-backend/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Dependence struct {
	Db               *gorm.DB
	DbConfig         db.Config
	HttpEngine       *gin.Engine
	HttpEngineConfig http_engine.Config
	Cache            cache.Cache
	Auther           auth.Authable
	AuthConfig       auth.Config
	Cron             cronjob.Cron
}

func (dep *Dependence) Start() {
	dep.Cron.Start()
	err := dep.HttpEngine.Run(dep.HttpEngineConfig.ListenAddr)
	if err != nil {
		logger.Fatal("http engine run failed", zap.Error(err))
	}
}

func AutoInit() *Dependence {
	var dbConfig db.Config
	env.Parse(&dbConfig)

	db := func() *gorm.DB {
		return db.New(dbConfig)
	}()

	httpEngine := http_engine.Default()
	var httpEngineConfig http_engine.Config
	env.Parse(&httpEngineConfig)

	cache := func() cache.Cache {
		var cacheConfig cache.Config
		env.Parse(&cacheConfig)
		return cache.New(cacheConfig)
	}()

	var authConfig auth.Config
	env.Parse(&authConfig)

	auther := func() auth.Authable {
		return auth.New(authConfig)
	}()

	var crossMicroServiceConfig cross_micro_service.CrossMicroServiceConfig
	env.Parse(&crossMicroServiceConfig)
	gin_context.Init(auther, crossMicroServiceConfig.HarmonyAuthorizationToken, nil)

	cron := func() cronjob.Cron {
		var cronConfig cronjob.Config
		env.Parse(&cronConfig)
		return cronjob.New(cronConfig)
	}()

	return &Dependence{
		Db:               db,
		DbConfig:         dbConfig,
		HttpEngine:       httpEngine,
		HttpEngineConfig: httpEngineConfig,
		Cache:            cache,
		Auther:           auther,
		AuthConfig:       authConfig,
		Cron:             cron,
	}
}
