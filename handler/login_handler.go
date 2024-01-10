package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/authentication"
	"github.com/lidiagaldino/go-first-api/schemas"
	"github.com/lidiagaldino/go-first-api/utils"
)

func LoginHandler(ctx *gin.Context) {
	request := LoginRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Login:    request.Login,
		Password: request.Password,
	}
	if err := db.Model(&user).Where("login = ? AND password = ?", user.Login, user.Password).First(&user).Error; err != nil {
		utils.SendError(ctx, http.StatusUnauthorized, "unauthorized user")
		return
	}

	token, err := authentication.CreateToken(user.Login)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error creating token")
		return
	}

	utils.SendSuccess(ctx, "login", LoginResponse{Token: token})
}
