package main

import (
	"github.com/enniome/go-todo/server"
	"github.com/labstack/echo"
)

func main() {
	// create server
	e := echo.New()

	// add routes
	registerRoutes(e)

	// start server
	e.Logger.Fatal(e.Start(":8080"))
}

func registerRoutes(e *echo.Echo) {
	e.GET("/info", server.GetInfo)
	e.POST("/", server.AddTodo)
	e.GET("/", server.GetTodos)
	e.GET("/:user", server.GetTodosForUser)
}
