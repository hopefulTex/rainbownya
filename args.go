package main

// TODO: Learn Cobra oml this is awful I left c++ to avoid doing these!!!
import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

// When I have the resolve to learn Cobra *shudder* this hack will be defunct
func parseFlags(args []string, vars variables) (variables, int) {
	var isArgs bool = true
	var beginText int = 0
	for i, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			if isArgs {
				isArgs = false
				continue
			}
			beginText = i
			break
		}
		isArgs = true
		if arg == "--help" {
			helpScreen()
			os.Exit(0)
		}
		if arg == "--bar" {
			return vars, i
		}
		if i >= len(args)-1 {
			log.Fatal(fmt.Errorf("missing argument after flag: %s", arg))
		}

		float, err := strconv.ParseFloat(args[i+1], 64)
		if err != nil {
			helpScreen()
			log.Fatal(err)
		}
		switch arg {
		case "-v", "--version":
			versionScreen()
			os.Exit(0)
		case "-s":
			vars.setSpread(float)
		case "-f":
			vars.setFrequency(float)
		case "-S":
			vars.seed = (float)
		case "-w":
			vars.setWidth(float)
		case "-c":
			vars.setCenter(float)
		case "-t":
			vars.forceTTY()
		}
	}
	if vars.seed == -1.0 {
		vars.setSeed(rand.Float64() * 128)
	}

	return vars, beginText
}

func helpScreen() {
	versionScreen()
	fmt.Printf("-s\tSpread\t\tSet spread\n")
	fmt.Printf("-f\tFrequency\tSet frequency\n")
	fmt.Printf("-S\tSeed\t\tSet seed\n")
	fmt.Printf("-w\tWidth\t\tSet width\n")
	fmt.Printf("-c\tCenter\t\tSet center\n")
	fmt.Printf("-v\tVersion\t\tShow version\n")
	fmt.Printf("-t\tTTY\t\tForce TTY(color) mode\n")
}

func versionScreen() {
	fmt.Printf("lulcat\nVersion 0.0.1\nA lolcat \"clone\" made in go\n")
}
