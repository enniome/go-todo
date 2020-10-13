package server

import (
	"github.com/enniome/go-todo/internal/model"
	"github.com/enniome/go-todo/internal/storage"
	"github.com/labstack/echo"
	"net/http"
)

var db storage.TodoDB

func init() {
	db = storage.NewDB()
}

func GetInfo(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World, I'm running!")
}

func AddTodo(c echo.Context) (err error) {
	t := new(model.Todo)
	user := c.Request().Header.Get("User")
	if err = c.Bind(t); err != nil {
		return
	}
	db.AddTodo(user, t.Todo, t.Done)
	return c.JSON(http.StatusOK, t)
}

func GetTodos(c echo.Context) error {
	todos := db.GetTodos()
	return c.JSON(http.StatusOK, todos)
}

func GetTodosForUser(c echo.Context) error {
	param := c.Param("user")
	todos := db.GetTodosFor(param)
	return c.JSON(http.StatusOK, todos)
}
