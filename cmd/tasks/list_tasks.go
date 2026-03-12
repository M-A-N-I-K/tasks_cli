package tasks

import (
	"fmt"
	"todo/cli/pkg/tasks"

	"os"
	"text/tabwriter"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var listTasksCmd = &cobra.Command{
    Use:   "list",
    Aliases: []string{"list"},
    Short:  "List pending tasks",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, '.', tabwriter.AlignRight|tabwriter.Debug)
        res,err := tasks.ListTasks("tasks.json")
		if err != nil{
			fmt.Println("Error listing tasks %w",err)
		}
		for _,todo := range res {
			createdAt := timediff.TimeDiff(todo.CreatedAt)
			fmt.Fprintln(w,todo.ID,todo.Description,createdAt)
		}

    },
}

func init() {
    rootCmd.AddCommand(listTasksCmd)
}