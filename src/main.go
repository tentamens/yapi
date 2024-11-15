package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var displayString = ""

type model struct {
	width     int
	height    int
	textinput textinput.Model
	err       error
}

func initModel() model {
	ti := textinput.New()
	ti.Placeholder = "Press Enter to search!"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20
	return model{textinput: ti}
}

func (m model) Init() tea.Cmd {

	return nil
}

func (m model) View() string {

	boxWidth := (m.width / 2) - 8
	boxStyle := lipgloss.NewStyle().
		Padding(2, 2).
		Width(boxWidth).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#5e81ac"))

	rightBox := boxStyle.Height(m.height - 4).Render("Left Box")
	layout := lipgloss.JoinHorizontal(lipgloss.Top, leftBox(m, boxStyle), rightBox)
	return lipgloss.NewStyle().Margin(1).Render(layout)
}

func leftBox(m model, containerStyle lipgloss.Style) string {
	bigBox := containerStyle.Height(m.height - 7).
		Render(displayString)
	ti := containerStyle.Height(1).
		Padding(0, 0).
		Render(m.textinput.View())

	layout := lipgloss.JoinVertical(lipgloss.Top, bigBox, ti)
	return layout
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			SearchRequest("hi")
		case "ctrl+q":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	}
	m.textinput, cmd = m.textinput.Update(msg)
	return m, cmd
}

func main() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
