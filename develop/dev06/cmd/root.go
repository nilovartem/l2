package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cut [FLAGS] TEXT",
	Short: "принимает строки через STDIN, разбивает по разделителю (TAB) на колонки и выводит запрошенные",
	Run: func(cmd *cobra.Command, args []string) {
		delimiter, _ := cmd.Flags().GetString("delimiter")
		fields, _ := cmd.Flags().GetUint("fields")
		separated, _ := cmd.Flags().GetBool("separated")
		reader := bufio.NewReader(os.Stdin)
		for {
			str, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			result := Cut(str[:len(str)-1], delimiter, fields, separated)

			fmt.Println(result)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("help", "h", false, "помощь по cut")
	rootCmd.Flags().UintP("fields", "f", 0, "выбрать поля (колонки). 0 по умолчанию")
	rootCmd.Flags().StringP("delimiter", "d", "\t", "использовать другой разделитель")
	rootCmd.Flags().BoolP("separated", "s", false, "только строки c разделителем")
}
