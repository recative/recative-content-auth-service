package main

import (
	"github.com/recative/recative-backend-sdk/pkg/auth"
	"github.com/recative/recative-backend-sdk/pkg/db"
	"github.com/recative/recative-backend-sdk/pkg/env"
	"github.com/recative/recative-backend-sdk/pkg/http_engine"
	"github.com/recative/recative-backend/domain"
	//"github.com/recative/recative-backend-sdk/pkg"
)

func main() {
	var dbConfig db.Config
	env.ForceParse(&dbConfig)
	db := db.New(dbConfig)

	var httpEngineConfig http_engine.Config
	env.ForceParse(&httpEngineConfig)
	httpEngine := http_engine.Default()

	var authConfig auth.Config
	env.ForceParse(&authConfig)
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
