package tasks

import (
	"fmt"
	"strconv"
	"todo/cli/pkg/tasks"

	"github.com/spf13/cobra"
)

var updateTodoCmd = &cobra.Command{
    Use:   "update",
    Aliases: []string{"update"},
    Short:  "Update a todo by id",
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
		
        res,err := tasks.UpdateTask(id,args[1],args[2] == "true")
		if err != nil{
			fmt.Println("Error deleting todo %w",err)
		}
        if res{
			fmt.Println("Successfully updated todo with id :",id)
		}
    },
}

func init() {
    rootCmd.AddCommand(updateTodoCmd)
}