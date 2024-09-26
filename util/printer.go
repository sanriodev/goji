package util

import (
	"github.com/fatih/color"
)

func PrintBlue(text string) {
	c := color.New(color.FgBlue)
	c.Println(text)
}

func PrintRed(text string) {
	c := color.New(color.FgRed)
	c.Println(text)
}

func PrintSelected(format string, text string) {
	c := color.New(color.FgGreen)
	c.Printf(format, text)
}
