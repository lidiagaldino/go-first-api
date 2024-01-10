package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfully", op), "data": data,
	})
}
