package todos

import (
	"fmt"
	"todo/cli/pkg/todos"

	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var listTodoCmd = &cobra.Command{
    Use:   "list",
    Aliases: []string{"list"},
    Short:  "List all todos",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
        res,err := todos.ListTodo("todos.json")
		if err != nil{
			fmt.Println("Error listing todo %w",err)
		}
		for _,todo := range res {
			fmt.Fprintln(w,todo.ID,todo.Description,todo.CreatedAt)
		}

    },
}

func init() {
    rootCmd.AddCommand(listTodoCmd)
}