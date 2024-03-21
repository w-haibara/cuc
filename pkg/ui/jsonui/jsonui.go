package jsonui

import (
	"encoding/json"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/w-haibara/cuc/pkg/ui"
)

type Model struct {
	Obj any
}

func NewModel(obj any) Model {
	return Model{obj}
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
	b, err := json.MarshalIndent(m.Obj, "", "  ")
	if err != nil {
		return heredoc.Docf(`
		{
			"error": %s
		}
		`, err.Error())
	}

	return fmt.Sprintf("%s\n", b)
}
