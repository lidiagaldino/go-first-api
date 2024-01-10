package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/authentication"
	"github.com/lidiagaldino/go-first-api/schemas"
	"github.com/lidiagaldino/go-first-api/security"
	"github.com/lidiagaldino/go-first-api/utils"
)

// @BasePath /api/v1

// @Summary Login
// @Description Login
// @Tags Users
// @Accept json
// @Produce json
// @Param user body LoginRequest true "User data to login"
// @Success 200 {object} LoginResponse
// @Failure 401 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func LoginHandler(ctx *gin.Context) {
	request := LoginRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{
		Login: request.Login,
	}
	if err := db.Model(&user).Where("login = ? ", user.Login).First(&user).Error; err != nil {
		utils.SendError(ctx, http.StatusUnauthorized, "unauthorized user")
		return
	}

	if err := security.CheckPasswordHash(request.Password, user.Password); err {
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
