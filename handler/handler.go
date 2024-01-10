package handler

import (
	"github.com/lidiagaldino/go-first-api/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSQLiteDatabase()
}
