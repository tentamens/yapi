package main

import (
	"fmt"
	"io"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var selectedResult = []item{}

type model struct {
	width     int
	height    int
	list      list.Model
	textinput textinput.Model
	inputMode int
}

type item struct {
	title string
}

func (i item) Title() string       { return i.title }
func (i item) FilterValue() string { return i.title }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	title := i.title
	if index == m.Index() {
		title = fmt.Sprintf("> %s", title) // Add a selection marker
	}

	fmt.Fprint(w, title)
}

func initModel() model {
	ti := textinput.New()
	ti.Placeholder = "Press Enter to search!"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	items := []list.Item{
		item{title: "Press enter to search"},
	}

	li := list.New(items, itemDelegate{}, 20, 20)
	li.Title = "hello world"
	return model{textinput: ti, inputMode: 1, list: li}
}

func (m model) Init() tea.Cmd {

	return nil
}

var processingData []AppHit

func (m model) View() string {

	boxWidth := (m.width / 2) - 8
	boxStyle := lipgloss.NewStyle().
		Padding(2, 2).
		Width(boxWidth).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#5e81ac"))

	rightBox := boxStyle.Height(m.height - 4).Width((m.width / 2) + 26).Render(renderRightBox())
	layout := lipgloss.JoinHorizontal(lipgloss.Top, leftBox(m, boxStyle), rightBox)
	return lipgloss.NewStyle().Margin(1).Render(layout)
}

func leftBox(m model, containerStyle lipgloss.Style) string {
	m.list.SetSize(20, 20)

	bigBox := containerStyle.Height(m.height - 7).
		Width((m.width / 2) - 30).
		Render(m.list.View())

	ti := containerStyle.Height(1).
		Width((m.width/2)-30).
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
		case "ctrl+q":
			return m, tea.Quit
		case "enter":
			processingData = append(SearchRequest("Discord").Hits)
			items := []list.Item{}
			for _, element := range processingData {
				items = append(items, item{title: element.Name})
			}
			m.list = list.New(items, itemDelegate{}, 0, 0)
			m.list.SetShowStatusBar(false)
			m.inputMode = 2
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	}
	switch m.inputMode {
	case 1:
		m.textinput, cmd = m.textinput.Update(msg)
	case 2:

		m.list, cmd = m.list.Update(msg)
	}

	return m, cmd
}

func main() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
