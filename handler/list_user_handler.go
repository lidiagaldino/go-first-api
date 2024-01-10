package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/schemas"
)

// @BasePath /api/v1

// @Summary List users
// @Description List all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} ListUsersResponse
// @Failure 500 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /users [get]
func ListUsersHandler(ctx *gin.Context) {
	users := []schemas.User{}

	if err := db.Find(&users).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error listing users: %v", err))
		return
	}

	sendSuccess(ctx, "list-users", users)
}
