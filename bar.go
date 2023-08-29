package main

import (
	"strings"
)

// returns a gradient colored solid block of text
func bar(vars variables, count int) string {
	var str strings.Builder

	i := 0
	for count > 0 {
		str.WriteString(rainbow(barString(80), i, vars))
		str.WriteRune('\n')
		count--
		i++
	}

	return str.String()
}

func barString(length int) string {
	var str strings.Builder
	for i := 0; i < length; i++ {
		str.WriteRune(9608)
	}
	return str.String()
}
