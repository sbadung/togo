package todo

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := "/api/v1"
	todos := router.Group(api + "/todos")
	{
		todos.GET("/", GetTodos)
		todos.GET("/:id", GetTodo)
		todos.POST("/", CreateTodo)
		todos.PUT("/:id", UpdateTodo)
		todos.DELETE("/:id", DeleteTodo)
		todos.DELETE("/", DeleteAllTodos)
	}
}
