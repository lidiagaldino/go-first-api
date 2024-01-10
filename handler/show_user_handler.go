package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/schemas"
	"github.com/lidiagaldino/go-first-api/utils"
)

// @BasePath /api/v1

// @Summary Show user
// @Description Show a user
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User identification"
// @Success 200 {object} ShowUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /user/{id} [get]
func ShowUserHandler(ctx *gin.Context) {
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

	utils.SendSuccess(ctx, "show-user", user)
}
