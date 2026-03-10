package todos

import (
	"fmt"
	"strconv"
	"todo/cli/pkg/todos"

	"github.com/spf13/cobra"
)

var addTodoCmd = &cobra.Command{
    Use:   "add",
    Aliases: []string{"add"},
    Short:  "Add all todos",
    Args:  cobra.ExactArgs(3),
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
            fmt.Println("Not Enough Arguments provided")
            return
        }
		id, err := strconv.Atoi(args[0])
		if err != nil{
			fmt.Println("Error converting id to int : %w",err)
		}
		todoItem := todos.Todo{id,args[1],args[2],false}
		
        res,err := todos.AddTodo(todoItem)
		if err != nil{
			fmt.Println("Error adding todo %w",err)
		}
        fmt.Println(res)
    },
}

func init() {
    rootCmd.AddCommand(addTodoCmd)
}