package cmd

import (
	"bufio"
	"fmt"
	"io"

	"os"

	"github.com/nilovartem/l2/develop/dev05/gogrep"
	"github.com/spf13/cobra"
)

var flags = gogrep.Flags{}

var Cmd = &cobra.Command{
	Use:   "gogrep [OPTIONS] PATTERN [FILE...]",
	Short: "gogrep - печатает строки, подходящие по шаблону",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		var strs []string
		if len(args) > 1 {
			files := args[1:]
			var err error
			strs, err = readFiles(files)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		} else {
			reader := bufio.NewReader(os.Stdin)
			for {
				line, err := reader.ReadString('\n')
				if err == io.EOF {
					break
				}
				strs = append(strs, line)
			}
		}
		result := gogrep.Grep(strs, args[0], flags)
		fmt.Print(result)
	},
}

func readFiles(files []string) ([]string, error) {
	var result []string

	for _, filepath := range files {
		file, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}
		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			result = append(result, line)
		}
	}
	return result, nil
}

func Execute() {
	err := Cmd.Execute()
	if err != nil {
		Cmd.Help()
		os.Exit(1)
	}
}

func init() {
	Cmd.Flags().BoolP("help", "h", false, "помощь по strsort")
	Cmd.Flags().UintVarP(&flags.After, "after", "A", 0, "печатать +N строк после совпадения")
	Cmd.Flags().UintVarP(&flags.Before, "before", "B", 0, "печатать +N строк до совпадения")
	Cmd.Flags().UintVarP(&flags.Context, "context", "C", 0, "(A+B) печатать ±N строк вокруг совпадения")
	Cmd.Flags().BoolVarP(&flags.Count, "count", "c", false, "количество строк")
	Cmd.Flags().BoolVarP(&flags.IgnoreCase, "ignore-case", "i", false, "игнорировать регистр")
	Cmd.Flags().BoolVarP(&flags.Invert, "invert", "v", false, "вместо совпадения, исключать")
	Cmd.Flags().BoolVarP(&flags.Fixed, "fixed", "F", false, "точное совпадение со строкой, не паттерн")
	Cmd.Flags().BoolVarP(&flags.LineNum, "line-num", "n", false, " печатать номер строки")
}
