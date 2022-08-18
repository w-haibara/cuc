package jsonui

import (
	"encoding/json"
	"fmt"

	"github.com/MakeNowJust/heredoc"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/w-haibara/cuc/pkg/ui"
)

type JsonModel struct {
	Obj any
}

func NewJsonModel(obj any) JsonModel {
	return JsonModel{obj}
}

func (m JsonModel) Render() error {
	return ui.Render(m)
}

func (m JsonModel) Init() tea.Cmd {
	return nil
}

func (m JsonModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, tea.Quit
}

func (m JsonModel) View() string {
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
