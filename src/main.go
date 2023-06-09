package main

import (
	"./config"
	"./parse"
	"fmt"
	"os"
	"regexp"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	parse.ParseToSlice()

	colums := []table.Column{
		{Title: "", Width: 70},
	}

	rows := []table.Row{}

	a := regexp.MustCompile("")
	for _, file := range config.FileNames {
		rows = append(rows, a.Split(file, 1))
	}

	t := table.New(
		table.WithColumns(colums),
		table.WithRows(rows),
		table.WithFocused(true),
	)

	s := table.DefaultStyles()

	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)

	t.SetStyles(s)

	if err := tea.NewProgram(&Model{t, 0}, tea.WithAltScreen()).Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Model struct {
	table table.Model
	count int
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "j":
			if m.count == (len(config.FileNames) - 1) {
				break
			} else {
				m.count++
			}
		case "down", "k":
			if m.count <= 0 {
				break
			} else {
				m.count--
			}
		}
	}

	m.table, cmd = m.table.Update(msg)

	return m, cmd
}

func (m *Model) View() string {
	return config.BaseStyle.Render(m.table.View() + "\n")
}
