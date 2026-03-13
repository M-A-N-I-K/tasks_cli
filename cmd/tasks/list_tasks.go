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
		var res  []tasks.Task
		var err error

		w := tabwriter.NewWriter(os.Stdout, 10, 0, 1, ' ', 0)
		all,_ := cmd.Flags().GetBool("all")

		if all {
			res,err = tasks.ListTasks("tasks.json")
		}else{
			res,err = tasks.ListPendingTasks("tasks.json")
		}

		if err != nil{
			fmt.Println("Error listing tasks %w",err)
		}
		if all {
			fmt.Fprintln(w, "ID\tTask\t\tCreatedAt\tStatus")
		}else {
			fmt.Fprintln(w, "ID\tTask\t\tCreatedAt")
		}
		for _,todo := range res {
			createdAt := timediff.TimeDiff(todo.CreatedAt)
			if all {
				fmt.Fprintf(w,"%d\t%s\t\t%s\t%t\n",todo.ID,todo.Description,createdAt,todo.IsComplete)
				}else{
				fmt.Fprintf(w,"%d\t%s\t\t%s\n",todo.ID,todo.Description,createdAt)
			}
		}
		w.Flush()
    },
}

func init() {
	listTasksCmd.Flags().Bool("all", false, "List all tasks")
    rootCmd.AddCommand(listTasksCmd)
}