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
	basePath := "/api/v1"
	v1 := router.Group(basePath)
	docs.SwaggerInfo.BasePath = basePath
	{
		v1.GET("/opening/:id", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening/:id", handler.DeleteOpeningHandler)
		v1.PUT("/opening/:id", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)

		v1.GET("/user/:id", handler.ShowUserHandler)
		v1.POST("/user", handler.CreateUserHandler)
		v1.DELETE("/user/:id", handler.DeleteUserHandler)
		v1.PUT("/user/:id", handler.UpdateUserHandler)
		v1.GET("/users", handler.ListUsersHandler)

		v1.POST("/login", handler.LoginHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
