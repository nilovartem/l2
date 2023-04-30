package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// killCmd represents the kill command
var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "terminate or signal a process",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(os.Stderr, "kill: not enough arguments")
		}
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		proc, err := os.FindProcess(pid)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = proc.Kill()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(killCmd)
}
