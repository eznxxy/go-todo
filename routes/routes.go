package routes

import (
	"github.com/eznxxy/go-todo/controllers"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/todos", controllers.FetchAllTodo)
	e.POST("/todos", controllers.CraeteTodo)
	e.PUT("/todos/:id", controllers.UpdateTodo)
	e.DELETE("/todos/:id", controllers.DeleteTodo)
	e.PATCH("/todos/:id/mark", controllers.MarkTodo)

	return e
}
