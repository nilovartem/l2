package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// cdCmd represents the cd command
var cdCmd = &cobra.Command{
	Use:   "cd",
	Short: "Displays the name of the current directory or changes the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			pwdCmd.Run(rootCmd, []string{})
		default:
			err := os.Chdir(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cdCmd)
}
