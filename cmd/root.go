package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goji",
	Short: "Create custom or random text emojis",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("\033[H\033[2J")
		ShowMainMenu()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
