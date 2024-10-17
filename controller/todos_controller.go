package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func ToggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "to do not found"})
		return
	}

	todo.Completed = !todo.Completed

	context.IndentedJSON(http.StatusOK, todo)
}

func GetTodoById(id string) (*Todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("todo not found")
}

func AddTOdo(context *gin.Context) {
	var newTodo Todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := GetTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "to do not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}
