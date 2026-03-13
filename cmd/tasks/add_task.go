package tasks

import (
	"fmt"
	"time"
	"todo/cli/pkg/tasks"

	"github.com/spf13/cobra"
)

var addTodoCmd = &cobra.Command{
    Use:   "add",
    Aliases: []string{"add"},
    Short:  "Add all tasks",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        var due_date time.Time
        var err error
		if len(args) == 0 {
            fmt.Println("No Arguments provided")
            return
        }

        due,_ := cmd.Flags().GetString("due")

        layout := "2006-01-02"
        if(due != ""){
            due_date, err = time.Parse(layout, due)
            if err != nil{
                fmt.Printf("Error adding todo :",err)
                return
            }
        }

        success,err := tasks.AddTask(args[0],&due_date)

		if err != nil{
			fmt.Println("Error adding todo %w",err)
            return
		}
		if success {
			fmt.Println("New task added successfully!")
		}
    },
}

func init() {
    addTodoCmd.Flags().String("due", "", "Enter due date for task")
    rootCmd.AddCommand(addTodoCmd)
}