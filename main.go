package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
	"github.com/muesli/termenv"
)

func main() {
	rand.Seed(time.Now().UnixMilli())
	var isPiped bool = false
	vars, beginText := parseFlags(os.Args, newVars())

	if !isatty.IsTerminal(os.Stdin.Fd()) {
		isPiped = true
	}

	if vars.forceColor || isatty.IsTerminal(os.Stdout.Fd()) {
		lipgloss.SetColorProfile(termenv.TrueColor)
	}

	if !isPiped && len(os.Args) < 2 {
		interactive(vars)
		quit()
	}
	// print rainbow bar
	if !isPiped && os.Args[1] == "--bar" {
		var barCount int64 = 1
		if len(os.Args) > 2 {
			barCount, _ = strconv.ParseInt(os.Args[2], 10, 64)
		}
		if barCount < 1 {
			barCount = 1
		}
		fmt.Println(bar(vars, int(barCount)))
		quit()
	}

	if isPiped {
		rainbowPipe(vars)
	} else {
		str := strings.Join(os.Args[beginText:], " ")
		fmt.Print(rainbow(str, 0, vars))
	}

	quit()
}

// pro tip: its rude to break peoples tty's
func quit() {
	lipgloss.DefaultRenderer().Output().Reset()
	fmt.Println("")
	os.Exit(0)
}
