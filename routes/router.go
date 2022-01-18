package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic-api/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	todo := router.Group("/todos")
	{
		todo.GET("", controllers.GetAllTodos)
		todo.POST("", controllers.CreateTodo)
		todo.GET("/:id", controllers.GetTodo)
		todo.PUT("/:id", controllers.UpdateTodo)
		todo.DELETE("/:id", controllers.DeleteTodo)
	}

	return router
}
