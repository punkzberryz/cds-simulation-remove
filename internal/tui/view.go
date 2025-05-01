package tui

import "fmt"

func (m Model) View() string {
	s := "Hello...\n\n"
	s += fmt.Sprintf("%s loading...\n", m.spinner.View())
	s += fmt.Sprintf("Directory... %s\n", m.textInput.View())
	s += "\nPress q to quit.\n"
	return s
}
