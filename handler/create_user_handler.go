package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/schemas"
	"github.com/lidiagaldino/go-first-api/security"
	"github.com/lidiagaldino/go-first-api/utils"
)

// @BasePath /api/v1

// @Summary Create user
// @Description Create a user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User data to create"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user [post]
func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}

	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %v", err.Error())
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	hash, err := security.HashPassword(request.Password)
	if err != nil {
		logger.Errorf("error hashing password: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error hashing password")
		return
	}

	user := schemas.User{
		Login:    request.Login,
		Password: hash,
		Name:     request.Name,
	}

	if err := db.Create(&user).Error; err != nil {
		logger.Errorf("error creating user: %v", err)
		utils.SendError(ctx, http.StatusInternalServerError, "error creating user")
		return
	}

	response := UserResponse{
		ID:        user.ID,
		Login:     user.Login,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.SendSuccess(ctx, "create-user", response)
}
