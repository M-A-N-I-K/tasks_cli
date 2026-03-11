package tasks

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:  "tasks",
    Short: "tasks - a simple CLI to add, update or delete tasks",
    Long: `Tasks is a super fancy CLI (kidding)
One can use tasks to modify or inspect daily tasks straight from the terminal`,
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}