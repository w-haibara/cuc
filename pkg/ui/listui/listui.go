package listui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/w-haibara/cuc/pkg/ui"
	"github.com/w-haibara/cuc/pkg/ui/errui"
)

type ListModel struct {
	List list.Model
	Cmd  func() tea.Msg
}

func NewListModel(title string, cmd func() tea.Msg) ListModel {
	m := ListModel{
		List: list.New(nil, list.NewDefaultDelegate(), 1, 1),
		Cmd:  cmd,
	}
	m.List.Title = title
	m.List.StartSpinner()
	return m
}

func (m ListModel) Render() error {
	return ui.Render(m)
}

func (m ListModel) Init() tea.Cmd {
	return tea.Batch(
		m.Cmd,
		m.List.StartSpinner(),
	)
}

func (m ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.List.SetSize(msg.Width, msg.Height)

	case ListMsg:
		m.List.StopSpinner()
		m.List.Title = msg.Title
		m.List.SetItems(msg.Items)

	case error:
		return errui.NewErrModel(msg), cmd
	}

	return m, cmd
}

func (m ListModel) View() string {
	return m.List.View()
}

type ListMsg struct {
	Title string
	Items []list.Item
}

func NewListMsg(title string, items []list.Item) ListMsg {
	return ListMsg{
		Title: title,
		Items: items,
	}
}

type ListItem struct {
	Title_ string
	Desc   string
}

func MakeListItems(size int) *[]list.Item {
	items := make([]list.Item, 0, size)
	return &items
}

func AppendItem(items *[]list.Item, title, desc string) {
	*items = append(*items, ListItem{title, desc})
}

func (item ListItem) Title() string       { return item.Title_ }
func (item ListItem) Description() string { return item.Desc }
func (item ListItem) FilterValue() string { return item.Title_ }
