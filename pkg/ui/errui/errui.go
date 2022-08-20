package errui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/w-haibara/cuc/pkg/ui"
)

type Model struct {
	err error
}

func NewModel(err error) Model {
	return Model{err}
}

func (m Model) Render() error {
	return ui.Render(m)
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m Model) View() string {
	return fmt.Sprintln("Error:", m.err.Error())
}
