package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var spinnerCmd tea.Cmd
	m.spinner, spinnerCmd = m.spinner.Update(msg)
	cmds = append(cmds, spinnerCmd)

	switch msg := msg.(type) {
	//if key is press...
	case tea.KeyMsg:
		//	key is press, let's see what key is pressed
		switch msg.Type {
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		}

		//we only allow user to type input in ASKING state
		if m.state == ASKING {
			//In ASKING state and user press enter
			if msg.Type == tea.KeyEnter {
				folderPath := strings.TrimSpace(m.textInput.Value())
				m.parentFolder = folderPath
				return m, func() tea.Msg {
				}
			}
			var textCmd tea.Cmd
			m.textInput, textCmd = m.textInput.Update(msg)
			cmds = append(cmds, textCmd)
		}

	}
	return m, tea.Batch(cmds...) //unpacks []Cmd into ...Cmd
}
