package models

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/7ruedzn/todos/internal/files"
)

type Todo struct {
	Id          int
	Description string
	CreatedAt   time.Time
	Done        bool
}

func New(description string) *Todo {
	return &Todo{
		// Id:          len(todos) + 1,
		Description: description,
		CreatedAt:   time.Now(),
		Done:        false,
	}
}

func GetTodo(id int, todos []Todo) (*Todo, error) {
	for _, v := range todos {
		if id == v.Id {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("todo with id %d not found\n", id)
}

func GetTodos() *[]Todo {
	b, err := files.Load()
	if err != nil {
		if errFile := files.Create([]byte("{}")); errFile != nil {
			panic(errFile)
		}
	}

	todos := []Todo{}
	if err := json.Unmarshal(b, &todos); err != nil {
		panic(err)
	}

	return &todos
}

func DeleteTodo(id int) error {
	todos := *GetTodos()

	if id == 0 || id > len(todos) {
		fmt.Fprintf(os.Stderr, "the todo with id %d doesnt exists!\n", id)
	}

	todos = slices.Delete(todos, (id - 1), id)

	//INFO: range returns a COPY of the element, not the pointer.
	// this way you access directly the todo reference
	for i := 0; i < len(todos); i++ {
		todos[i].Id = i + 1
	}

	b, err := json.Marshal(&todos)

	if err != nil {
		return err
	}

	err = files.Write(b)

	if err != nil {
		return err
	}

	return nil
}

func AddTodo(todos []Todo, description string) []Todo {
	todo := Todo{
		Id:          len(todos) + 1,
		Description: description,
		CreatedAt:   time.Now(),
		Done:        false,
	}

	todos = append(todos, todo)

	return todos
}

func (todo *Todo) UpdateTodos() (*[]Todo, error) {
	if todo.Done == true {
		return nil, fmt.Errorf("the todo with id %d is already complete\n", todo.Id)
	}

	todos := *GetTodos() //TODO: load the file from app
	for i := 0; i < len(todos); i++ {
		if todo.Id == todos[i].Id {
			todos[i].Done = true
			return &todos, nil
		}
	}

	return nil, fmt.Errorf("couldnt update the todo with id %q\n", todo.Id)
}
