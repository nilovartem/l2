package cmd

import (
	"os"
	"sync"

	"github.com/nilovartem/l2/develop/dev09/wget"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "wget URL...",
	Short: "wget - скачивает весь сайт",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		wg := sync.WaitGroup{}
		for _, addr := range args {
			wg.Add(1)
			go wget.Wget(addr, &wg)
		}
		wg.Wait()
	},
}

func Execute() {
	err := Cmd.Execute()
	if err != nil {
		Cmd.Help()
		os.Exit(1)
	}
}

func init() {
	Cmd.Flags().BoolP("help", "h", false, "помощь по gowget")

}
