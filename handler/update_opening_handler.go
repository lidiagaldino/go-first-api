package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lidiagaldino/go-first-api/schemas"
)

func UpdateOpeningHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	request := UpdateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("error validating request: %s", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}	

	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "param").Error())
		return
	}

	opening := schemas.Opening{
		Role: request.Role,
		Company: request.Company,
		Location: request.Location,
		Remote: *request.Remote,
		Link: request.Link,
		Salary: request.Salary,
	}

	if err := db.Model(&opening).Where("id = ?", id).Updates(opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}

	sendSuccess(ctx, "update-opening", opening)
}