package cmd

import (
	"os"

	shell "github.com/brianstrauch/cobra-shell"
	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "msh",
	Short: "My shell",
}

func Execute() {
	rootCmd.AddCommand(shell.New(rootCmd, rootCmd.Root,
		prompt.OptionMaxSuggestion(2),
		prompt.OptionInputBGColor(prompt.Red),
		prompt.OptionDescriptionBGColor(prompt.Red)))
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func init() {
}
