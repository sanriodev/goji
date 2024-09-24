package main

import (
	"fmt"
	"os"

	"github.com/eiannone/keyboard"
	"github.com/spf13/cobra"
)

// Define available characters for each part of the emoji
var leftArms = []string{"٩", "∿", "☝", "<"}
var rightArms = []string{"۶", "∿", "☝", ">"}
var eyes = []string{"o", "O", "^", "◕", "•", "°", "ʘ"}
var mouths = []string{"O", "‿", ".̫ ", "⊖", "ω", "ʖ"}

func main() {
	var rootCmd = &cobra.Command{Use: "gomoji"}

	var newCmd = &cobra.Command{
		Use:   "new",
		Short: "Create a new text emoji",
		Run: func(cmd *cobra.Command, args []string) {
			createEmoji()
		},
	}

	rootCmd.AddCommand(newCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func createEmoji() {
	fmt.Println("Use arrow keys to select and press Enter to confirm.")

	// Open the keyboard for capturing inputs
	if err := keyboard.Open(); err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer keyboard.Close()

	leftArm := pickPart(leftArms, "Pick a left arm:")

	leftEye := pickPart(eyes, "Pick a left eye:")

	mouth := pickPart(mouths, "Pick a mouth:")

	rightEye := pickPart(eyes, "Pick a right eye:")

	rightArm := pickPart(rightArms, "Pick a right arm:")

	emoji := fmt.Sprintf("%s %s %s %s %s %s %s", leftArm, "(", leftEye, mouth, rightEye, ")", rightArm)
	fmt.Println("Your custom emoji:", emoji)
}

func pickPart(options []string, message string) string {
	selectedIndex := 0

	for {
		fmt.Println(message)
		// Display options with highlighting for the current selection
		displayOptions(options, selectedIndex)

		// Read the keyboard input
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error reading key and char:", err, char)
			break
		}

		// Handle key input
		switch key {
		case keyboard.KeyArrowUp:
			// Move up in the selection list
			if selectedIndex > 0 {
				selectedIndex--
			}
		case keyboard.KeyArrowDown:
			// Move down in the selection list
			if selectedIndex < len(options)-1 {
				selectedIndex++
			}
		case keyboard.KeyEnter:
			// Confirm selection with Enter
			fmt.Print("\033[H\033[2J")
			return options[selectedIndex]
		case keyboard.KeyEsc:
			// Exit on ESC (optional)
			fmt.Println("Exiting...")
			os.Exit(0)
		}
					
					fmt.Print("\033[H\033[2J")
	}
	return options[selectedIndex]
}

func displayOptions(options []string, selectedIndex int) {
	for i, option := range options {
		if i == selectedIndex {
			// Highlight the selected option
			fmt.Printf("> %s\n", option)
		} else {
			fmt.Printf("  %s\n", option)
		}
	}
}