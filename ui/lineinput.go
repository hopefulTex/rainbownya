package ui

import (
	"fmt"
	"log"
	"os"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/hopefulTex/rainbownya/rainbow"
)

var textBoxStyle lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

type lineInputModel struct {
	textBox      textinput.Model
	width        int
	height       int
	style        lipgloss.Style
	coloredInput bool
}

func (m lineInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m lineInputModel) Update(msg tea.Msg) (lineInputModel, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.style = m.style.Width(msg.Width - 2)
		m.textBox.SetWidth(msg.Width - len(m.textBox.Prompt) - 2 - 1)
	case tea.KeyPressMsg:
		m.textBox, cmd = m.textBox.Update(msg)
	}

	return m, cmd
}

func (m lineInputModel) View() string {
	var view strings.Builder
	var box string

	box = m.style.Render(m.textBox.View())

	view.WriteString(box)
	return view.String()
}

func (m lineInputModel) Value() string {
	return m.textBox.Value()
}

func (m *lineInputModel) Clear() {
	m.textBox.SetValue("")
}

func getLineModel() lineInputModel {
	t := textinput.New()
	t.Focus()
	t.SetWidth(20)
	t.Prompt = "> "
	t.Placeholder = "nya! :3c"

	return lineInputModel{
		width:   0,
		height:  0,
		textBox: t,
		style:   textBoxStyle,
	}
}

func GetLine() string {
	p := tea.NewProgram(initialModel(true, rainbow.DefaultVars()), tea.WithOutput(os.Stderr))
	m, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}

	result, ok := m.(model)
	if !ok {
		log.Print("Failed type assertion in getLine()")
		return ""
	}
	v := m.View().Content
	lines := strings.Split(v, "\n")
	fmt.Println(lines[len(lines)-1])
	return result.Result
}
