package db

import (
	"database/sql"
	"github.com/recative/recative-backend/pkg/env"
	"github.com/recative/recative-backend/pkg/logger"
	"go.uber.org/zap"
	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type Transactionable interface {
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
}

type AutoMigrater interface {
	AutoMigrate()
}

type Config struct {
	IsAutoMigrate bool   `env:"IS_AUTO_MIGRATE" envDefault:"true"`
	PgsqlUri      string `env:"PGSQL_URI"`
	PgsqlMaxIdle  int    `env:"PGSQL_MAX_IDLE" envDefault:"50"`
	PgsqlMaxOpen  int    `env:"PGSQL_MAX_OPEN" envDefault:"50"`
}

func New(config Config) *gorm.DB {
	var gormLogger gormlogger.Interface

	if env.Environment() == env.Prod {
		gormLogger = NewDevelopmentGormLoggerConfig().BuildWith(logger.RawLogger().Sugar())
	}

	db, err := gorm.Open(driver.Open(config.PgsqlUri), &gorm.Config{Logger: gormLogger})
	if err != nil {
		panic("Open PostgreSQL DB failed: " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Panic("init database failed", zap.Error(err))
	}
	sqlDB.SetMaxIdleConns(config.PgsqlMaxIdle)
	sqlDB.SetMaxOpenConns(config.PgsqlMaxOpen)

	return db
}
