package detailui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/w-haibara/cuc/pkg/ui/message"
)

type Model struct {
	Title string
	Desc  string
	Data  []map[string]string
	Index int
}

func (m *Model) SetTitle(title string) {
	m.Title = title
}

func (m *Model) SetDesc(desc string) {
	m.Desc = desc
}

func (m *Model) SetData(data []map[string]string) {
	m.Data = data
}

func (m *Model) SetIndex(i int) error {
	if i < 0 || i >= len(m.Data) {
		return fmt.Errorf("Invalid index: %d", i)
	}

	m.Index = i
	return nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "esc" {
			return m, m.showListCmd
		}
	}

	return m, nil
}

func (m Model) showListCmd() tea.Msg {
	return message.ShowListMsg{}
}

func (m Model) View() string {
	in := ""

	if m.Title != "" {
		in += fmt.Sprintln("#", m.Title)
	}

	if m.Desc != "" {
		in += m.Desc + "\n"
	}

	in += "| Key | Value |\n| --- | --- |\n"
	for k, v := range m.Data[m.Index] {
		in += fmt.Sprintf("| %s | %s |\n", k, v)
	}

	out, err := glamour.Render(in, "dark")
	if err != nil {
		panic(err.Error())
	}

	return out
}
