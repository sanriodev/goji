package util

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func CopyToClipboard(emoji string, cannotReSave bool) {
	fmt.Println("Do you want to copy this emoji to the clipboard? (y/n/save)")

	var input string
	fmt.Scanln(&input)

	if input == "y" || input == "Y" {
		err := clipboard.WriteAll(emoji)
		if err != nil {
			fmt.Println("Error copying to clipboard:", err)
		} else {
			fmt.Println("Emoji copied to clipboard!")
		}
	} else if input == "save" && !cannotReSave {
		addFavorite(emoji)
		fmt.Println("Emoji saved to favorites!")
	} else {
		fmt.Println("Emoji not copied.")
	}
}
