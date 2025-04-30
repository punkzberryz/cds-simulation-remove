package tui

import "fmt"

func (m Model) View() string {
	s := "Hello...\n\n"
	s += fmt.Sprintf("%s loading...", m.spinner.View())
	s += "\nPress q to quit.\n"
	return s
}
