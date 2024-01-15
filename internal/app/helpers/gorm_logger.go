package helpers

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

type GormLogger struct {
	log                   *logrus.Logger
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Debug                 bool
}

func NewGormLogger(log *logrus.Logger) *GormLogger {
	return &GormLogger{
		log:                   log,
		SkipErrRecordNotFound: true,
		Debug:                 true,
	}
}

func (gormLogger *GormLogger) LogMode(logger.LogLevel) logger.Interface {
	return gormLogger
}

func (gormLogger *GormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	gormLogger.log.WithContext(ctx).Infof(s, args)
}

func (gormLogger *GormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	gormLogger.log.WithContext(ctx).Warnf(s, args)
}

func (gormLogger *GormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	gormLogger.log.WithContext(ctx).Errorf(s, args)
}

func (gormLogger *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := logrus.Fields{}

	if gormLogger.SourceField != "" {
		fields[gormLogger.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && gormLogger.SkipErrRecordNotFound) {
		fields[logrus.ErrorKey] = err
		gormLogger.log.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if gormLogger.SlowThreshold != 0 && elapsed > gormLogger.SlowThreshold {
		gormLogger.log.WithContext(ctx).WithFields(fields).Warnf("%s [%s]", sql, elapsed)
		return
	}

	if gormLogger.Debug {
		gormLogger.log.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
	}
}
