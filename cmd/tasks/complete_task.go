package tasks

import (
	"fmt"
	"strconv"
	"todo/cli/pkg/tasks"

	"github.com/spf13/cobra"
)

var completeTaskCmd = &cobra.Command{
    Use:   "complete",
    Aliases: []string{"complete"},
    Short:  "Complete a todo by id",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
            fmt.Println("No Argument provided")
            return
        }
		id, err := strconv.Atoi(args[0])
		if err != nil{
			fmt.Println("Error converting id to int : %w",err)
		}
		
        success,err := tasks.CompleteTask(id)
		if err != nil{
			fmt.Println("Error deleting todo %w",err)
		}
		if success {
			fmt.Printf("Task with id %d completed successfully!",id)
		}
    },
}

func init() {
    rootCmd.AddCommand(completeTaskCmd)
}