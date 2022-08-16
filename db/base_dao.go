package db

import (
	"go-start/core/logger"
	"gorm.io/gorm"
)

type BaseDao struct {
	Db  *gorm.DB
	Log *logger.Log
}
