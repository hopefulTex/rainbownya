package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/charmbracelet/fang"
	"github.com/charmbracelet/x/term"
	"github.com/hopefulTex/rainbownya/rainbow"
	"github.com/hopefulTex/rainbownya/ui"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
)

const DEFAULT_WIDTH = 80

func main() {

	var forceColor bool = false
	var isPiped bool = !isatty.IsTerminal(os.Stdin.Fd())
	var mathVariables rainbow.Variables = rainbow.DefaultVars()
	var barSize int = 0
	var barWidth int = -1

	var lineMode bool = false

	cmd := &cobra.Command{
		Use:   "rainbowcat",
		Short: "A nice lolcat 'clone'",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(c *cobra.Command, args []string) error {
			if forceColor || !isPiped {
				lipgloss.Writer.Profile = colorprofile.Detect(os.Stdout, []string{"COLORTERM"})
				// lipgloss.Writer = colorprofile.NewWriter(os.Stdout, os.Environ())
			}

			if barSize > 0 {
				var width int = 0

				if barWidth != -1 {
					width = barWidth
				} else {
					w, _, err := term.GetSize(os.Stdout.Fd())
					if err != nil {
						w = DEFAULT_WIDTH
					}
					width = min(w, DEFAULT_WIDTH)
				}

				t := rainbow.Bar(mathVariables, barSize, width)
				fmt.Println(t)
				return nil
			}

			if isPiped {
				rainbowPipe(mathVariables)
				return nil
			}

			if lineMode {
				text := ui.GetLine()
				fmt.Print(text)
				return nil
			}

			if len(args) < 1 {
				ui.Interactive(mathVariables)
				return nil
			}

			var buf strings.Builder
			for _, arg := range args {
				t, err := os.ReadFile(arg)
				if err != nil {
					log.Printf("error opening file: `%s`", arg)
					continue
				}
				buf.Write(t)
			}
			fmt.Print(rainbow.Rainbow(buf.String(), 0, mathVariables))

			return nil
		},
	}
	cmd.Version = VERSION

	cmd.Flags().IntVar(&barSize, "bar", 0, "Print a pretty bar of size n")
	cmd.Flags().IntVar(&barWidth, "bar-width", -1, "Change the bar width")

	cmd.Flags().Float64VarP(&mathVariables.Spread, "spread", "s", 3.0, "Set spread")
	cmd.Flags().Float64VarP(&mathVariables.Freq, "frequency", "f", 0.3, "Set frequency")
	cmd.Flags().Float64VarP(&mathVariables.Width, "width", "w", 127.5, "Set width of color band")
	cmd.Flags().Float64VarP(&mathVariables.Center, "center", "c", 127.5, "Set center")
	cmd.Flags().Float64VarP(&mathVariables.Seed, "seed", "S", rand.Float64()*128, "Set seed")

	cmd.Flags().BoolVarP(&forceColor, "TTY", "t", false, "Force TTY(color) mode")
	cmd.Flags().BoolVar(&lineMode, "get-line", false, "Get a single line and direct it to stdout")

	if err := fang.Execute(
		context.Background(),
		cmd,
		fang.WithNotifySignal(os.Interrupt, os.Kill),
		fang.WithVersion(VERSION),
	); err != nil {
		os.Exit(1)
	}
}
