package tasks

import (
	"fmt"
	"todo/cli/pkg/tasks"

	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:     "stats",
	Aliases: []string{"stats"},
	Short:   "List Statistics",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var res []tasks.Task
		var completedTasks int
		var pendingTasks int

		w := tabwriter.NewWriter(os.Stdout, 10, 0, 1, ' ', 0)

		res, err := tasks.ListTasks("tasks.json")

		if err != nil {
			fmt.Println("Error listing stats", err)
		}

		completedTasks = 0
		pendingTasks = 0

		for _, todo := range res {
			if todo.IsComplete {
				completedTasks = completedTasks + 1
			} else {
				pendingTasks = pendingTasks + 1
			}

		}
		fmt.Fprintf(w, "Completed Tasks : %d\n", completedTasks)
		fmt.Fprintf(w, "Pending Tasks : %d\n", pendingTasks)

		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
