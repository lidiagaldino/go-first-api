package router

import (
	"github.com/gin-gonic/gin"
	docs "github.com/lidiagaldino/go-first-api/docs"
	"github.com/lidiagaldino/go-first-api/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handler.Init()
	v1 := router.Group("/api/v1/")
	docs.SwaggerInfo.BasePath = "/api/v1"
	{
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
