package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Pick a left arm:")
	leftArm := pickPart(reader, leftArms)

	fmt.Println("Pick a left eye:")
	leftEye := pickPart(reader, eyes)

	fmt.Println("Pick a mouth:")
	mouth := pickPart(reader, mouths)

	fmt.Println("Pick a right eye:")
	rightEye := pickPart(reader, eyes)

	fmt.Println("Pick a right arm:")
	rightArm := pickPart(reader, rightArms)

	// Generate and display the final emoji
	emoji := fmt.Sprintf("%s%s%s%s%s%s%s", leftArm, "(", leftEye, mouth, rightEye, ")", rightArm)
	fmt.Println("Your custom emoji:", emoji)
}

func pickPart(reader *bufio.Reader, options []string) string {
	for i, option := range options {
		fmt.Printf("%d: %s\n", i+1, option)
	}

	fmt.Print("Enter the number for your choice: ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr) // Remove any trailing newlines or spaces

	choiceIndex, err := strconv.Atoi(choiceStr) // Convert the string to an integer
	if err != nil || choiceIndex < 1 || choiceIndex > len(options) {
		fmt.Println("Invalid choice, please try again.")
		return pickPart(reader, options)
	}

	return options[choiceIndex-1]
}
