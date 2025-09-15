package router

import (
	"github.com/huynh-fs/gin-api/internal/handler"
	"github.com/huynh-fs/gin-api/internal/middleware"
	"github.com/huynh-fs/gin-api/pkg/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/huynh-fs/gin-api/docs"
)

func Setup(cfg *config.Config, todoHandler *handler.TodoHandler, authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.New()

	r.Use(middleware.LoggerMiddleware())

	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")
	{
		authRoutes := apiV1.Group("/auth")
		{
			authRoutes.POST("/register", authHandler.Register)
			authRoutes.POST("/login", authHandler.Login)
			authRoutes.POST("/refresh", authHandler.RefreshToken)
			authRoutes.POST("/logout", authHandler.Logout)
		}
		
		todos := apiV1.Group("/todos")
		todos.Use(middleware.AuthMiddleware(cfg))
		{
			todos.POST("", todoHandler.CreateTodo)
			todos.GET("", todoHandler.ListTodos)
			todos.GET("/:id", todoHandler.GetTodo)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}