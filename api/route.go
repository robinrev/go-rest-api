package api

import (
	"controller/todos_controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	routes := router.Group("/todos")
	{
		routes.GET("/", todos_controller.GetTodos)
		// routes.GET("/todos/:id", getTodo)
		// routes.PATCH("/todos/:id", toggleTodoStatus)
		// routes.POST("/todos", addTOdo)
		// routes.Run("localhost:9090")
	}
}
