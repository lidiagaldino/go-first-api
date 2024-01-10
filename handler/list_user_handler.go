package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/schemas"
	"github.com/lidiagaldino/go-first-api/utils"
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
		utils.SendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error listing users: %v", err))
		return
	}

	response := []UserResponse{}
	for _, user := range users {
		response = append(response, UserResponse{
			ID:        user.ID,
			Login:     user.Login,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	utils.SendSuccess(ctx, "list-users", response)
}
