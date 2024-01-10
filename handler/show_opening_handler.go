package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/schemas"
	"github.com/lidiagaldino/go-first-api/utils"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path int true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening/{id} [get]
func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		utils.SendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "param").Error())
		return
	}
	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		utils.SendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id %s not found", id))
		return
	}

	utils.SendSuccess(ctx, "show-opening", opening)
}
