package main

//TODO: Fix the second line always having wrong offset

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func rainbowPipe(vars variables) {
	var str string
	var eof bool = false
	var i int = 0
	var leftover string = ""
	var lastLine int
	for !eof {
		str, eof = readChunk(os.Stdin)

		str = leftover + str
		lastLine = strings.LastIndex(str, "\n")
		if lastLine == -1 {
			leftover = ""
		} else {
			leftover = str[lastLine:]
		}
		fmt.Print(rainbow(str[:lastLine], i, vars))
		i++
	}
}

// basically a copy of io.ReadAll() without the loop
// returns most recent text and if its at EOF
func readChunk(r io.Reader) (string, bool) {
	var eof bool = false

	var bytes []byte = make([]byte, 1024)
	if len(bytes) == cap(bytes) {
		bytes = append(bytes, 0)[:len(bytes)]
	}

	n, err := r.Read(bytes[len(bytes):cap(bytes)])
	bytes = bytes[:len(bytes)+n]

	if err != nil {
		eof = true
	}
	return string(bytes), eof
}
