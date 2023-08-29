package main

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var textBoxStyle lipgloss.Style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

type model struct {
	textBox textinput.Model
	width   int
	height  int
	view    []string
	offset  int
	vars    variables
}

func (m model) Init() tea.Cmd {
	return m.textBox.Cursor.BlinkCmd()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		textBoxStyle.Width(msg.Width - 2)
		m.textBox.Width = msg.Width - len(m.textBox.Prompt) - 2 - 1
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "enter":
			if m.textBox.Value() != "" {
				m.view = append(m.view, rainbow(m.textBox.Value(), m.offset, m.vars))
				// offset should be 1, but users would get impatient
				m.offset += 4
				m.textBox.SetValue("")
			}
		default:
			m.textBox, cmd = m.textBox.Update(msg)
		}
	}

	return m, cmd
}

func (m model) View() string {
	var view strings.Builder

	var box string
	var offset int = 0

	box = textBoxStyle.Render(m.textBox.View())
	offset = len(m.view) - m.height - 3
	if offset < 0 {
		offset = 0
	}

	view.WriteString(strings.Join(m.view[offset:], "\n"))
	if len(m.view) > 0 {
		view.WriteRune('\n')
	}

	view.WriteString(box)
	return view.String()
}

func initialModel(vars variables) tea.Model {
	t := textinput.New()
	t.Focus()
	t.Width = 20
	t.Prompt = "> "
	t.Placeholder = "nya! :3c"

	return model{
		width:   0,
		height:  0,
		textBox: t,
		view:    make([]string, 0),
		offset:  0,
		vars:    vars,
	}
}
func interactive(vars variables) {
	p := tea.NewProgram(initialModel(vars))
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
