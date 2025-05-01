package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state string

const (
	ASKING    state = "ASKING"
	SEARCHING state = "SEARCHING"
)

type Model struct {
	cursor            int              //which list item you want to delete...
	selected          map[int]struct{} //which items are selected
	dataPaths         []string         //directory containing simulation results
	parentFolder      string
	state             state         //state of the program
	availableRoutines int           //available routines
	spinner           spinner.Model //Spinner ui
	spinnerFrame      int
	textInput         textinput.Model
}

func InitialModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Monkey
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	ti := textinput.New()
	ti.Placeholder = "Enter folder path (e.g. /kang/project/)"
	ti.Focus()
	ti.Width = 60
	return Model{
		state:             ASKING,
		cursor:            0,
		selected:          make(map[int]struct{}),
		dataPaths:         []string{},
		availableRoutines: 10,
		spinner:           s,
		textInput:         ti,
		parentFolder:      "",
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
		m.spinner.Tick,
	)
}

type findFilesMsg struct {
	files []string
}
