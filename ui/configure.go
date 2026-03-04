package ui

import (
	"fmt"
	"image/color"
	"log"
	"strconv"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/hopefulTex/rainbownya/rainbow"
)

type coordinates struct {
	x int
	y int
}

const (
	SEEDINDEX int = iota
	FREQINDEX
	SPREADINDEX
	WIDTHINDEX
	CENTERINDEX
)

type textBox struct {
	input textinput.Model
	value float64
}

type Configure struct {
	baseStyle  lipgloss.Style
	inputStyle lipgloss.Style

	preview   string
	barWidth  int
	barHeight int

	config      rainbow.Variables
	inputs      []textBox
	activeInput int

	pointer coordinates

	wantsOutput bool
}

var ACTIVECOLOR color.Color = lipgloss.Color("#ff99c8")

func NewConfigureScreen() Configure {
	v := rainbow.DefaultVars()
	t := textBox{input: textinput.New(), value: 0.0}

	style := textinput.DefaultDarkStyles()
	style.Focused.Prompt = lipgloss.NewStyle().Foreground(ACTIVECOLOR)
	t.input.SetStyles(style)
	t.input.SetWidth(8)

	t.input.Blur()

	var boxes []textBox = make([]textBox, 5)

	for i := range boxes {
		switch i {
		case SEEDINDEX:
			t.input.Placeholder = fmt.Sprint(v.Seed)
			t.value = v.Seed
		case SPREADINDEX:
			t.input.Placeholder = fmt.Sprint(v.Spread)
			t.value = v.Spread
		case WIDTHINDEX:
			t.input.Placeholder = fmt.Sprint(v.Width)
			t.value = v.Width
		case FREQINDEX:
			t.input.Placeholder = fmt.Sprint(v.Freq)
			t.value = v.Freq
		case CENTERINDEX:
			t.input.Placeholder = fmt.Sprint(v.Center)
			t.value = v.Center
		}
		boxes[i] = t
	}
	boxes[0].input.Focus()

	return Configure{
		baseStyle:  lipgloss.NewStyle().Width(80).Border(lipgloss.RoundedBorder()),
		inputStyle: lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Width(16).Height(1),

		preview:   rainbow.Bar(v, 6, 16),
		barWidth:  16,
		barHeight: 6,

		config: v,

		inputs:      boxes,
		activeInput: 0,

		pointer: coordinates{x: 0, y: 0},

		wantsOutput: false,
	}
}

func (c Configure) Init() tea.Cmd {
	return textinput.Blink
}

func (c Configure) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.baseStyle = c.baseStyle.Width(min(msg.Width, 30))
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+q", "esc", "enter":
			if msg.String() == "enter" {
				c.wantsOutput = true
			}
			return c, tea.Quit
		case "up", "ctrl+tab":
			if c.activeInput > 0 {
				c.inputs[c.activeInput].input.Blur()
				c.activeInput--
				c.inputs[c.activeInput].input.Focus()
			}
		case "down", "tab":
			if c.activeInput < len(c.inputs)-1 {
				c.inputs[c.activeInput].input.Blur()
				c.activeInput++
				c.inputs[c.activeInput].input.Focus()
			}
		default:
			input := c.inputs[c.activeInput].input
			prev := input.Value()
			if prev == "" {
				prev = input.Placeholder
			}

			// fix '.5' -> '0.5'
			if msg.Code == '.' && input.Value() == "" {
				input.SetValue("0")
			}

			c.inputs[c.activeInput].input, cmd = input.Update(msg)
			input = c.inputs[c.activeInput].input

			current := input.Value()
			if current == "" {
				current = input.Placeholder
			}

			if prev != current {
				fl, err := strconv.ParseFloat(current, 64)
				if err != nil {
					fmt.Println("ERROR")
				}
				if err == nil {
					c.inputs[c.activeInput].value = fl
					c = c.updatePreview()
				}
			}
		}
	}
	return c, cmd
}

func (c Configure) View() tea.View {
	// get background color @ x,y
	return tea.NewView(lipgloss.JoinHorizontal(lipgloss.Center, c.preview, " ", c.controlsView()))
}

func (c Configure) updatePreview() Configure {
	c.config.Seed = c.inputs[SEEDINDEX].value
	c.config.Spread = c.inputs[SPREADINDEX].value
	c.config.Center = c.inputs[CENTERINDEX].value
	c.config.Width = c.inputs[WIDTHINDEX].value
	c.config.Freq = c.inputs[FREQINDEX].value

	c.preview = rainbow.Bar(c.config, c.barHeight, c.barWidth)
	return c
}

func (c Configure) controlsView() string {
	var view strings.Builder
	var boxLabel string

	for i, textbox := range c.inputs {
		switch i {
		case SEEDINDEX:
			boxLabel = "Seed:     "
		case SPREADINDEX:
			boxLabel = "Spread:   "
		case WIDTHINDEX:
			boxLabel = "Width:    "
		case FREQINDEX:
			boxLabel = "Frequency:"
		case CENTERINDEX:
			boxLabel = "Center:   "
		}

		if i == c.activeInput {
			c.inputStyle = c.inputStyle.BorderForeground(ACTIVECOLOR)
		}

		view.WriteString(lipgloss.JoinHorizontal(lipgloss.Center, boxLabel, c.inputStyle.Render(textbox.input.View())))
		view.WriteRune('\n')
		if i == c.activeInput {
			c.inputStyle = c.inputStyle.UnsetBorderForeground()
		}
	}

	return c.baseStyle.Render(strings.TrimSuffix(view.String(), "\n"))
}

var colorPreviewStyle lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Width(9).Height(1)
var colorTileStyle lipgloss.Style = lipgloss.NewStyle().Width(1).Height(1).Background(lipgloss.Color("323232"))

func colorPreview(hex string) string {
	hex = strings.TrimSpace(hex)
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return colorPreviewStyle.Render("🚫No Color")
	}
	colorTileStyle = colorTileStyle.Background(lipgloss.Color(hex))
	return colorPreviewStyle.Render(colorTileStyle.String() + " #" + hex)
}

func LivePreview() string {
	p := tea.NewProgram(NewConfigureScreen())
	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	result, ok := m.(Configure)
	if !ok {
		log.Print("Failed type assertion in getLine()")
		return ""
	}
	v := m.View().Content
	lines := strings.Split(v, "\n")
	fmt.Println(lines[len(lines)-1])
	if result.wantsOutput {
		return string(result.config.String())
	}
	return ""
}
