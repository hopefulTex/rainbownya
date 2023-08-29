package main

import (
	"strings"
	"unicode"

	"github.com/charmbracelet/lipgloss"
)

var charColor string
var lineStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(charColor))

// returns rainbow'd text
func rainbow(input string, offset int, calculator variables) string {
	var spaceCount int = 0
	var view strings.Builder

	a := []rune(input)
	for i, char := range a {
		// don't increment gradient if there's no text to color
		if unicode.IsSpace(char) {
			view.WriteRune(char)

			spaceCount++
			if char == '\n' || char == '\r' {
				offset += 2
			}
			continue
		}

		charColor = calculator.calcColor(charColor, i+offset-spaceCount)
		lineStyle.Foreground(lipgloss.Color(charColor))
		view.WriteString(lineStyle.Render(string(char)))
	}
	return view.String()
}
