package output

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/7ruedzn/todos/models"
	"github.com/dustin/go-humanize"
)

func ListTodos(todos []models.Todo, all bool) {
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)

	if len(todos) <= 0 {
		fmt.Println("You have no todos added! See usage with help or -h to add your todos!")
		return
	}

	for i, v := range todos {
		if all {
			if i == 0 {
				fmt.Fprintf(w, "ID\tTask\tCreated\tDone\n")
			}
			fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", v.Id, v.Description, humanize.Time(v.CreatedAt), v.Done)
			continue
		} else {
			if i == 0 {
				fmt.Fprintf(w, "ID\tTask\tCreated\n")
			}
			if !v.Done {
				fmt.Fprintf(w, "%d\t%s\t%s\n", v.Id, v.Description, humanize.Time(v.CreatedAt))
				continue
			}
		}
	}

	w.Flush()
}

func ListAddedTodo(todo models.Todo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 2, 4, ' ', 0)

	fmt.Fprintf(w, "ID\tTask\tCreated\tDone\n")
	fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", todo.Id, todo.Description, humanize.Time(todo.CreatedAt), todo.Done)

	w.Flush()
}
