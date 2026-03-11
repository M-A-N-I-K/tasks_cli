package tasks

import (
	"fmt"
	"todo/cli/pkg/tasks"

	"github.com/spf13/cobra"
)

var addTodoCmd = &cobra.Command{
    Use:   "add",
    Aliases: []string{"add"},
    Short:  "Add all tasks",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
            fmt.Println("No Arguments provided")
            return
        }
		
        success,err := tasks.AddTask(args[0])

		if err != nil{
			fmt.Println("Error adding todo %w",err)
		}
		if success {
			fmt.Println("New task added successfully!")
		}
    },
}

func init() {
    rootCmd.AddCommand(addTodoCmd)
}