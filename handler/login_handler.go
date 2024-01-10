package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/schemas"
)

func LoginHandler(ctx *gin.Context) {
	request := LoginRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Login:    request.Login,
		Password: request.Password,
	}
	if err := db.Model(&user).Where("login = ? AND password = ?", user.Login, user.Password).First(&user).Error; err != nil {
		sendError(ctx, http.StatusUnauthorized, "unauthorized user")
		return
	}

	token, err := createToken.CreateToken(user.Login)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error creating token")
		return
	}

	sendSuccess(ctx, "login", LoginResponse{Token: token})
}
