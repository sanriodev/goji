package emoji

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/sanriodev/goji/util"
)

var leftArms = []string{"٩", "∿", "☝", "<", "ヽ", "(", "へ", "ᕕ", "ノ"}
var rightArms = []string{"۶", "∿", "☝", ">", "ﾉ", ")", "へ", "ᕗ", "ノ"}
var eyes = []string{"o", "O", "^", "◕", "•", "°", "ʘ", "ʕ", "ಠ", "눈"}
var mouths = []string{"O", "‿", ".̫ ", "⊖", "ω", "ʖ", "﹏", "▽", "益"}

const maxVisibleOptions = 6

func CreateCustomEmoji() {
	fmt.Println("Use arrow keys to select and press Enter to confirm.")

	leftArm := PickPart(leftArms, "Pick a left arm:")
	leftEye := PickPart(eyes, "Pick a left eye:")
	mouth := PickPart(mouths, "Pick a mouth:")
	rightEye := PickPart(eyes, "Pick a right eye:")
	rightArm := PickPart(rightArms, "Pick a right arm:")

	emoji := fmt.Sprintf("%s %s %s %s %s %s %s", leftArm, "(", leftEye, mouth, rightEye, ")", rightArm)
	fmt.Println("Your custom emoji:", emoji)

	util.CopyToClipboard(emoji)
}

func CreateRandomEmoji() {
	rand.Seed(time.Now().UnixNano())

	leftArm := leftArms[rand.Intn(len(leftArms))]
	leftEye := eyes[rand.Intn(len(eyes))]
	mouth := mouths[rand.Intn(len(mouths))]
	rightEye := eyes[rand.Intn(len(eyes))]
	rightArm := rightArms[rand.Intn(len(rightArms))]

	emoji := fmt.Sprintf("%s %s %s %s %s %s %s", leftArm, "(", leftEye, mouth, rightEye, ")", rightArm)
	fmt.Print("\033[H\033[2J")

	fmt.Println("Random emoji generated:", emoji)
	util.CopyToClipboard(emoji)
}

func DisplayOptions(options []string, selectedIndex, startIndex int) {
	endIndex := startIndex + maxVisibleOptions
	if endIndex > len(options) {
		endIndex = len(options)
	}

	for i := startIndex; i < endIndex; i++ {
		if i == selectedIndex {
			util.PrintSelected("> %s\n", options[i])
		} else {
			fmt.Printf("  %s\n", options[i])
		}
	}
}

func PickPart(options []string, message string) string {
	selectedIndex := 0
	startIndex := 0

	for {
		fmt.Println(message)
		DisplayOptions(options, selectedIndex, startIndex)
		util.PrintBlue("move up or down to display more choices")
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
			util.PrintRed("Exiting...")
			os.Exit(0)
		}

		// Clear terminal
		fmt.Print("\033[H\033[2J")
	}
	return options[selectedIndex]
}