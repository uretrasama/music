package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	dir, errf := os.ReadDir(os.Args[1])
	fileNames := []string{}

	if errf != nil {
		fmt.Println(errf)
		os.Exit(1)
	}

	for _, file := range dir {
		fileNames = append(fileNames, file.Name())
	}

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
		case "up", "k":
			m.count++
		case "down", "j":
			m.count--
		}
	}

	return m, nil
}

func (m *Model) View() string {
	return fmt.Sprintf("count: %d\n\n up\tdown", m.count)
}
