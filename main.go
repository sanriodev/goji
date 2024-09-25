package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Define available characters for each part of the emoji
var leftArms = []string{"٩", "∿", "☝", "<", "ヽ", "(", "へ", "ᕕ", "ノ"}
var rightArms = []string{"۶", "∿", "☝", ">", "ﾉ", ")", "へ", "ᕗ", "ノ"}
var eyes = []string{"o", "O", "^", "◕", "•", "°", "ʘ", "ʕ", "ಠ", "눈"}
var mouths = []string{"O", "‿", ".̫ ", "⊖", "ω", "ʖ", "﹏", "▽", "益"}

const maxVisibleOptions = 6

func main() {
	var rootCmd = &cobra.Command{
		Use:   "goji",
		Short: "Create custom or random text emojis",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("\033[H\033[2J")
			showMainMenu()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func showMainMenu() {
	// Open the keyboard for capturing inputs
	if err := keyboard.Open(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer keyboard.Close()

	// Main menu options
	menuOptions := []string{"Create New Emoji", "Create Random Emoji", "Exit"}

	selectedIndex := 0
	startIndex := 0

	for {
		printBue(`
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
		displayOptions(menuOptions, selectedIndex, startIndex)

		// Read the keyboard input
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key and char:", err, char)
			break
		}

		// Handle key input for menu
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
				createEmoji()
			case 1:
				createRandomEmoji()
			case 2:
				printRed("Exiting...")
				os.Exit(0)
			}
		case keyboard.KeyEsc:
			printRed("Exiting...")
			os.Exit(0)
		}

		// Clear terminal after each key press
		fmt.Print("\033[H\033[2J")
	}
}

func createEmoji() {
	fmt.Println("Use arrow keys to select and press Enter to confirm.")

	leftArm := pickPart(leftArms, "Pick a left arm:")
	leftEye := pickPart(eyes, "Pick a left eye:")
	mouth := pickPart(mouths, "Pick a mouth:")
	rightEye := pickPart(eyes, "Pick a right eye:")
	rightArm := pickPart(rightArms, "Pick a right arm:")

	emoji := fmt.Sprintf("%s %s %s %s %s %s %s", leftArm, "(", leftEye, mouth, rightEye, ")", rightArm)
	fmt.Println("Your custom emoji:", emoji)

	// Prompt to copy the emoji to clipboard
	copyToClipboard(emoji)
}

func createRandomEmoji() {
	// Generate random selections for emoji parts
	rand.Seed(time.Now().UnixNano())

	leftArm := leftArms[rand.Intn(len(leftArms))]
	leftEye := eyes[rand.Intn(len(eyes))]
	mouth := mouths[rand.Intn(len(mouths))]
	rightEye := eyes[rand.Intn(len(eyes))]
	rightArm := rightArms[rand.Intn(len(rightArms))]

	emoji := fmt.Sprintf("%s %s %s %s %s %s %s", leftArm, "(", leftEye, mouth, rightEye, ")", rightArm)
	// Clear terminal
	fmt.Print("\033[H\033[2J")

	fmt.Println("Random emoji generated:", emoji)

	// Prompt to copy the emoji to clipboard
	copyToClipboard(emoji)
}

func pickPart(options []string, message string) string {
	selectedIndex := 0
	startIndex := 0

	for {
		fmt.Println(message)
		displayOptions(options, selectedIndex, startIndex)
		printBue("move up or down to display more choices")
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key and char:", err, char)
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
			if selectedIndex < len(options)-1 {
				selectedIndex++
				if selectedIndex >= startIndex+maxVisibleOptions {
					startIndex++
				}
			}
		case keyboard.KeyEnter:
			fmt.Print("\033[H\033[2J")
			return options[selectedIndex]
		case keyboard.KeyEsc:
			printRed("Exiting...")
			os.Exit(0)
		}

		// Clear terminal
		fmt.Print("\033[H\033[2J")
	}
	return options[selectedIndex]
}

// displayOptions shows a window of maxVisibleOptions at a time
func displayOptions(options []string, selectedIndex, startIndex int) {
	endIndex := startIndex + maxVisibleOptions
	if endIndex > len(options) {
		endIndex = len(options)
	}

	for i := startIndex; i < endIndex; i++ {
		if i == selectedIndex {
			printSelected("> %s\n", options[i])
		} else {
			fmt.Printf("  %s\n", options[i])
		}
	}
}

func copyToClipboard(emoji string) {
	fmt.Println("Do you want to copy this emoji to the clipboard? (y/n)")

	var input string
	fmt.Scanln(&input)

	if input == "y" || input == "Y" {
		err := clipboard.WriteAll(emoji)
		if err != nil {
			fmt.Println("Error copying to clipboard:", err)
		} else {
			fmt.Println("Emoji copied to clipboard!")
		}
	} else {
		fmt.Println("Emoji not copied.")
	}
}

func printBue(text string) {
	c := color.New(color.FgBlue)
	c.Println(text)
}

func printSelected(format string, text string) {
	c := color.New(color.FgGreen)
	c.Printf(format, text)
}

func printRed(text string) {
	c := color.New(color.FgRed)
	c.Println(text)
}
