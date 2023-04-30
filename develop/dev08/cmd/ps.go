/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	ps "github.com/mitchellh/go-ps"
	"github.com/spf13/cobra"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "process status",
	Long: `The ps utility displays a header line, followed by lines containing
	information about all of your processes that have controlling terminals.`,
	Run: func(cmd *cobra.Command, args []string) {
		w := new(tabwriter.Writer)
		w = w.Init(os.Stdout, 5, 30, 5, ' ', 0)
		list, err := ps.Processes()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Fprintln(w, "NAME\tPID\tPPID")
		for _, p := range list {
			fmt.Fprintf(w, "%s\t%d\t%d\n", p.Executable(), p.Pid(), p.PPid())
		}
		w.Flush()
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
}
