package mdui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/w-haibara/cuc/pkg/ui/message"
)

type Model struct {
	MD  string
	Cmd func(data any) tea.Cmd
}

func NewModel(cmd func(data any) tea.Cmd) Model {
	return Model{
		Cmd: cmd,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case message.InitDetailMsg:
		return m, m.Cmd(msg.Data)

	case message.ShowItemDetailMsg:
		m.MD = msg.MD
		return m, nil

	case tea.KeyMsg:
		if msg.String() == "esc" {
			return m, m.showListCmd
		}
		return m, nil

	default:
		return m, nil
	}
}

func (m Model) showListCmd() tea.Msg {
	return message.ShowListMsg{}
}

func (m Model) View() string {
	out, err := glamour.Render(m.MD, "dark")
	if err != nil {
		panic(err.Error())
	}

	return out
}
