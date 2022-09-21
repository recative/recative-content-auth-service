package db

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type databaseLogger struct {
	zap *zap.SugaredLogger

	LogLevel                  gormLogger.LogLevel
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

var _ gormLogger.Interface = &databaseLogger{}

type LoggerConfig struct {
	LogLevel                  gormLogger.LogLevel
	SlowThreshold             time.Duration
	IgnoreRecordNotFoundError bool
}

func NewDevelopmentGormLoggerConfig() LoggerConfig {
	return LoggerConfig{
		LogLevel:                  gormLogger.Info,
		SlowThreshold:             time.Second,
		IgnoreRecordNotFoundError: false,
	}
}

func NewProductionGormLoggerConfig() LoggerConfig {
	return LoggerConfig{
		LogLevel:                  gormLogger.Error,
		SlowThreshold:             3 * time.Second,
		IgnoreRecordNotFoundError: false,
	}
}

func (loggerConfig LoggerConfig) BuildWith(zap *zap.SugaredLogger) gormLogger.Interface {
	return NewGormLogger(zap, loggerConfig)
}

func NewGormLogger(zap *zap.SugaredLogger, databaseLoggerConfig LoggerConfig) gormLogger.Interface {
	return &databaseLogger{
		zap:                       zap,
		LogLevel:                  databaseLoggerConfig.LogLevel,
		SlowThreshold:             databaseLoggerConfig.SlowThreshold,
		IgnoreRecordNotFoundError: databaseLoggerConfig.IgnoreRecordNotFoundError,
	}
}

func (d *databaseLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := *d
	newLogger.LogLevel = level
	return &newLogger
}

func (d *databaseLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if d.LogLevel >= gormLogger.Info {
		d.zap.Info(append([]interface{}{zap.String("msg", msg), utils.FileWithLineNum()}, data...)...)
	}
}

func (d *databaseLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if d.LogLevel >= gormLogger.Warn {
		d.zap.Warn(append([]interface{}{zap.String("msg", msg), utils.FileWithLineNum()}, data...)...)
	}
}

func (d *databaseLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if d.LogLevel >= gormLogger.Error {
		d.zap.Error(append([]interface{}{zap.String("msg", msg), utils.FileWithLineNum()}, data...)...)
	}
}

func (d *databaseLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if d.LogLevel <= gormLogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && d.LogLevel >= gormLogger.Error && (!errors.Is(err, gorm.ErrRecordNotFound) || !d.IgnoreRecordNotFoundError):
		sql, rows := fc()
		if rows == -1 {
			d.zap.Error(utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			d.zap.Error(utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > d.SlowThreshold && d.SlowThreshold != 0 && d.LogLevel >= gormLogger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", d.SlowThreshold)
		if rows == -1 {
			d.zap.Warn("warn", utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			d.zap.Warn("warn", utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case d.LogLevel == gormLogger.Info:
		sql, rows := fc()
		if rows == -1 {
			d.zap.Info("trace", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			d.zap.Info("trace", utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
