package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var spinnerCmd tea.Cmd
	m.spinner, spinnerCmd = m.spinner.Update(msg)
	cmds = append(cmds, spinnerCmd)

	switch msg := msg.(type) {
	//if key is press...
	case tea.KeyMsg:
		//	key is press, let's see what key is pressed
		if msg.String() == "ctrl+c" || msg.Type == tea.KeyEnter {
			return m, tea.Quit
		}
	}
	return m, tea.Batch(cmds...) //unpacks []Cmd into ...Cmd
}
