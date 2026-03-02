package rainbow

import (
	"strings"
)

// returns a gradient colored solid block of text
func Bar(vars Variables, count, length int) string {
	var str strings.Builder

	i := 0
	for count > 0 {
		str.WriteString(Rainbow(barString(length), i, vars))
		str.WriteRune('\n')
		count--
		i++
	}

	return str.String()
}

func barString(length int) string {
	var str strings.Builder
	for range length {
		str.WriteRune(9608)
	}
	return str.String()
}
