package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sort ",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Short: "sort records (lines) of text file",
	Long:  `The sort utility sorts text file by lines.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetInt("key")
		numeric, _ := cmd.Flags().GetBool("numeric-sort")
		reverse, _ := cmd.Flags().GetBool("reverse")
		unique, _ := cmd.Flags().GetBool("unique")

		strs, err := readFile(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		strs = Sort(strs, key, numeric, reverse, unique)

		fmt.Fprintln(os.Stdin, strings.Join(strs, "\n"))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
func readFile(filename string) ([]string, error) {
	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	result := strings.Split(string(fileBytes), "\n")
	return result, nil
}
func init() {
	rootCmd.Flags().IntP("key", "k", 0, "указание колонки для сортировки")
	rootCmd.Flags().BoolP("numeric-sort", "n", false, "сортировать по числовому значению")
	rootCmd.Flags().BoolP("reverse", "r", false, "сортировать в обратном порядке")
	rootCmd.Flags().BoolP("unique", "u", false, "не выводить повторяющиеся строки")
}
