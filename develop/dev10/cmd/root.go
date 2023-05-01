/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/nilovartem/l2/develop/dev10/telnet"
	"github.com/spf13/cobra"
)

var (
	timeout uint
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-telnet [--timeout] host port",
	Short: "go-telnet",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(timeout)
		if len(args) == 2 {
			host := args[0]
			port := args[1]
			fmt.Println(host, port)
			telnet.Connect(timeout, host, port)
		} else {
			fmt.Println("Must be host and port")
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().UintVar(&timeout, "timeout", 10, "--timeout=10")
}
