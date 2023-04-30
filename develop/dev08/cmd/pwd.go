/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// pwdCmd represents the pwd command
var pwdCmd = &cobra.Command{
	Use:   "pwd",
	Short: "return working directory name",
	Long: `The pwd utility writes the absolute pathname of the current working
	directory to the standard output.`,
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println(pwd)
	},
}

func init() {
	rootCmd.AddCommand(pwdCmd)
}
