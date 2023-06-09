package config

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Path      string
	FileNames []string
	BaseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.HiddenBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderTop(false).
			BorderLeft(false)
)
