package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "specifies a command line for another command",
	Run: func(cmd *cobra.Command, args []string) {
		ex := exec.Command(args[0], args[1:]...)
		ex.Stderr = os.Stderr
		ex.Stdout = os.Stdout
		err := ex.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
