package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/schemas"
	"github.com/lidiagaldino/go-first-api/utils"
)

// @BasePath /api/v1

// @Summary Delete user
// @Description Delete a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User Identification"
// @Success 200 {object} DeleteUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/{id} [delete]
func DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "param").Error())
		return
	}
	user := schemas.User{}

	if err := db.First(&user, id).Error; err != nil {
		utils.SendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id %s not found", id))
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting user with id %s", id))
		return
	}

	utils.SendSuccess(ctx, "delete-user", user)
}
