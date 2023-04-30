package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// echoCmd represents the echo command
var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "write arguments to the standard output",
	Long: `The echo utility writes any specified operands, separated by single blank
	(‘ ’) characters and followed by a newline (‘\n’) character, to the
	standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(strings.Join(args[0:], " "))
	},
}

func init() {
	rootCmd.AddCommand(echoCmd)
}
