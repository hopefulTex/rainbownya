package ui

import (
	"fmt"
	"log"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/hopefulTex/rainbownya/rainbow"
)

type model struct {
	Result  string
	textBox lineInputModel

	width  int
	height int

	view         []string
	history      []string
	historyIndex int
	offset       int
	increment    int

	vars rainbow.Variables

	getLineMode bool
}

func (m model) Init() tea.Cmd {
	return m.textBox.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.textBox, cmd = m.textBox.Update(msg)
		return m, cmd
	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			fmt.Println("")
			return m, tea.Quit
		case "enter":
			if m.getLineMode {
				m.Result = rainbow.Rainbow(m.textBox.Value(), m.offset, m.vars)
				return m, tea.Quit
			}
			if m.textBox.Value() != "" {
				m.view = append(m.view, rainbow.Rainbow(m.textBox.Value(), m.offset, m.vars))
				m.history = append(m.history, m.textBox.Value())
				m.offset += m.increment
				m.textBox.Clear()
				m.historyIndex = len(m.history)
			}
		case "up":
			if m.historyIndex > 0 {
				m.historyIndex--
			}
			if len(m.history) > m.historyIndex {
				m.textBox.textBox.SetValue(m.history[m.historyIndex])
			}
		case "down":
			if m.historyIndex < len(m.history) {
				m.historyIndex++
			}
			if m.historyIndex >= len(m.history) {
				m.textBox.textBox.Reset()
			} else {
				m.textBox.textBox.SetValue(m.history[m.historyIndex])
			}
		default:
			m.textBox, cmd = m.textBox.Update(msg)
		}
	}

	return m, cmd
}

func (m model) View() tea.View {
	var view strings.Builder

	var offset int = 0
	offset = max(0, len(m.view)-m.height-3)

	view.WriteString(strings.Join(m.view[offset:], "\n"))
	if len(m.view) > 0 {
		view.WriteRune('\n')
	}

	viewHeight := lipgloss.Height(view.String())
	boxHeight := lipgloss.Height(m.textBox.View())

	for viewHeight+boxHeight < m.height {
		viewHeight++
		view.WriteRune('\n')
	}

	view.WriteString(m.textBox.View())
	return tea.NewView(view.String())
}

func initialModel(getLine bool, vars rainbow.Variables) tea.Model {

	t := getLineModel()
	return model{
		width:        0,
		height:       0,
		textBox:      t,
		view:         make([]string, 0),
		history:      make([]string, 0),
		offset:       0,
		increment:    4,
		vars:         vars,
		getLineMode:  getLine,
		historyIndex: 0,
		Result:       "",
	}
}
func Interactive(vars rainbow.Variables) string {
	p := tea.NewProgram(initialModel(false, vars))
	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	return ""
}
