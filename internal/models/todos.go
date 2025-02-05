package models

import (
	"encoding/json"
	"fmt"
	"slices"
	"time"

	"github.com/7ruedzn/todos/internal/files"
	"github.com/spf13/viper"
)

type Todo struct {
	Id          int
	Description string
	CreatedAt   time.Time
	Done        bool
}

func GetTodo(id int, todos []Todo) (*Todo, error) {
	for _, v := range todos {
		if id == v.Id {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("todo with id %d not found\n", id)
}

func GetTodos() []Todo {
	todos := []Todo{}
	todosPath := viper.GetString("todos.path")
	if todosPath == "" {
		fmt.Println("Todos path is empty: ", todosPath)
	}
	b, err := files.Load(todosPath)
	if err != nil {
		if err := files.Create(todosPath); err != nil {
			fmt.Println("err creating file on get todos", err)
			panic(err)
		}
		return todos
	}

	b, err = files.Load(todosPath)

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(b, &todos); err != nil && len(b) > 0 {
		panic(err)
	}

	return todos
}

func DeleteTodo(id int) error {
	todosPath := viper.GetString("todos.path")
	todos := GetTodos()

	if id == 0 || id > len(todos) {
		return fmt.Errorf("the todo with id %d doesnt exists! See usage with %q or use %q to list your todos!\n", id, "help", "todos list -a")
	}

	fmt.Println("todos before delete: ", todos)
	todos = slices.Delete(todos, (id - 1), id)
	fmt.Println("todos after delete: ", todos)

	//INFO: range returns a COPY of the element, not the pointer.
	// this way you access directly the todo reference
	for i := 0; i < len(todos); i++ {
		todos[i].Id = i + 1
	}

	fmt.Println("todos after id re assign: ", todos)

	b, err := json.Marshal(&todos)

	if err != nil {
		return err
	}

	if err := files.Write(b, todosPath); err != nil {
		return err
	}

	return nil
}

func AddTodo(todos []Todo, description string) ([]Todo, Todo) {
	todo := Todo{
		Id:          len(todos) + 1,
		Description: description,
		CreatedAt:   time.Now(),
		Done:        false,
	}

	todos = append(todos, todo)

	return todos, todo
}

func (todo *Todo) UpdateTodos() (*[]Todo, error) {
	if todo.Done == true {
		return nil, fmt.Errorf("the todo with id %d is already complete\n", todo.Id)
	}

	todos := GetTodos() //TODO: load the file from app
	for i := 0; i < len(todos); i++ {
		if todo.Id == todos[i].Id {
			todos[i].Done = true
			return &todos, nil
		}
	}

	return nil, fmt.Errorf("couldnt update the todo with id %q\n", todo.Id)
}
