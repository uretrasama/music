package main

import (
	"./parse"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	path      string
	fileNames []string
)

func main() {
	if len(os.Args) == 1 {
		path = "."
	} else {
		path = os.Args[1]
	}

	dir, errf := os.ReadDir(path)

	if errf != nil {
		fmt.Println(errf)
		os.Exit(1)
	}

	fileNames = parse.ParseFileNames(dir)

	err := tea.NewProgram(&Model{}, tea.WithAltScreen()).Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Model struct {
	count int
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "j":
			if m.count == (len(fileNames) - 1) {
				m.count = m.count
			} else {
				m.count++
			}
		case "down", "k":
			if m.count <= 0 {
				m.count = m.count
			} else {
				m.count--
			}
		}
	}

	return m, nil
}

func (m *Model) View() string {
	return fmt.Sprintf("%s    %d ", fileNames[m.count], m.count)
}
