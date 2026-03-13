package tasks

import (
	"fmt"
	"todo/cli/pkg/tasks"

	"os"
	"text/tabwriter"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var completedTasksCmd = &cobra.Command{
    Use:   "completed",
    Aliases: []string{"completed"},
    Short:  "Returns all completed tasks",
    Args:  cobra.ExactArgs(0),
    Run: func(cmd *cobra.Command, args []string) {
		w := tabwriter.NewWriter(os.Stdout, 10, 0, 1, ' ', 0)
        res,err := tasks.ListCompleteTasks("tasks.json")
		if err != nil{
			fmt.Println("Error listing tasks %w",err)
		}
		fmt.Fprintln(w, "ID\tTask\t\tCreatedAt")
		for _,todo := range res {
			createdAt := timediff.TimeDiff(todo.CreatedAt)
			fmt.Fprintf(w,"%d\t%s\t\t%s\n",todo.ID,todo.Description,createdAt)
		}
		w.Flush()
    },
}

func init() {
    rootCmd.AddCommand(completedTasksCmd)
}