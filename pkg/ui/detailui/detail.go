package detailui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/w-haibara/cuc/pkg/ui/message"
)

type Model struct {
	Keys  []string
	Data  []map[string]string
	Index int
}

func (m *Model) SetData(keys []string, data []map[string]string) {
	m.Keys = keys
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
	str := ""

	for _, key := range m.Keys {
		str += fmt.Sprintf("%s: %s\n", key, m.Data[m.Index][key])
	}

	return str
}
