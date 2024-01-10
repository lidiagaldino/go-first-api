package handler

import (
	"github.com/lidiagaldino/go-first-api/authentication"
	"github.com/lidiagaldino/go-first-api/config"
	"gorm.io/gorm"
)

var (
	logger      *config.Logger
	createToken *authentication.CreateToken
	verifyToken *authentication.VerifyToken
	db          *gorm.DB
)

func Init() {
	logger = config.GetLogger("handler")
	db = config.GetSQLiteDatabase()
	createToken = authentication.GetCreateToken()
	verifyToken = authentication.GetVerifyToken()
}
