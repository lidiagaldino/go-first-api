package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/schemas"
	"github.com/lidiagaldino/go-first-api/utils"
)

// @BasePath /api/v1

// @Summary Update user
// @Description Update a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User Identification"
// @Param user body UpdateUserRequest true "User data to Update"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/{id} [put]
func UpdateUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	request := UpdateUserRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %s", err.Error())
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "param").Error())
		return
	}

	user := schemas.User{
		Login:    request.Login,
		Password: request.Password,
		Name:     request.Name,
	}

	if err := db.Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	utils.SendSuccess(ctx, "update-user", user)
}
