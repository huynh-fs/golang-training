package router

import (
	"github.com/huynh-fs/gin-api/internal/handler"
	"github.com/huynh-fs/gin-api/internal/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/huynh-fs/gin-api/docs"
)

func Setup(todoHandler *handler.TodoHandler) *gin.Engine {
	r := gin.New()

	r.Use(middleware.LoggerMiddleware())

	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")
	{
		todos := apiV1.Group("/todos")
		todos.Use(middleware.AuthMiddleware())
		{
			todos.POST("", todoHandler.CreateTodo)
			todos.GET("", todoHandler.GetTodos)
			todos.GET("/:id", todoHandler.GetTodo)
			todos.PUT("/:id", todoHandler.UpdateTodo)
			todos.DELETE("/:id", todoHandler.DeleteTodo)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}