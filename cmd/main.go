package main

import (
	"github.com/recative/recative-backend-sdk/pkg/auth"
	"github.com/recative/recative-backend-sdk/pkg/config"
	"github.com/recative/recative-backend-sdk/pkg/db"
	"github.com/recative/recative-backend-sdk/pkg/http_engine"
	"github.com/recative/recative-backend/domain"
	//"github.com/recative/recative-backend-sdk/pkg"
)

func main() {
	var dbConfig db.Config
	config.ForceParseByKey("database", &dbConfig)
	db := db.New(dbConfig)

	var httpEngineConfig http_engine.Config
	config.ForceParseByKey("http_engine", &httpEngineConfig)
	httpEngine := http_engine.Default(httpEngineConfig)

	var authConfig auth.Config
	config.ForceParseByKey("auth", &authConfig)
	auther := auth.New(authConfig)

	domain.Init(&domain.Dependence{
		Db:         db,
		HttpEngine: httpEngine,
		Auther:     auther,
	})

	err := httpEngine.Run(httpEngineConfig.ListenAddr)
	if err != nil {
		panic(err)
	}
}
