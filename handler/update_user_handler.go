package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/schemas"
	"github.com/lidiagaldino/go-first-api/security"
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
// @Param Authorization header string true "With the bearer started"
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

	userForVerify := schemas.User{}

	//verify if login already exists
	if err := db.Where("login = ? AND id != ?", request.Login, id).First(&userForVerify).Error; err == nil {
		logger.Errorf("error creating user: %v", err)
		utils.SendError(ctx, http.StatusBadRequest, "login already exists")
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

	//verify if user exists
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		utils.SendError(ctx, http.StatusNotFound, "user not found")
		return
	}

	if err := db.Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, "error updating user")
		return
	}

	response := UserResponse{
		ID:        user.ID,
		Login:     user.Login,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	utils.SendSuccess(ctx, "update-user", response)
}
