package storage

import (
	"fmt"
	"github.com/enniome/go-todo/internal/model"
)

type TodoDB struct {
	items map[string][]model.Todo
}

func NewDB() TodoDB {
	return TodoDB{items: map[string][]model.Todo{}}
}

func (t TodoDB) AddTodo(user, todo string, done bool) {
	t.items[user] = append(t.items[user], model.Todo{
		Todo: todo,
		Done: done,
	})
	fmt.Printf("User %v has following todos: %v \n", user, t.items[user])
}

func (t TodoDB) GetTodos() *map[string][]model.Todo {
	return &t.items
}

func (t TodoDB) GetTodosFor(user string) []model.Todo {
	return t.items[user]
}
