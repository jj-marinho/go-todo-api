package models

import "errors"

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

var TodoCounter int = 1
var Todos []Todo

func GetAllTodos() []Todo {
	return Todos
}

func CreateTodo(todo *Todo) {
	todo.Id = TodoCounter
	Todos = append(Todos, *todo)

	TodoCounter++
}

func GetTodo(id int) (*Todo, error) {
	for _, todo := range Todos {
		if todo.Id == id {
			return &todo, nil
		}
	}
	return nil, errors.New("Todo associated with id provided not found")
}

func UpdateTodo(id int, new_todo *Todo) error {
	todo, err := GetTodo(id)
	if err != nil {
		return err
	}

	todo.Title = new_todo.Title
	todo.Description = new_todo.Description
	todo.Done = new_todo.Done

	return nil
}

func DeleteTodo(id int) error {
	for i, todo := range Todos {
		if todo.Id == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			return nil
		}
	}

	return errors.New("Todo associated with id provided not found")
}
