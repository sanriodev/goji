package cmd

import (
	"fmt"
	"os"

	"github.com/sanriodev/goji/emoji"
	"github.com/spf13/cobra"
)

var newFlag bool
var randomFlag bool
var favoritesFlag bool

var rootCmd = &cobra.Command{
	Use:   "goji",
	Short: "Create custom or random text emojis",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("\033[H\033[2J")
		if newFlag {
			emoji.CreateCustomEmoji()
		} else if randomFlag {
			emoji.CreateRandomEmoji()
		} else if favoritesFlag {
			PickFavorite()
		} else {
			ShowMainMenu()
		}
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&newFlag, "new", "n", false, "Create a new emoji")
	rootCmd.Flags().BoolVarP(&randomFlag, "random", "r", false, "Create a random emoji")
	rootCmd.Flags().BoolVarP(&favoritesFlag, "favorites", "f", false, "Pick from favorites")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
