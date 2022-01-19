package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-basic-api/controllers"
	"github.com/go-basic-api/middlewares"
)

// Start all current routes available on API
func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	todo := router.Group("/todos")
	{
		todo.GET("", controllers.GetAllTodos)
		todo.POST("", controllers.CreateTodo)

		// Routes with id parameter need custom validation middleware
		idRoutes := todo.Group("/", middlewares.ValidateMongoId())
		{
			idRoutes.GET(":id", controllers.GetTodo)
			idRoutes.PUT(":id", controllers.UpdateTodo)
			idRoutes.DELETE(":id", controllers.DeleteTodo)
		}
	}

	return router
}
