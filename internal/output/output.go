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

	if all == true {
		fmt.Fprintf(w, "ID\tTask\tCreated\tDone\n")
	} else {
		fmt.Fprintf(w, "ID\tTask\tCreated\n")
	}

	for _, v := range todos {
		if all == true {
			fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", v.Id, v.Description, humanize.Time(v.CreatedAt), v.Done)
		} else {
			fmt.Fprintf(w, "%d\t%s\t%s\n", v.Id, v.Description, humanize.Time(v.CreatedAt))
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

// func PrintHelp() {
// 	fmt.Print("A simple todo list for the terminal written in go\n\nUsage:\n\ttasks [-command=value]\n\nAvailable Commands:\nadd\t\t\tAdd a new task to the todo list\ncomplete\t\tSet a task as being completed\ndelete\t\t\tRemoves a task for the todo list by it's id\nlist\t\t\tLists all of the tasks in your todo list\nhelp\t\t\tHelp about any command\n\nFlags:\n-a, --all\t\tList all todos, including the already done\n-h, --help\t\tHelp for tasks\n\nUse tasks [command] --help for more information about a command.\n")
// }
