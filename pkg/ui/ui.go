package ui

import tea "github.com/charmbracelet/bubbletea"

type Model interface {
	Render() error
	Init() tea.Cmd
	Update(tea.Msg) (tea.Model, tea.Cmd)
	View() string
}

func Render(m Model) error {
	prog := tea.NewProgram(m)
	if err := prog.Start(); err != nil {
		return err
	}

	return nil
}
