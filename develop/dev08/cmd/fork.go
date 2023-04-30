package cmd

import (
	"syscall"

	"github.com/spf13/cobra"
)

// forkCmd represents the fork command
var forkCmd = &cobra.Command{
	Use:   "fork",
	Short: "create a new process",
	Run: func(cmd *cobra.Command, args []string) {
		syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	},
}

func init() {
	rootCmd.AddCommand(forkCmd)
}
