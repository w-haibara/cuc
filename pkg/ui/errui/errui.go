package errui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/w-haibara/cuc/pkg/ui"
)

type ErrModel struct {
	err error
}

func NewErrModel(err error) ErrModel {
	return ErrModel{err}
}

func (m ErrModel) Render() error {
	return ui.Render(m)
}

func (m ErrModel) Init() tea.Cmd {
	return nil
}

func (m ErrModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m ErrModel) View() string {
	return fmt.Sprintln("Error:", m.err.Error())
}
