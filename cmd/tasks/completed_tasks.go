package tasks

import (
	"fmt"
	"todo/cli/pkg/tasks"

	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var completedTasksCmd = &cobra.Command{
    Use:   "complete",
    Aliases: []string{"complete"},
    Short:  "Returns all completed tasks",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
        res,err := tasks.ListTasks("tasks.json")
		if err != nil{
			fmt.Println("Error listing tasks %w",err)
		}
		for _,todo := range res {
			fmt.Fprintln(w,todo.ID,todo.Description,todo.CreatedAt)
		}

    },
}

func init() {
    rootCmd.AddCommand(completedTasksCmd)
}