package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializeSqlite()

	if err != nil {
		return fmt.Errorf("error initializing database: %v", err)
	}
	return nil
}

func GetSQLiteDatabase() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}