package logger

import (
	"golang.org/x/net/context"
	"gorm.io/gorm/logger"
	"time"
)

type Log struct {
	logger logger.Interface
}

func (log *Log) Info(str string, data ...interface{}) {
	log.logger.Info(context.Background(), str, data)
}

func (log *Log) Warn(str string, data ...interface{}) {
	log.logger.Warn(context.Background(), str, data)
}

func (log *Log) Error(str string, data ...interface{}) {
	log.logger.Error(context.Background(), str, data)
}

func (log *Log) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	log.logger.Trace(ctx, begin, fc, err)
}

func NewLog(logger logger.Interface) *Log {
	log := Log{logger: logger}
	return &log
}
