package rainbow

import (
	"strings"
	"unicode"

	"charm.land/lipgloss/v2"
)

var charColor string
var lineStyle lipgloss.Style = lipgloss.NewStyle().Foreground(lipgloss.Color(charColor))

// returns rainbow'd text
func Rainbow(input string, offset int, calculator Variables) string {
	var spaceCount int = 0
	var view strings.Builder

	var tmpStyle lipgloss.Style = lineStyle

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

		charColor = calculator.calcColor(i + offset - spaceCount)
		tmpStyle = tmpStyle.Foreground(lipgloss.Color(charColor))
		view.WriteString(tmpStyle.Render(string(char)))
	}
	return view.String()
}
