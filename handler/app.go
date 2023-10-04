package handler

import (
	"final_project_1/docs"
	"final_project_1/infrastructure/config"
	"final_project_1/infrastructure/database"
	"final_project_1/repository/todo_repository/todo_repo"
	"final_project_1/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {

	config.LoadAppConfig()
	db := database.GetDatabaseInstance()

	// Dependency injection
	todoRepo := todo_repository.NewTodoRepo(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := NewTodoHandler(todoService)

	route := gin.Default()

	//Swagger
	docs.SwaggerInfo.Title = "Final Project 1 Kelompok 3"
	docs.SwaggerInfo.Description = "Todos API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:" + config.GetAppConfig().Port
	docs.SwaggerInfo.Schemes = []string{"http"}
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	todoRoute := route.Group("/todos")
	{
		todoRoute.GET("/", todoHandler.GetTodos)
		todoRoute.POST("/", todoHandler.CreateTodo)
		todoRoute.GET("/:todoId", todoHandler.GetTodo)
		todoRoute.PUT("/:todoId", todoHandler.UpdateTodo)
		todoRoute.DELETE("/:todoId", todoHandler.DeleteTodo)
	}

	route.Run(":" + config.GetAppConfig().Port)
}