package cmd

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/sanriodev/goji/emoji"
	"github.com/sanriodev/goji/util"
)

const maxVisibleOptions = 6

func ShowMainMenu() {
	if err := keyboard.Open(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer keyboard.Close()
	menuOptions := []string{"Create New Emoji", "Create Random Emoji", "Copy from Favorites", "Exit"}
	selectedIndex := 0
	startIndex := 0

	for {
		util.PrintBlue(`
      ___         ___                       
     /\__\       /\  \      ___             
    /:/ _/_     /::\  \    /\__\    ___      
   /:/ /\  \   /:/\:\  \  /:/__/   /\__\     
  /:/ /::\  \ /:/  \:\  \/::\  \  /:/__/     
 /:/__\/\:\__/:/__/ \:\__\/\:\  \/::\  \     
 \:\  \ /:/  \:\  \ /:/  /~~\:\  \/\:\  \__ 
  \:\  /:/  / \:\  /:/  /    \:\__~~\:\/\__\
   \:\/:/  /   \:\/:/  /     /:/  /  \::/  /
    \::/  /     \::/  /     /:/  /   /:/  /
     \/__/       \/__/      \/__/    \/__/  
`)
		fmt.Println("Select an option:")
		emoji.DisplayOptions(menuOptions, selectedIndex, startIndex)

		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err, char)
			break
		}

		switch key {
		case keyboard.KeyArrowUp:
			if selectedIndex > 0 {
				selectedIndex--
				if selectedIndex < startIndex {
					startIndex--
				}
			}
		case keyboard.KeyArrowDown:
			if selectedIndex < len(menuOptions)-1 {
				selectedIndex++
				if selectedIndex >= startIndex+maxVisibleOptions {
					startIndex++
				}
			}
		case keyboard.KeyEnter:
			fmt.Print("\033[H\033[2J")
			switch selectedIndex {
			case 0:
				emoji.CreateCustomEmoji()
			case 1:
				emoji.CreateRandomEmoji()
			case 2:
				PickFavorite()
				os.Exit(0)
			case 3:
				util.PrintRed("Exiting...")
				os.Exit(0)
			}
		case keyboard.KeyEsc:
			util.PrintRed("Exiting...")
			os.Exit(0)
		}
		fmt.Print("\033[H\033[2J")
	}
}

/* func listFavorites() {
	favorites := util.LoadFavorites()

	if len(favorites.Emojis) == 0 {
		fmt.Println("No favorite emojis yet!")
		return
	}

	fmt.Println("Your favorite emojis:")
	for i, emoji := range favorites.Emojis {
		emojiStr := fmt.Sprintf("%d: %s", i+1, emoji.Content)
		util.PrintBlue(emojiStr)
	}
} */

func PickFavorite() {
	favorites := util.LoadFavorites()

	if len(favorites.Emojis) == 0 {
		fmt.Println("No favorite emojis saved yet!")
		return
	}

	selectedIndex := 0
	startIndex := 0

	for {
		fmt.Println("Pick a favorite emoji to copy:")
		emoji.DisplayOptions(util.GetFavoriteContents(favorites), selectedIndex, startIndex)

		// Capture user input
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key:", err, char)
			break
		}

		switch key {
		case keyboard.KeyArrowUp:
			if selectedIndex > 0 {
				selectedIndex--
				if selectedIndex < startIndex {
					startIndex--
				}
			}
		case keyboard.KeyArrowDown:
			if selectedIndex < len(favorites.Emojis)-1 {
				selectedIndex++
				if selectedIndex >= startIndex+maxVisibleOptions {
					startIndex++
				}
			}
		case keyboard.KeyEnter:
			emoji := favorites.Emojis[selectedIndex].Content
			util.CopyToClipboard(emoji, true)
			return
		case keyboard.KeyEsc:
			fmt.Println("Exiting favorites menu.")
			return
		}

		// Clear terminal after each key press
		fmt.Print("\033[H\033[2J")
	}
}
