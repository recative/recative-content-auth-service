package main

import (
	"github.com/recative/recative-backend/domain"
	"github.com/recative/recative-backend/mock_data"
	"github.com/recative/recative-service-sdk/pkg/auth"
	"github.com/recative/recative-service-sdk/pkg/config"
	"github.com/recative/recative-service-sdk/pkg/db"
	"github.com/recative/recative-service-sdk/pkg/http_engine"
	"github.com/spf13/viper"
	//"github.com/recative/recative-service-sdk/pkg"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(err)
	}

	var domainConfig domain.Config
	config.ForceParseByKey("domain", &domainConfig)

	var dbConfig db.Config
	config.ForceParseByKey("database", &dbConfig)
	db := db.New(dbConfig)

	isApiTest := viper.GetBool("database.is_api_test")
	if isApiTest && config.Environment() != config.Prod {
		mock_data.Init(db)
	}

	var httpEngineConfig http_engine.Config
	config.ForceParseByKey("http_engine", &httpEngineConfig)
	httpEngine := http_engine.Default(httpEngineConfig)
	httpEngine.UseRawPath = true
	httpEngine.AddPing()

	var authConfig auth.Config
	config.ForceParseByKey("auth", &authConfig)
	auther := auth.New(authConfig)

	domain.Init(&domain.Dependence{
		Db:         db,
		HttpEngine: httpEngine,
		Auther:     auther,
		DbConfig:   dbConfig,
	}, domainConfig)

	err = httpEngine.Run(httpEngineConfig.ListenAddr)
	if err != nil {
		panic(err)
	}
}
