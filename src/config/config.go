package config

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	Path      string
	FileNames []string
	BaseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))
)
